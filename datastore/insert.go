package datastore

import (
	"fmt"
	"reflect"
	"sync/atomic"
)

// InsertRow 插入数据，长度必须和heads长度一样
func (c *DataStore) InsertRow(row Row) (err error) {
	err = c.checkRow(row)
	if err != nil {
		return err
	}
	c.appendRow(row)
	return nil
}

func (c *DataStore) appendRow(row Row) {
	row.SetID(atomic.AddUint64(&c.rowIncreaseID, 1))
	c.syncTable.Lock()
	c.store = append(c.store, row)
	c.syncTable.Unlock()
	atomic.AddUint64(&c.rowCount, 1)
}

// InsertRows 要么全成功，要么全失败。
func (c *DataStore) InsertRows(rows []Row) (err error) {
	if len(rows) == 0 {
		return nil
	}
	for _, row := range rows {
		err = c.checkRow(row)
		if err != nil {
			return err
		}
	}
	for _, row := range rows {
		c.appendRow(row)
	}
	return nil
}

// 严格检查数据录入,如果为字段为nil会失败
func (c *DataStore) checkRow(row Row) (err error) {
	if row == nil {
		return fmt.Errorf("row is nil")
	}

	// 判断row是否为指针类型
	rt := reflect.TypeOf(row)
	if rt == nil {
		return fmt.Errorf("row is nil")
	}
	if rt.Kind() != reflect.Ptr {
		return fmt.Errorf("row: non-pointer " + rt.Kind().String())
	}

	// 判断row的长度是否和columnLength相等
	rowLength := row.Length()
	if rowLength != c.columnLength {
		return fmt.Errorf("row length is error")
	}

	// 判断每个字段值和类型
	var errNull error
	var errType error
	row.Range(func(i int, v interface{}) {
		if errNull != nil || errType != nil {
			return
		}
		if v == nil {
			errNull = fmt.Errorf("row column is null:%s", c.columnSlice[i].ColumnName)
		}
		rt := reflect.TypeOf(v)
		if rt == nil {
			errType = fmt.Errorf("row column is nil:%s", c.columnSlice[i].ColumnName)
		}
		if rt != c.columnSlice[i].columnType {
			errType = fmt.Errorf("row column type is error:%s,%s,%s", c.columnSlice[i].ColumnName,
				c.columnSlice[i].columnType.String(), rt.String())
		}
	})
	if errNull != nil {
		return errNull
	}

	if errType != nil {
		return errType
	}
	return nil
}

package datastore

import (
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
)

// DataStore 二维数据表
type DataStore struct {
	// 保存数据
	store     []Row
	syncTable *sync.RWMutex

	// 字段初始化初始化的时候完成
	columnMap    map[string]*ColumnDef
	columnSlice  []*ColumnDef
	columnLength int

	rowIncreaseID uint64
	rowCount      uint64
}

// NewDataStore 初始化后，row的长度就固定了，并且同一列的数据类型必须是一样，并且row必须是指针类型。
func NewDataStore(columns ...ColumnDef) (ds *DataStore, err error) {
	if len(columns) == 0 {
		return nil, fmt.Errorf("columns is empty")
	}
	columnMap := make(map[string]*ColumnDef)
	columnSlice := make([]*ColumnDef, len(columns))
	for k, v := range columns {
		if v.ZeroValue == nil {
			return nil, fmt.Errorf("columns zeroValue is nil:%s", v.ColumnName)
		}
		col := &ColumnDef{
			id:             k,
			ColumnName:     v.ColumnName,
			ZeroValue:      v.ZeroValue,
			columnType:     reflect.TypeOf(v.ZeroValue),
			CalculateSlice: v.CalculateSlice,
			Compare:        v.Compare,
		}
		columnMap[v.ColumnName] = col
		columnSlice[k] = col
	}

	ds = &DataStore{
		store:        make([]Row, 0),
		columnMap:    columnMap,
		columnSlice:  columnSlice,
		columnLength: len(columns),
		syncTable:    new(sync.RWMutex),
	}
	return ds, nil
}

// GetRowCount 获取总行数
func (c *DataStore) GetRowCount() (count uint64) {
	return atomic.LoadUint64(&c.rowCount)
}

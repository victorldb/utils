package datastore

import (
	"fmt"
	"reflect"
)

// ColumnDef --
type ColumnDef struct {
	id int

	// 列名
	ColumnName string

	// zero值
	ZeroValue interface{}
	// value类型由ZeroValue反推得到
	columnType reflect.Type

	// 集合处理,把符合条件数据放到一个切片里，统一处理
	CalculateSlice func(values []interface{}) (newValue interface{}, err error)

	// 比较 CompareLess:a<b CompareEqual:a==b CompareGreater:a>b
	Compare func(a, b interface{}) (res CompareType, err error)
}

func (c *DataStore) getColumnID(column string) (id int, ok bool) {
	colDef, colOK := c.columnMap[column]
	if !colOK {
		return id, false
	}
	id = colDef.id
	return id, true
}

func (c *DataStore) getColInfoWithNames(colNames []string) (currentCols []*ColumnDef, err error) {
	if len(colNames) == 0 {
		return nil, fmt.Errorf("colNames is empty")
	}
	currentCols = make([]*ColumnDef, len(colNames))
	for k, v := range colNames {
		nv, ok := c.columnMap[v]
		if !ok {
			return nil, fmt.Errorf("illegal column:%s", v)
		}
		currentCols[k] = nv
	}
	return currentCols, nil
}

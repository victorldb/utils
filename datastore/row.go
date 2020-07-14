package datastore

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Row --
type Row interface {
	GetItem(index int) (value interface{})
	Range(handle func(i int, v interface{}))
	SetID(id uint64)
	GetID() (id uint64)
	Length() int
}

// DataRow rowData里面不容许为nil，长度受table影响
type DataRow struct {
	id        uint64
	rowData   []interface{}
	rowLength int
	syncRow   sync.RWMutex
}

// NewDataRow --
func NewDataRow(data []interface{}) (row *DataRow) {
	row = &DataRow{
		rowData:   data,
		rowLength: len(data),
	}
	return row
}

// GetItem 不做index越界检查
func (c *DataRow) GetItem(index int) (value interface{}) {
	c.syncRow.RLock()
	value = c.rowData[index]
	c.syncRow.RUnlock()
	return value
}

// Length --
func (c *DataRow) Length() int {
	return c.rowLength
}

// Range --
func (c *DataRow) Range(handle func(i int, v interface{})) {
	c.syncRow.RLock()
	for i, v := range c.rowData {
		handle(i, v)
	}
	c.syncRow.RUnlock()
}

// SetID --
func (c *DataRow) SetID(id uint64) {
	atomic.CompareAndSwapUint64(&c.id, 0, id)
}

// GetID --
func (c *DataRow) GetID() (id uint64) {
	return atomic.LoadUint64(&c.id)
}

// String --
func (c *DataRow) String() (s string) {
	c.syncRow.RLock()
	s = fmt.Sprintf("%+v", c.rowData)
	c.syncRow.RUnlock()
	return s
}

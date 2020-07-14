package datastore

// GetItemInt --
func (c *DataStore) GetItemInt(rowIndex int, column string) (value int, ok bool) {
	nativeValue := c.GetItemWithIndex(rowIndex, column)
	value, ok = nativeValue.(int)
	if !ok {
		return value, false
	}
	return value, true
}

// GetItemInt8 --
func (c *DataStore) GetItemInt8(rowIndex int, column string) (value int8, ok bool) {
	nativeValue := c.GetItemWithIndex(rowIndex, column)
	value, ok = nativeValue.(int8)
	if !ok {
		return value, false
	}
	return value, true
}

// GetItemInt16 --
func (c *DataStore) GetItemInt16(rowIndex int, column string) (value int16, ok bool) {
	nativeValue := c.GetItemWithIndex(rowIndex, column)
	value, ok = nativeValue.(int16)
	if !ok {
		return value, false
	}
	return value, true
}

// GetItemInt32 --
func (c *DataStore) GetItemInt32(rowIndex int, column string) (value int32, ok bool) {
	nativeValue := c.GetItemWithIndex(rowIndex, column)
	value, ok = nativeValue.(int32)
	if !ok {
		return value, false
	}
	return value, true
}

// GetItemInt64 --
func (c *DataStore) GetItemInt64(rowIndex int, column string) (value int64, ok bool) {
	nativeValue := c.GetItemWithIndex(rowIndex, column)
	value, ok = nativeValue.(int64)
	if !ok {
		return value, false
	}
	return value, true
}

// GetItemUint --
func (c *DataStore) GetItemUint(rowIndex int, column string) (value uint, ok bool) {
	nativeValue := c.GetItemWithIndex(rowIndex, column)
	value, ok = nativeValue.(uint)
	if !ok {
		return value, false
	}
	return value, true
}

// GetItemUint8 --
func (c *DataStore) GetItemUint8(rowIndex int, column string) (value uint8, ok bool) {
	nativeValue := c.GetItemWithIndex(rowIndex, column)
	value, ok = nativeValue.(uint8)
	if !ok {
		return value, false
	}
	return value, true
}

// GetItemUint16 --
func (c *DataStore) GetItemUint16(rowIndex int, column string) (value uint16, ok bool) {
	nativeValue := c.GetItemWithIndex(rowIndex, column)
	value, ok = nativeValue.(uint16)
	if !ok {
		return value, false
	}
	return value, true
}

// GetItemUint32 --
func (c *DataStore) GetItemUint32(rowIndex int, column string) (value uint32, ok bool) {
	nativeValue := c.GetItemWithIndex(rowIndex, column)
	value, ok = nativeValue.(uint32)
	if !ok {
		return value, false
	}
	return value, true
}

// GetItemUint64 --
func (c *DataStore) GetItemUint64(rowIndex int, column string) (value uint64, ok bool) {
	nativeValue := c.GetItemWithIndex(rowIndex, column)
	value, ok = nativeValue.(uint64)
	if !ok {
		return value, false
	}
	return value, true
}

// GetItemFloat32 --
func (c *DataStore) GetItemFloat32(rowIndex int, column string) (value float32, ok bool) {
	nativeValue := c.GetItemWithIndex(rowIndex, column)
	value, ok = nativeValue.(float32)
	if !ok {
		return value, false
	}
	return value, true
}

// GetItemFloat64 --
func (c *DataStore) GetItemFloat64(rowIndex int, column string) (value float64, ok bool) {
	nativeValue := c.GetItemWithIndex(rowIndex, column)
	value, ok = nativeValue.(float64)
	if !ok {
		return value, false
	}
	return value, true
}

// GetItemString --
func (c *DataStore) GetItemString(rowIndex int, column string) (value string, ok bool) {
	nativeValue := c.GetItemWithIndex(rowIndex, column)
	value, ok = nativeValue.(string)
	if !ok {
		return value, false
	}
	return value, true
}

// GetItemBool --
func (c *DataStore) GetItemBool(rowIndex int, column string) (value bool, ok bool) {
	nativeValue := c.GetItemWithIndex(rowIndex, column)
	value, ok = nativeValue.(bool)
	if !ok {
		return value, false
	}
	return value, true
}

// GetItemRune --
func (c *DataStore) GetItemRune(rowIndex int, column string) (value rune, ok bool) {
	nativeValue := c.GetItemWithIndex(rowIndex, column)
	value, ok = nativeValue.(rune)
	if !ok {
		return value, false
	}
	return value, true
}

// GetItemWithIndex --
func (c *DataStore) GetItemWithIndex(rowIndex int, column string) (value interface{}) {
	if rowIndex > int(c.GetRowCount()) || rowIndex < 1 {
		return nil
	}
	colIndex, ok := c.getColumnID(column)
	if !ok {
		return nil
	}
	c.syncTable.Lock()
	value = c.store[rowIndex-1].GetItem(colIndex)
	c.syncTable.Unlock()
	return value
}

// GetRow --
func (c *DataStore) GetRow(rowIndex int) (row Row, ok bool) {
	if rowIndex > int(c.GetRowCount()) || rowIndex < 1 {
		return nil, false
	}
	row = c.store[rowIndex-1]
	return row, true
}

// GetAllRow --
func (c *DataStore) GetAllRow() (rows []Row, ok bool) {
	c.syncTable.RLock()
	count := int(c.GetRowCount())
	if count == 0 {
		c.syncTable.RUnlock()
		return nil, false
	}
	rows = make([]Row, count)
	for i := 0; i < count; i++ {
		rows[i] = c.store[i]
	}
	c.syncTable.RUnlock()
	return rows, true
}

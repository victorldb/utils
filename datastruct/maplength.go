package datastruct

import (
	"sync"
)

// SpecLenMP --
// 缓存指定长度的src==>value
type SpecLenMP struct {
	length            int
	currentLength     int
	currentDelIndex   int
	currentWriteIndex int
	keys              []interface{}
	mp                map[interface{}]interface{}

	syncmp sync.RWMutex
}

// NewSpecLenMP --
func NewSpecLenMP(length int) (nat *SpecLenMP) {
	nat = &SpecLenMP{
		length: length,
		keys:   make([]interface{}, length),
		mp:     make(map[interface{}]interface{}),
	}
	return nat
}

// Set --
func (c *SpecLenMP) Set(key, value interface{}) {
	c.syncmp.Lock()
	if c.currentLength == c.length {
		delKey := c.keys[c.currentDelIndex]
		delete(c.mp, delKey)
		c.currentDelIndex++
		if c.currentDelIndex == c.length {
			c.currentDelIndex = 0
		}
	}
	c.keys[c.currentWriteIndex] = key
	c.currentWriteIndex++
	if c.currentWriteIndex == c.length {
		c.currentWriteIndex = 0
	}
	c.mp[key] = value
	if c.currentLength < c.length {
		c.currentLength++
	}
	c.syncmp.Unlock()
}

// Get --
func (c *SpecLenMP) Get(key interface{}) (value interface{}, ok bool) {
	c.syncmp.RLock()
	value, ok = c.mp[key]
	c.syncmp.RUnlock()
	if value == nil {
		return nil, false
	}
	return value, ok
}

// Del --
func (c *SpecLenMP) Del(key interface{}) {
	c.syncmp.Lock()
	v, ok := c.mp[key]
	if ok && v != nil {
		c.mp[key] = nil
	}
	c.syncmp.Unlock()
}

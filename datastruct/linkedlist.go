package datastruct

import (
	"sync"
)

/*
	单向长度限制链表
	当链表长度达到限定长度时，先删除链表中最早录入的数据，再追加新的数据
*/

type singleLinkData struct {
	data interface{}
	next *singleLinkData
}

// SingleLinkListDel --
type SingleLinkListDel struct {
	length        uint64
	currentLength uint64
	lastValue     *singleLinkData
	firstValue    *singleLinkData
	syncValue     sync.Mutex
}

// NewSingleLinkListDel --
func NewSingleLinkListDel(length uint64) (l *SingleLinkListDel) {
	l = &SingleLinkListDel{
		length: length,
	}
	return l
}

// Get 当数据为空时，客户端有义务控制延迟
func (c *SingleLinkListDel) Get() (v interface{}, ok bool) {
	c.syncValue.Lock()
	if c.currentLength == 0 {
		c.syncValue.Unlock()
		return nil, false
	}
	v = c.firstValue.data
	c.currentLength--
	if c.currentLength == 0 {
		c.firstValue = nil
		c.lastValue = nil
	} else {
		c.firstValue = c.firstValue.next
	}
	c.syncValue.Unlock()
	return v, true
}

// Set --
func (c *SingleLinkListDel) Set(v interface{}) (dv interface{}, ok bool) {
	nv := &singleLinkData{
		data: v,
	}
	c.syncValue.Lock()
	if c.currentLength == 0 {
		c.firstValue = nv
		c.lastValue = nv
		c.currentLength++
	} else if c.currentLength < c.length {
		c.lastValue.next = nv
		c.lastValue = nv
		c.currentLength++
	} else {
		c.lastValue.next = nv
		c.lastValue = nv
		dv = c.firstValue.data
		ok = true
		c.firstValue = c.firstValue.next
	}
	c.syncValue.Unlock()
	return dv, ok
}

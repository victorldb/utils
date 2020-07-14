package datastore

import (
	"fmt"
	"strings"
	"sync/atomic"
)

// LogicType --
type LogicType uint8

const (
	// AND --
	AND LogicType = 1

	// OR --
	OR LogicType = 2
)

var (
	conditionAddID uint64 = 1
)

// GetLogicTypeName --
func GetLogicTypeName(lgType LogicType) string {
	switch lgType {
	case AND:
		return "&&"
	case OR:
		return "||"
	}
	return ""
}

type colvalueCompare struct {
	colName string
	comType CompareType
	value   interface{}
	lgType  LogicType
}

// DsCondition --
type DsCondition struct {
	id uint64

	colCond []*colvalueCompare
	lgType  LogicType

	// 内联条件
	innerDsCond *DsCondition
	innerLgType LogicType

	lastDsCond *DsCondition
	nextDsCond *DsCondition
}

// NewDsCondition --
func NewDsCondition() (nc *DsCondition) {
	return &DsCondition{
		id:      atomic.AddUint64(&conditionAddID, 1),
		colCond: make([]*colvalueCompare, 0),
	}
}

// InnerWithAnd --
func (c *DsCondition) InnerWithAnd(oc *DsCondition) (nc *DsCondition) {
	return c.setInnerWithOr(oc, AND)
}

// InnerWithOr --
func (c *DsCondition) InnerWithOr(oc *DsCondition) (nc *DsCondition) {
	return c.setInnerWithOr(oc, OR)
}

func (c *DsCondition) setInnerWithOr(oc *DsCondition, lgType LogicType) (nc *DsCondition) {
	c.innerDsCond = oc
	c.innerLgType = lgType
	return c
}

// JoinWithAnd ()&&()
func (c *DsCondition) JoinWithAnd(oc *DsCondition) (nc *DsCondition) {
	return c.setJoinLogic(oc, AND)
}

// JoinWithOr ()||()
func (c *DsCondition) JoinWithOr(oc *DsCondition) (nc *DsCondition) {
	return c.setJoinLogic(oc, OR)
}

func (c *DsCondition) setJoinLogic(oc *DsCondition, lgType LogicType) (nc *DsCondition) {
	c.lgType = lgType
	c.nextDsCond = oc
	oc.lastDsCond = c
	return oc
}

// And &&
func (c *DsCondition) And() (nc *DsCondition) {
	return c.setLogic(AND)
}

// Or ||
func (c *DsCondition) Or() (nc *DsCondition) {
	return c.setLogic(OR)
}

func (c *DsCondition) setLogic(lgType LogicType) (nc *DsCondition) {
	length := len(c.colCond)
	if length == 0 {
		return c
	}
	c.colCond[length-1].lgType = lgType
	return c
}

// Greater >
func (c *DsCondition) Greater(colName string, value interface{}) (nc *DsCondition) {
	c.createNewCond(colName, value, Greater)
	return c
}

// Equal ==
func (c *DsCondition) Equal(colName string, value interface{}) (nc *DsCondition) {
	c.createNewCond(colName, value, Equal)
	return c
}

// Less <
func (c *DsCondition) Less(colName string, value interface{}) (nc *DsCondition) {
	c.createNewCond(colName, value, Less)
	return c
}

// GreaterEqual >=
func (c *DsCondition) GreaterEqual(colName string, value interface{}) (nc *DsCondition) {
	c.createNewCond(colName, value, GreaterEqual)
	return c
}

// LessEqual <=
func (c *DsCondition) LessEqual(colName string, value interface{}) (nc *DsCondition) {
	c.createNewCond(colName, value, LessEqual)
	return c
}

// NotEqual !=
func (c *DsCondition) NotEqual(colName string, value interface{}) (nc *DsCondition) {
	c.createNewCond(colName, value, NotEqual)
	return c
}

func (c *DsCondition) createNewCond(colName string, value interface{}, compType CompareType) {
	colVC := &colvalueCompare{
		colName: colName,
		comType: compType,
		value:   value,
	}
	c.colCond = append(c.colCond, colVC)
}

func (c *DsCondition) getCloumnsName() (colNames []string) {
	mp := make(map[string]int)
	c.allCloumnsName(c, mp)
	n := 0
	colNames = make([]string, len(mp))
	for k := range mp {
		colNames[n] = k
		n++
	}
	return colNames
}

func (c *DsCondition) allCloumnsName(oc *DsCondition, mp map[string]int) {
	swapDsCond := oc
	for {
		if swapDsCond.lastDsCond == nil {
			break
		}
		swapDsCond = swapDsCond.lastDsCond
	}
	for {
		for _, v := range swapDsCond.colCond {
			mp[v.colName] = 1
		}
		if swapDsCond.innerDsCond != nil {
			c.allCloumnsName(swapDsCond.innerDsCond, mp)
		}
		if swapDsCond.nextDsCond == nil {
			break
		}
		swapDsCond = swapDsCond.nextDsCond
	}
}

// Check 等待完成
func (c *DsCondition) Check() (err error) {

	return nil
}

func (c *DsCondition) legalityCheck(oc *DsCondition) (err error) {
	swapDsCond := oc
	for {
		if swapDsCond.lastDsCond == nil {
			break
		}
		swapDsCond = swapDsCond.lastDsCond
	}
	for {
		for _, v := range swapDsCond.colCond {
			if v.comType == 0 {
				return fmt.Errorf("comType failed")
			}
		}
		if swapDsCond.innerDsCond != nil {
			err = c.legalityCheck(swapDsCond.innerDsCond)
			if err != nil {
				return nil
			}
		}
		if swapDsCond.nextDsCond == nil {
			break
		}
		swapDsCond = swapDsCond.nextDsCond
	}
	return nil
}

// 通常情况下,采用字符串相加更有优势
func (c *DsCondition) String() (res string) {
	var left, right string
	c.decode(c, &left, &right)
	res = left + right
	return res
}

func (c *DsCondition) decode(oc *DsCondition, left, right *string) {
	swapDsCond := oc
	for {
		if swapDsCond.lastDsCond == nil {
			break
		}
		swapDsCond = swapDsCond.lastDsCond
	}

	*left = *left + "("
	*right = *right + ")"

	for {
		for _, v := range swapDsCond.colCond {
			*left = *left + v.colName + GetCompareTypeName(v.comType) + fmt.Sprintf("%+v", v.value) + GetLogicTypeName(v.lgType)
		}
		if swapDsCond.innerDsCond != nil {
			*left = *left + GetLogicTypeName(swapDsCond.innerLgType)
			c.decode(swapDsCond.innerDsCond, left, right)
		}
		*right = *right + GetLogicTypeName(swapDsCond.lgType)
		if swapDsCond.nextDsCond == nil {
			break
		}
		*left = *left + *right + "("
		*right = ")"
		swapDsCond = swapDsCond.nextDsCond
	}
}

// StringWithSlice 条件非常多时使用slice有一点优势
func (c *DsCondition) StringWithSlice() (res string) {
	var left, right []string
	c.decodeWithSlice(c, &left, &right)
	res = strings.Join(append(left, right...), "")
	return res
}

func (c *DsCondition) decodeWithSlice(oc *DsCondition, left, right *[]string) {
	swapDsCond := oc
	for {
		if swapDsCond.lastDsCond == nil {
			break
		}
		swapDsCond = swapDsCond.lastDsCond
	}

	*left = append(*left, "(")
	*right = append(*right, ")")

	for {
		for _, v := range swapDsCond.colCond {
			*left = append(*left, v.colName, GetCompareTypeName(v.comType), fmt.Sprintf("%+v", v.value), GetLogicTypeName(v.lgType))
		}
		if swapDsCond.innerDsCond != nil {
			*left = append(*left, GetLogicTypeName(swapDsCond.innerLgType))
			c.decodeWithSlice(swapDsCond.innerDsCond, left, right)
		}
		*right = append(*right, GetLogicTypeName(swapDsCond.lgType))
		if swapDsCond.nextDsCond == nil {
			break
		}
		*left = append(append(*left, *right...), "(")
		*right = []string{")"}
		swapDsCond = swapDsCond.nextDsCond
	}
}

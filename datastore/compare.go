package datastore

import (
	"fmt"
	"reflect"
	"time"
)

// CompareType --
type CompareType uint8

const (
	// Greater 大于
	Greater CompareType = 1 << 0

	// Equal 等于
	Equal CompareType = 1 << 1

	// Less 小于
	Less CompareType = 1 << 2

	// GreaterEqual 大于等于
	GreaterEqual CompareType = Greater | Equal

	// LessEqual 小于等于
	LessEqual CompareType = Equal | Less

	// NotEqual 不等于
	NotEqual CompareType = Greater | Less
)

// GetCompareTypeName --
func GetCompareTypeName(comType CompareType) string {
	switch comType {
	case Greater:
		return ">"
	case Equal:
		return "=="
	case Less:
		return "<"
	case GreaterEqual:
		return ">="
	case LessEqual:
		return "<="
	case NotEqual:
		return "!="
	}
	return ""
}

func negateCompareType(a CompareType) (v CompareType) {
	switch a {
	case Greater:
		v = LessEqual
	case GreaterEqual:
		v = Less
	case Equal:
		v = NotEqual
	case Less:
		v = GreaterEqual
	case LessEqual:
		v = Greater
	case NotEqual:
		v = Equal
	}
	return v
}

// CompareNumberFunc --
func CompareNumberFunc(a, b interface{}) (res CompareType, err error) {
	if a == nil {
		return res, fmt.Errorf("a is nil")
	}
	if b == nil {
		return res, fmt.Errorf("b is nil")
	}
	rav := reflect.ValueOf(a)
	rbv := reflect.ValueOf(b)

	rat := rav.Type()
	rbt := rbv.Type()

	if rat != rbt {
		return res, fmt.Errorf("value type is error:%s,%s", rat.String(), rbt.String())
	}
	switch rat.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		raIn64 := rav.Int()
		rbInt64 := rbv.Int()
		if raIn64 < rbInt64 {
			res = Less
		} else if raIn64 == rbInt64 {
			res = Equal
		} else {
			res = Greater
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		raUin64 := rav.Uint()
		rbUint64 := rbv.Uint()
		if raUin64 < rbUint64 {
			res = Less
		} else if raUin64 == rbUint64 {
			res = Equal
		} else {
			res = Greater
		}
	case reflect.Float32, reflect.Float64:
		raFloat64 := rav.Float()
		rbFloat64 := rbv.Float()
		if raFloat64 < rbFloat64 {
			res = Less
		} else if raFloat64 == rbFloat64 {
			res = Equal
		} else {
			res = Greater
		}
	default:
		return res, fmt.Errorf("unsport value type:%s", rat.Kind())
	}
	return res, nil
}

// CompareIntFunc --
func CompareIntFunc(a, b interface{}) (res CompareType, err error) {
	av, ok := a.(int)
	if !ok {
		return res, fmt.Errorf("a type is error")
	}
	bv, ok := b.(int)
	if !ok {
		return res, fmt.Errorf("b type is error")
	}
	if av < bv {
		res = Less
	} else if av == bv {
		res = Equal
	} else {
		res = Greater
	}
	return res, nil
}

// CompareInt64Func --
func CompareInt64Func(a, b interface{}) (res CompareType, err error) {
	av, ok := a.(int64)
	if !ok {
		return res, fmt.Errorf("a type is error")
	}
	bv, ok := b.(int64)
	if !ok {
		return res, fmt.Errorf("b type is error")
	}
	if av < bv {
		res = Less
	} else if av == bv {
		res = Equal
	} else {
		res = Greater
	}
	return res, nil
}

// CompareUint64Func --
func CompareUint64Func(a, b interface{}) (res CompareType, err error) {
	av, ok := a.(uint64)
	if !ok {
		return res, fmt.Errorf("a type is error")
	}
	bv, ok := b.(uint64)
	if !ok {
		return res, fmt.Errorf("b type is error")
	}
	if av < bv {
		res = Less
	} else if av == bv {
		res = Equal
	} else {
		res = Greater
	}
	return res, nil
}

// CompareFloat64Func --
func CompareFloat64Func(a, b interface{}) (res CompareType, err error) {
	av, ok := a.(float64)
	if !ok {
		return res, fmt.Errorf("a type is error")
	}
	bv, ok := b.(float64)
	if !ok {
		return res, fmt.Errorf("b type is error")
	}
	if av < bv {
		res = Less
	} else if av == bv {
		res = Equal
	} else {
		res = Greater
	}
	return res, nil
}

// CompareStringFunc --
func CompareStringFunc(a, b interface{}) (res CompareType, err error) {
	av, ok := a.(string)
	if !ok {
		return res, fmt.Errorf("a type is error")
	}
	bv, ok := b.(string)
	if !ok {
		return res, fmt.Errorf("b type is error")
	}
	if av < bv {
		res = Less
	} else if av == bv {
		res = Equal
	} else {
		res = Greater
	}
	return res, nil
}

// CompareTimeFunc --
func CompareTimeFunc(a, b interface{}) (res CompareType, err error) {
	av, ok := a.(time.Time)
	if !ok {
		return res, fmt.Errorf("a type is error")
	}
	bv, ok := b.(time.Time)
	if !ok {
		return res, fmt.Errorf("b type is error")
	}
	if av.Before(bv) {
		res = Less
	} else if av.Equal(bv) {
		res = Equal
	} else {
		res = Greater
	}
	return res, nil
}

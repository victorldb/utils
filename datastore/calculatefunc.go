package datastore

import (
	"fmt"
	"reflect"
)

// CalculateSumNumber --
func CalculateSumNumber(a, b interface{}) (nv interface{}, err error) {
	if a == nil {
		return nil, fmt.Errorf("a is nil")
	}
	if b == nil {
		return nil, fmt.Errorf("b is nil")
	}
	rav := reflect.ValueOf(a)
	rbv := reflect.ValueOf(b)

	rat := rav.Type()
	rbt := rbv.Type()

	if rat != rbt {
		return nil, fmt.Errorf("value type is error:%s,%s", rat.String(), rbt.String())
	}
	switch rat.Kind() {
	case reflect.Int:
		v := rav.Int() + rbv.Int()
		nv = int(v)
	case reflect.Int64:
		nv = rav.Int() + rbv.Int()
	case reflect.Uint64:
		nv = rav.Uint() + rbv.Uint()
	case reflect.Float64:
		nv = rav.Float() + rbv.Float()
	case reflect.Int8:
		v := rav.Int() + rbv.Int()
		nv = int8(v)
	case reflect.Int16:
		v := rav.Int() + rbv.Int()
		nv = int16(v)
	case reflect.Int32:
		v := rav.Int() + rbv.Int()
		nv = int32(v)
	case reflect.Uint:
		v := rav.Uint() + rbv.Uint()
		nv = uint(v)
	case reflect.Uint8:
		v := rav.Uint() + rbv.Uint()
		nv = uint8(v)
	case reflect.Uint16:
		v := rav.Uint() + rbv.Uint()
		nv = uint16(v)
	case reflect.Uint32:
		v := rav.Uint() + rbv.Uint()
		nv = uint32(v)
	case reflect.Float32:
		v := rav.Float() + rbv.Float()
		nv = float32(v)
	default:
		return nil, fmt.Errorf("unsport value type:%s", rat.Kind())
	}
	return nv, nil
}

// CalculateSumNumberSlice --
func CalculateSumNumberSlice(sli []interface{}) (nv interface{}, err error) {
	if sli == nil {
		return nil, fmt.Errorf("sli is nil")
	}

	var firstType reflect.Type
	for _, v := range sli {
		if v == nil {
			return nil, fmt.Errorf("CalculateSumNumberSlice vaulue is nil")
		}
		rt := reflect.TypeOf(v)
		if firstType == nil {
			firstType = rt
			continue
		}
		if rt != firstType {
			return nil, fmt.Errorf("value type is error:%s,%s", firstType.String(), rt.String())
		}
		kindType := rt.Kind()
		if kindType != reflect.Int && kindType != reflect.Int8 && kindType != reflect.Int16 &&
			kindType != reflect.Int32 && kindType != reflect.Int64 &&
			kindType != reflect.Uint && kindType != reflect.Uint8 && kindType != reflect.Uint16 &&
			kindType != reflect.Uint32 && kindType != reflect.Uint64 &&
			kindType != reflect.Float64 && kindType != reflect.Float32 {
			return nil, fmt.Errorf("value kind is error:%s", kindType.String())
		}
	}

	var add interface{}
	switch firstType.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		addInt64 := int64(0)
		for _, v := range sli {
			addInt64 += reflect.ValueOf(v).Int()
		}
		add = int64(addInt64)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		addUint64 := uint64(0)
		for _, v := range sli {
			addUint64 += reflect.ValueOf(v).Uint()
		}
		add = uint64(addUint64)
	case reflect.Float32, reflect.Float64:
		addFloat64 := float64(0)
		for _, v := range sli {
			addFloat64 += reflect.ValueOf(v).Float()
		}
		add = float64(addFloat64)
	default:
		return nil, fmt.Errorf("unsport value type:%s", firstType.Kind())
	}

	switch firstType.Kind() {
	case reflect.Int:
		nv = int(add.(int64))
	case reflect.Int8:
		nv = int8(add.(int64))
	case reflect.Int16:
		nv = int16(add.(int64))
	case reflect.Int32:
		nv = int32(add.(int64))
	case reflect.Int64:
		nv = add.(int64)
	case reflect.Uint:
		nv = uint(add.(uint64))
	case reflect.Uint8:
		nv = uint8(add.(uint64))
	case reflect.Uint16:
		nv = uint16(add.(uint64))
	case reflect.Uint32:
		nv = uint32(add.(uint64))
	case reflect.Uint64:
		nv = add.(uint64)
	case reflect.Float32:
		nv = float32(add.(float64))
	case reflect.Float64:
		nv = add.(float64)
	default:
		return nil, fmt.Errorf("unsport value type:%s", firstType.Kind())
	}
	return nv, nil
}

// CalculateSumInt --
func CalculateSumInt(a, b interface{}) (nv interface{}, err error) {
	av, ok := a.(int)
	if !ok {
		return 0, fmt.Errorf("a type is error")
	}
	bv, ok := b.(int)
	if !ok {
		return 0, fmt.Errorf("b type is error")
	}
	nv = av + bv
	return nv, nil
}

// CalculateSumInt64 --
func CalculateSumInt64(a, b interface{}) (nv interface{}, err error) {
	av, ok := a.(int64)
	if !ok {
		return 0, fmt.Errorf("a type is error")
	}
	bv, ok := b.(int64)
	if !ok {
		return 0, fmt.Errorf("b type is error")
	}
	nv = av + bv
	return nv, nil
}

// CalculateSumUint64 --
func CalculateSumUint64(a, b interface{}) (nv interface{}, err error) {
	av, ok := a.(uint64)
	if !ok {
		return 0, fmt.Errorf("a type is error")
	}
	bv, ok := b.(uint64)
	if !ok {
		return 0, fmt.Errorf("b type is error")
	}
	nv = av + bv
	return nv, nil
}

// CalculateSumFloat64 --
func CalculateSumFloat64(a, b interface{}) (nv interface{}, err error) {
	av, ok := a.(float64)
	if !ok {
		return 0, fmt.Errorf("a type is error")
	}
	bv, ok := b.(float64)
	if !ok {
		return 0, fmt.Errorf("b type is error")
	}
	nv = av + bv
	return nv, nil
}

// CalculateSumIntSlice --
func CalculateSumIntSlice(sli []interface{}) (nv interface{}, err error) {
	var add int
	for _, v := range sli {
		nv, ok := v.(int)
		if !ok {
			return 0, fmt.Errorf("a type is error")
		}
		add += nv
	}
	return add, nil
}

// CalculateSumInt64Slice --
func CalculateSumInt64Slice(sli []interface{}) (nv interface{}, err error) {
	var add int64
	for _, v := range sli {
		nv, ok := v.(int64)
		if !ok {
			return 0, fmt.Errorf("a type is error")
		}
		add += nv
	}
	return add, nil
}

// CalculateSumUint64Slice --
func CalculateSumUint64Slice(sli []interface{}) (nv interface{}, err error) {
	var add uint64
	for _, v := range sli {
		nv, ok := v.(uint64)
		if !ok {
			return 0, fmt.Errorf("a type is error")
		}
		add += nv
	}
	return add, nil
}

// CalculateSumFloat64Slice --
func CalculateSumFloat64Slice(sli []interface{}) (nv interface{}, err error) {
	var add float64
	for _, v := range sli {
		nv, ok := v.(float64)
		if !ok {
			return 0, fmt.Errorf("a type is error")
		}
		add += nv
	}
	return add, nil
}

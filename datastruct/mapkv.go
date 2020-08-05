package datastruct

// MapKV --
type MapKV map[string]interface{}

// GetValueWithFloat64 --
func (c MapKV) GetValueWithFloat64(k string) (v float64, ok bool) {
	nv, ok := c[k]
	if !ok {
		return v, false
	}
	v, ok = nv.(float64)
	return v, ok
}

// GetValueWithString --
func (c MapKV) GetValueWithString(k string) (v string, ok bool) {
	nv, ok := c[k]
	if !ok {
		return v, false
	}
	v, ok = nv.(string)
	return v, ok
}

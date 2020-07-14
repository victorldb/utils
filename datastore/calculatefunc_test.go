package datastore

import (
	"testing"
	"time"
)

func TestCalculateSumNumber(t *testing.T) {
	var a, b, v interface{}
	var err error
	// int
	a = int(1)
	b = int(2)
	v, err = CalculateSumNumber(a, b)
	if err != nil {
		t.Fatal(err)
	}
	if rv, ok := v.(int); !ok {
		t.Fatal("v type is error")
	} else {
		if rv != 3 {
			t.Fatalf("v is error:3,%+v", rv)
		}
	}

	// int8
	a = int8(1)
	b = int8(2)
	v, err = CalculateSumNumber(a, b)
	if err != nil {
		t.Fatal(err)
	}
	if rv, ok := v.(int8); !ok {
		t.Fatal("v type is error")
	} else {
		if rv != 3 {
			t.Fatalf("v is error:3,%+v", rv)
		}
	}

	// int16
	a = int16(1)
	b = int16(2)
	v, err = CalculateSumNumber(a, b)
	if err != nil {
		t.Fatal(err)
	}
	if rv, ok := v.(int16); !ok {
		t.Fatal("v type is error")
	} else {
		if rv != 3 {
			t.Fatalf("v is error:3,%+v", rv)
		}
	}

	// int32
	a = int32(1)
	b = int32(2)
	v, err = CalculateSumNumber(a, b)
	if err != nil {
		t.Fatal(err)
	}
	if rv, ok := v.(int32); !ok {
		t.Fatal("v type is error")
	} else {
		if rv != 3 {
			t.Fatalf("v is error:3,%+v", rv)
		}
	}

	// int64
	a = int64(1)
	b = int64(2)
	v, err = CalculateSumNumber(a, b)
	if err != nil {
		t.Fatal(err)
	}
	if rv, ok := v.(int64); !ok {
		t.Fatal("v type is error")
	} else {
		if rv != 3 {
			t.Fatalf("v is error:3,%+v", rv)
		}
	}

	// uint
	a = uint(1)
	b = uint(2)
	v, err = CalculateSumNumber(a, b)
	if err != nil {
		t.Fatal(err)
	}
	if rv, ok := v.(uint); !ok {
		t.Fatal("v type is error")
	} else {
		if rv != 3 {
			t.Fatalf("v is error:3,%+v", rv)
		}
	}

	// uint8
	a = uint8(1)
	b = uint8(2)
	v, err = CalculateSumNumber(a, b)
	if err != nil {
		t.Fatal(err)
	}
	if rv, ok := v.(uint8); !ok {
		t.Fatal("v type is error")
	} else {
		if rv != 3 {
			t.Fatalf("v is error:3,%+v", rv)
		}
	}

	// uint16
	a = uint16(1)
	b = uint16(2)
	v, err = CalculateSumNumber(a, b)
	if err != nil {
		t.Fatal(err)
	}
	if rv, ok := v.(uint16); !ok {
		t.Fatal("v type is error")
	} else {
		if rv != 3 {
			t.Fatalf("v is error:3,%+v", rv)
		}
	}

	// uint32
	a = uint32(1)
	b = uint32(2)
	v, err = CalculateSumNumber(a, b)
	if err != nil {
		t.Fatal(err)
	}
	if rv, ok := v.(uint32); !ok {
		t.Fatal("v type is error")
	} else {
		if rv != 3 {
			t.Fatalf("v is error:3,%+v", rv)
		}
	}

	// uint64
	a = uint64(1)
	b = uint64(2)
	v, err = CalculateSumNumber(a, b)
	if err != nil {
		t.Fatal(err)
	}
	if rv, ok := v.(uint64); !ok {
		t.Fatal("v type is error")
	} else {
		if rv != 3 {
			t.Fatalf("v is error:3,%+v", rv)
		}
	}

	// float32
	a = float32(1)
	b = float32(2)
	v, err = CalculateSumNumber(a, b)
	if err != nil {
		t.Fatal(err)
	}
	if rv, ok := v.(float32); !ok {
		t.Fatal("v type is error")
	} else {
		if rv != 3 {
			t.Fatalf("v is error:3,%+v", rv)
		}
	}

	// float64
	a = float64(1)
	b = float64(2)
	v, err = CalculateSumNumber(a, b)
	if err != nil {
		t.Fatal(err)
	}
	if rv, ok := v.(float64); !ok {
		t.Fatal("v type is error")
	} else {
		if rv != 3 {
			t.Fatalf("v is error:3,%+v", rv)
		}
	}

	// err check
	a = float32(1)
	b = int8(3)
	v, err = CalculateSumNumber(a, b)
	if err == nil {
		t.Fatal("float32 int8")
	}
}

func Benchmark_calNumber(b *testing.B) {
	var a, d, v interface{}
	var err error

	a = int(1)
	d = int(2)
	for i := 0; i < b.N; i++ {
		v, err = CalculateSumNumber(a, d)
		if err != nil {
			b.Fatal(err)
		}
		if rv, ok := v.(int); !ok {
			b.Fatal("v type is error")
		} else {
			if rv != 3 {
				b.Fatalf("v is error:3,%+v", rv)
			}
		}
	}
}

func Benchmark_calInt(b *testing.B) {
	var a, d, v interface{}
	var err error

	a = int(1)
	d = int(2)
	for i := 0; i < b.N; i++ {
		v, err = CalculateSumFloat64(a, d)
		if err != nil {
			b.Fatal(err)
		}
		if rv, ok := v.(int); !ok {
			b.Fatal("v type is error")
		} else {
			if rv != 3 {
				b.Fatalf("v is error:3,%+v", rv)
			}
		}
	}
}

func Test_calIntSlice(t *testing.T) {
	var v interface{}
	var err error

	sli := make([]interface{}, 1e7)
	for i := 0; i < 1e7; i++ {
		sli[i] = int(i)
	}

	now := time.Now()
	v, err = CalculateSumIntSlice(sli)
	if err != nil {
		t.Fatal(err)
	}
	if rv, ok := v.(int); !ok {
		t.Fatal("v type is error")
	} else {
		println(rv)
	}
	println(time.Since(now).String())
}

func Test_calInts(t *testing.T) {
	var v interface{}
	var add int
	var err error

	sli := make([]interface{}, 1e7)
	for i := 0; i < 1e7; i++ {
		sli[i] = int(i)
	}

	now := time.Now()
	for _, n := range sli {
		v, err = CalculateSumInt(add, n)
		if err != nil {
			t.Fatal(err)
		}
		if rv, ok := v.(int); !ok {
			t.Fatal("v type is error")
		} else {
			add = rv
		}
	}
	println(add)
	println(time.Since(now).String())
}

func Test_calNumbers(t *testing.T) {
	var v interface{}
	var add int
	var err error

	sli := make([]interface{}, 1e7)
	for i := 0; i < 1e7; i++ {
		sli[i] = int(i)
	}

	now := time.Now()
	for _, n := range sli {
		v, err = CalculateSumNumber(add, n)
		if err != nil {
			t.Fatal(err)
		}
		if rv, ok := v.(int); !ok {
			t.Fatal("v type is error")
		} else {
			add = rv
		}
	}
	println(add)
	println(time.Since(now).String())
}

func Test_calNumberSlice(t *testing.T) {
	var v interface{}
	var err error

	sli := make([]interface{}, 1e7)
	for i := 0; i < 1e7; i++ {
		sli[i] = int(i)
	}

	now := time.Now()
	v, err = CalculateSumNumberSlice(sli)
	if err != nil {
		t.Fatal(err)
	}
	if rv, ok := v.(int); !ok {
		t.Fatal("v type is error")
	} else {
		println(rv)
	}
	println(time.Since(now).String())
}

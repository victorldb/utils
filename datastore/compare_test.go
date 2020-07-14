package datastore

import (
	"testing"
	"time"
)

func TestCompareNumberFunc(t *testing.T) {
	var a, b, v interface{}
	var err error
	// int
	a = int(1)
	b = int(2)
	v, err = CompareNumberFunc(a, b)
	if err != nil {
		t.Fatal(err)
	}
	if v != Less {
		t.Fatalf("res is error")
	}

	// int8
	a = int8(1)
	b = int8(2)
	v, err = CompareNumberFunc(a, b)
	if err != nil {
		t.Fatal(err)
	}
	if v != Less {
		t.Fatalf("res is error")
	}

	// int16
	a = int16(1)
	b = int16(2)
	v, err = CompareNumberFunc(a, b)
	if err != nil {
		t.Fatal(err)
	}
	if v != Less {
		t.Fatalf("res is error")
	}

	// int32
	a = int32(1)
	b = int32(2)
	v, err = CompareNumberFunc(a, b)
	if err != nil {
		t.Fatal(err)
	}
	if v != Less {
		t.Fatalf("res is error")
	}

	// int64
	a = int64(1)
	b = int64(2)
	v, err = CompareNumberFunc(a, b)
	if err != nil {
		t.Fatal(err)
	}
	if v != Less {
		t.Fatalf("res is error")
	}

	// uint
	a = uint(1)
	b = uint(2)
	v, err = CompareNumberFunc(a, b)
	if err != nil {
		t.Fatal(err)
	}
	if v != Less {
		t.Fatalf("res is error")
	}

	// uint8
	a = uint8(1)
	b = uint8(2)
	v, err = CompareNumberFunc(a, b)
	if err != nil {
		t.Fatal(err)
	}
	if v != Less {
		t.Fatalf("res is error")
	}

	// uint16
	a = uint16(1)
	b = uint16(2)
	v, err = CompareNumberFunc(a, b)
	if err != nil {
		t.Fatal(err)
	}
	if v != Less {
		t.Fatalf("res is error")
	}

	// uint32
	a = uint32(1)
	b = uint32(2)
	v, err = CompareNumberFunc(a, b)
	if err != nil {
		t.Fatal(err)
	}
	if v != Less {
		t.Fatalf("res is error")
	}

	// uint64
	a = uint64(1)
	b = uint64(2)
	v, err = CompareNumberFunc(a, b)
	if err != nil {
		t.Fatal(err)
	}
	if v != Less {
		t.Fatalf("res is error")
	}

	// float32
	a = float32(1)
	b = float32(2)
	v, err = CompareNumberFunc(a, b)
	if err != nil {
		t.Fatal(err)
	}
	if v != Less {
		t.Fatalf("res is error")
	}

	// float64
	a = float64(1)
	b = float64(2)
	v, err = CompareNumberFunc(a, b)
	if err != nil {
		t.Fatal(err)
	}
	if v != Less {
		t.Fatalf("res is error")
	}

	// err check
	a = float32(1)
	b = int8(3)
	v, err = CompareNumberFunc(a, b)
	if err == nil {
		t.Fatal("failed")
	}

	// less
	a = int8(1)
	b = int8(3)
	v, err = CompareNumberFunc(a, b)
	if err != nil {
		t.Fatal("failed")
	}
	if v != Less {
		t.Fatalf("res is error")
	}

	// equal
	a = int8(1)
	b = int8(1)
	v, err = CompareNumberFunc(a, b)
	if err != nil {
		t.Fatal("failed")
	}
	if v != Equal {
		t.Fatalf("res is error")
	}

	// greater
	a = int8(3)
	b = int8(1)
	v, err = CompareNumberFunc(a, b)
	if err != nil {
		t.Fatal("failed")
	}
	if v != Greater {
		t.Fatalf("res is error")
	}
}

func Benchmark_comNumber(b *testing.B) {
	var a, d interface{}
	var v CompareType
	var err error
	a = int(3)
	d = int(1)
	for i := 0; i < b.N; i++ {
		v, err = CompareNumberFunc(a, d)
		if err != nil {
			b.Fatal("failed")
		}
		if v != Greater {
			b.Fatalf("res is error")
		}
	}
}

func Benchmark_comInt(b *testing.B) {
	var a, d interface{}
	var v CompareType
	var err error
	a = int(3)
	d = int(1)
	for i := 0; i < b.N; i++ {
		v, err = CompareIntFunc(a, d)
		if err != nil {
			b.Fatal("failed")
		}
		if v != Greater {
			b.Fatalf("res is error")
		}
	}
}

func Benchmark_comTime(b *testing.B) {
	var a, d interface{}
	var v CompareType
	var err error
	a = time.Now().Add(time.Hour)
	d = time.Now()
	for i := 0; i < b.N; i++ {
		v, err = CompareTimeFunc(a, d)
		if err != nil {
			b.Fatal("failed", err)
		}
		if v != Greater {
			b.Fatalf("res is error")
		}
	}
}

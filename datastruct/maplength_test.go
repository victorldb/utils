package datastruct

import (
	"fmt"
	"testing"
	"time"
)

func TestSpecLenMP_Set(t *testing.T) {
	now := time.Now()
	nat := NewSpecLenMP(1e5)

	for i := 0; i < 100005; i++ {
		nat.Set(i, i)
	}
	// for i := 0; i < 2000002; i++ {
	// 	nat.Get(uint64(i))
	// }

	fmt.Println(len(nat.mp))
	fmt.Println(time.Since(now))
	v1, ok := nat.Get(30)
	if !ok {
		t.Fatal("30-false")
	}
	v2, ok := nat.Get(100001)
	if !ok {
		t.Fatal("100001-false")
	}
	v3, ok := nat.Get(2)
	if ok {
		t.Fatal("2-false")
	}
	fmt.Println(v1, nat.keys[0], v2, nat.keys[1], v3)
	time.Sleep(time.Hour)
}

func Benchmark_SpecLenMP_set(b *testing.B) {
	nat := NewSpecLenMP(1e5)

	for i := 0; i < b.N; i++ {
		nat.Set(i, i)
	}
}

func Benchmark_SpecLenMP_get(b *testing.B) {
	nat := NewSpecLenMP(1e5)
	var key, value int
	key = 1
	value = 2
	nat.Set(key, value)

	for i := 0; i < b.N; i++ {
		v, ok := nat.Get(1)
		if !ok {
			b.Fatal("failed")
		}
		if nv, nok := v.(int); !nok || nv != value {
			b.Fatal("failed-v")
		}
	}
}

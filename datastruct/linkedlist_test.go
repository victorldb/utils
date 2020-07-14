package datastruct

import (
	"fmt"
	"testing"
)

func TestNewSingleLinkListDel(t *testing.T) {
	l := NewSingleLinkListDel(1)
	dv, ok := l.Set(1)
	fmt.Println(dv, ok)
	dv, ok = l.Set(2)
	fmt.Println(dv, ok)
	dv, ok = l.Set(3)
	fmt.Println(dv, ok)
	v, ok := l.Get()
	fmt.Println(v, ok)
	v, ok = l.Get()
	fmt.Println(v, ok)
}

func Test_set(t *testing.T) {
	l := NewSingleLinkListDel(20)
	for i := 1; i <= 100; i++ {
		v, ok := l.Set(i)
		if ok {
			fmt.Println("del: ", v, ok)
		}
	}
	for {
		v, ok := l.Get()
		if !ok {
			break
		}
		fmt.Println("get: ", v, ok)
	}
}

func Benchmark_set(b *testing.B) {
	l := NewSingleLinkListDel(20)
	for i := 0; i < b.N; i++ {
		l.Set(i)
	}
}

func Benchmark_get(b *testing.B) {
	l := NewSingleLinkListDel(20)
	for i := 0; i < 1000; i++ {
		l.Set(i)
	}
	for i := 0; i < b.N; i++ {
		l.Get()
	}
}

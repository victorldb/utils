package datastore

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func Test_sort(t *testing.T) {
	type sa struct {
		a int
		b string
		c int
	}

	type sc struct {
		a int
		b string
		c int
	}

	var ra interface{}
	var rb interface{}
	var rc interface{}

	ra = sa{a: 2, b: "a", c: 1}
	rb = sc{a: 0, b: "a", c: 2}
	rc = &sa{a: 2, b: "a", c: 1}
	ta := reflect.TypeOf(ra)
	tb := reflect.TypeOf(rb)
	tc := reflect.TypeOf(rc)
	fmt.Printf("%s--%s--%t\n", ta.String(), tb.String(), ta == tb)
	fmt.Println(tc.Kind().String(), tc.String(), tc.Kind() == reflect.Ptr)

	ss := []sa{
		sa{a: 2, b: "a", c: 1},
		sa{a: 2, b: "c", c: 9},
		sa{a: 1, b: "b", c: 2},
		sa{a: 0, b: "a", c: 2},
		sa{a: 3, b: "d", c: 5},
		sa{a: 2, b: "a", c: 10},
		sa{a: 1, b: "e", c: 3},
		sa{a: 0, b: "a", c: 3},
	}

	sort.Slice(ss, func(i, j int) bool {
		if ss[i].a < ss[j].a {
			return true
		}
		if ss[i].a > ss[j].a {
			return false
		}
		if ss[i].b < ss[j].b {
			return true
		}
		if ss[i].b > ss[j].b {
			return false
		}
		return ss[i].c < ss[j].c
	})

	for k, v := range ss {
		fmt.Printf("%d--%+v\n", k, v)
	}
}

func TestDataStore_filter(t *testing.T) {
	ds, err := getTestDataStore()
	if err != nil {
		t.Fatal(err)
	}
	rows, ok := ds.GetAllRow()
	if !ok {
		t.Fatal("GetAllRow failed")
	}
	filters := []Filter{
		Filter{
			Key: "score",
			Dms: []Determine{
				Determine{
					Value:   float64(99),
					ComType: Equal,
				},
			},
		},
		Filter{
			Key: "age",
			Dms: []Determine{
				Determine{
					Value:   int(15),
					ComType: Greater,
				},
			},
		},
	}
	rows, err = ds.filter(rows, filters)
	if err != nil {
		t.Fatal(err)
	}
	if true {
		for _, v := range rows {
			fmt.Println(v.(*DataRow).String())
		}
	}
}

func Benchmark_filterOne(b *testing.B) {
	ds, err := getTestDataStore()
	if err != nil {
		b.Fatal(err)
	}
	rows, ok := ds.GetAllRow()
	if !ok {
		b.Fatal("GetAllRow failed")
	}
	filters := []Filter{
		Filter{
			Key: "score",
			Dms: []Determine{
				Determine{
					Value:   float64(65),
					ComType: GreaterEqual,
				},
			},
		},
	}

	for i := 0; i < b.N; i++ {
		rows, err = ds.filter(rows, filters)
		if err != nil {
			b.Fatal(err)
		}
		if false {
			for _, v := range rows {
				fmt.Println(v.(*DataRow).String())
			}
		}
	}
}

func Benchmark_filterTwo(b *testing.B) {
	ds, err := getTestDataStore()
	if err != nil {
		b.Fatal(err)
	}
	rows, ok := ds.GetAllRow()
	if !ok {
		b.Fatal("GetAllRow failed")
	}
	filters := []Filter{
		Filter{
			Key: "score",
			Dms: []Determine{
				Determine{
					Value:   float64(65),
					ComType: LessEqual,
				},
			},
		},
		Filter{
			Key: "age",
			Dms: []Determine{
				Determine{
					Value:   int(15),
					ComType: Greater,
				},
			},
		},
	}

	for i := 0; i < b.N; i++ {
		rows, err = ds.filter(rows, filters)
		if err != nil {
			b.Fatal(err)
		}
		if false {
			for _, v := range rows {
				fmt.Println(v.(*DataRow).String())
			}
		}
	}
}

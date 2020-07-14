package datastore

import (
	"fmt"
	"testing"
)

func TestDataStore_Group(t *testing.T) {
	ds, err := getTestDataStore()
	if err != nil {
		t.Fatal(err)
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
		Filter{
			Key: "age",
			Dms: []Determine{
				Determine{
					Value:   int(1),
					ComType: Greater,
				},
			},
		},
	}

	calCols := []string{
		"score",
	}

	groupCols := []string{
		"name",
		"phone",
	}
	res, err := ds.Group(calCols, groupCols, filters)
	if err != nil {
		t.Fatal(err)
	}
	if true {
		fmt.Printf("%+v\n", res)
	}
}

func Benchmark_group(b *testing.B) {
	ds, err := getTestDataStore()
	if err != nil {
		b.Fatal(err)
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
		Filter{
			Key: "age",
			Dms: []Determine{
				Determine{
					Value:   int(1),
					ComType: Greater,
				},
			},
		},
	}

	calCols := []string{
		"score",
	}

	groupCols := []string{
		"name",
		"phone",
	}
	for i := 0; i < b.N; i++ {
		res, err := ds.Group(calCols, groupCols, filters)
		if err != nil {
			b.Fatal(err)
		}
		if false {
			fmt.Printf("%+v\n", res)
		}
	}
}

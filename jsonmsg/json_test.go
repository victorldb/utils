package jsonmsg

import (
	"fmt"
	"testing"
)

func Test_Marshal(t *testing.T) {
	type a struct {
		BB string
		CC int
	}
	data := Marshal(StatusOK, "StatusOK", a{BB: "abcdefg", CC: 888})
	fmt.Println(string(data))

	// data = []byte(`{"status":200,"time":"2018-04-12 18:32:43","msg":"StatusOK"}`)
	newA := &a{}
	res, err := UnmarshalWithInt(data, newA)
	if err != nil {
		t.Fatal(err)
	}

	status, ok := res.Status.(int)
	fmt.Println("===============")
	fmt.Printf("%d,%t\n", status, ok)
	fmt.Println("===============")
	fmt.Printf("%+v\n%+v\n", res, *newA)
}

func Test_MarshalList(t *testing.T) {
	data := Marshal(StatusOK, "StatusOK", []int{1, 2, 3, 4, 5, 6, 7, 8})
	fmt.Println(string(data))
}

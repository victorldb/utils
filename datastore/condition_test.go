package datastore

import (
	"fmt"
	"testing"
)

func Test_dsCondition_cloumnsName(t *testing.T) {
	dsc := NewDsCondition().Greater("a", 1).And().Less("b", 100).Or().GreaterEqual("c", 80)
	dsc2 := NewDsCondition().Greater("m", 1).And().Less("n", 100).Or().GreaterEqual("h", 80)
	dsc3 := NewDsCondition().Greater("o", 1).And().Less("p", 100).Or().GreaterEqual("k", 80)
	nds := dsc.JoinWithAnd(dsc2).JoinWithOr(dsc3)
	colsName := nds.getCloumnsName()
	fmt.Printf("%+v\n", colsName)
}

func Test_dsCondition_decode(t *testing.T) {
	dsc := NewDsCondition().Greater("c1", 1).And().Less("c2", 100).Or().GreaterEqual("c3", 80)
	dscInner := NewDsCondition().Greater("c4", 1).And().Less("c5", 100).Or().GreaterEqual("c6", 80)
	dscInner2 := NewDsCondition().Greater("c7", 1).And().Less("c8", 100).Or().GreaterEqual("c9", 80).InnerWithOr(dscInner)
	nds := dsc.JoinWithOr(dscInner2)
	decodeStr := nds.String()
	if true {
		fmt.Println(decodeStr)
	}
}

func Benchmark__dsCondition_decode_string(b *testing.B) {
	dsc := NewDsCondition().Greater("c1", 1).And().Less("c2", 100).Or().GreaterEqual("c3", 80)
	dscInner := NewDsCondition().Greater("c4", 1).And().Less("c5", 100).Or().GreaterEqual("c6", 80)
	dscInner2 := NewDsCondition().Greater("c7", 1).And().Less("c8", 100).Or().GreaterEqual("c9", 80).InnerWithOr(dscInner)
	dscInner3 := NewDsCondition().Greater("c7", 1).And().Less("c8", 100).Or().GreaterEqual("c9", 80).InnerWithOr(dscInner2)
	dscInner5 := NewDsCondition().Greater("c7", 1).And().Less("c8", 100).Or().GreaterEqual("c9", 80).InnerWithOr(dscInner3)
	dscInner6 := NewDsCondition().Greater("c7", 1).And().Less("c8", 100).Or().GreaterEqual("c9", 80).InnerWithOr(dscInner5)
	dscInner7 := NewDsCondition().Greater("c7", 1).And().Less("c8", 100).Or().GreaterEqual("c9", 80).InnerWithOr(dscInner6)
	dscInner8 := NewDsCondition().Greater("c7", 1).And().Less("c8", 100).Or().GreaterEqual("c9", 80).InnerWithOr(dscInner7)
	dscInner9 := NewDsCondition().Greater("c7", 1).And().Less("c8", 100).Or().GreaterEqual("c9", 80).InnerWithOr(dscInner8)
	dscInner10 := NewDsCondition().Greater("c7", 1).And().Less("c8", 100).Or().GreaterEqual("c9", 80).InnerWithOr(dscInner9)
	dscInner11 := NewDsCondition().Greater("c7", 1).And().Less("c8", 100).Or().GreaterEqual("c9", 80).InnerWithOr(dscInner10)
	dscInner12 := NewDsCondition().Greater("c7", 1).And().Less("c8", 100).Or().GreaterEqual("c9", 80).InnerWithOr(dscInner11)
	dscInner13 := NewDsCondition().Greater("c7", 1).And().Less("c8", 100).Or().GreaterEqual("c9", 80).InnerWithOr(dscInner12)
	dsc2 := NewDsCondition().Greater("c10", 1).And().Less("c11", 100).Or().GreaterEqual("c12", 80).InnerWithOr(dscInner13)
	nds := dsc.JoinWithOr(dsc2)
	for i := 0; i < b.N; i++ {
		decodeStr := nds.String()
		if false {
			fmt.Println(decodeStr)
		}
	}
}

func Benchmark__dsCondition_decode_slice(b *testing.B) {
	dsc := NewDsCondition().Greater("c1", 1).And().Less("c2", 100).Or().GreaterEqual("c3", 80)
	dscInner := NewDsCondition().Greater("c4", 1).And().Less("c5", 100).Or().GreaterEqual("c6", 80)
	dscInner2 := NewDsCondition().Greater("c7", 1).And().Less("c8", 100).Or().GreaterEqual("c9", 80).InnerWithOr(dscInner)
	dscInner3 := NewDsCondition().Greater("c7", 1).And().Less("c8", 100).Or().GreaterEqual("c9", 80).InnerWithOr(dscInner2)
	dscInner5 := NewDsCondition().Greater("c7", 1).And().Less("c8", 100).Or().GreaterEqual("c9", 80).InnerWithOr(dscInner3)
	dscInner6 := NewDsCondition().Greater("c7", 1).And().Less("c8", 100).Or().GreaterEqual("c9", 80).InnerWithOr(dscInner5)
	dscInner7 := NewDsCondition().Greater("c7", 1).And().Less("c8", 100).Or().GreaterEqual("c9", 80).InnerWithOr(dscInner6)
	dscInner8 := NewDsCondition().Greater("c7", 1).And().Less("c8", 100).Or().GreaterEqual("c9", 80).InnerWithOr(dscInner7)
	dscInner9 := NewDsCondition().Greater("c7", 1).And().Less("c8", 100).Or().GreaterEqual("c9", 80).InnerWithOr(dscInner8)
	dscInner10 := NewDsCondition().Greater("c7", 1).And().Less("c8", 100).Or().GreaterEqual("c9", 80).InnerWithOr(dscInner9)
	dscInner11 := NewDsCondition().Greater("c7", 1).And().Less("c8", 100).Or().GreaterEqual("c9", 80).InnerWithOr(dscInner10)
	dscInner12 := NewDsCondition().Greater("c7", 1).And().Less("c8", 100).Or().GreaterEqual("c9", 80).InnerWithOr(dscInner11)
	dscInner13 := NewDsCondition().Greater("c7", 1).And().Less("c8", 100).Or().GreaterEqual("c9", 80).InnerWithOr(dscInner12)
	dsc2 := NewDsCondition().Greater("c10", 1).And().Less("c11", 100).Or().GreaterEqual("c12", 80).InnerWithOr(dscInner13)
	nds := dsc.JoinWithOr(dsc2)

	for i := 0; i < b.N; i++ {
		decodeStr := nds.StringWithSlice()
		if false {
			fmt.Println(decodeStr)
		}
	}
}

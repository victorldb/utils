package utils

import (
	"fmt"
	"testing"
)

func TestRemoveTextRemark(t *testing.T) {
	res, err := RemoveTextRemark(testText)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n%s\n", testText, res)
}

var testText = `
A
//AA
// AA
 //AA
 // AA
    //AA
    // AA
     //AA
     // AA
      //AA
      // AA
B
/*
BB
*/
C
/*
// CC
*/
D
/*
DD
/*
DD
*/
*/
E
/*
EE
/*
// EE
EE
*/
*/
SUCCESS
`

package converttofile

import (
	"fmt"
	"io/ioutil"
	"testing"
	"time"
)

func TestConvertToXlsx(t *testing.T) {
	data := ConToXlsx{
		Header: []interface{}{
			"name",
			"年龄",
			"备注",
		},
		Rows: [][]interface{}{
			[]interface{}{
				"N1",
				23,
				"VVVVVV",
			},
			[]interface{}{
				"N2",
				28,
				"VVVVVV",
			},
			[]interface{}{
				"N3",
				26,
				"VVVVVV",
			},
		},
		SheetName: "sh1",
	}
	buff, err := ConvertToXlsx(data)
	if err != nil {
		t.Fatal(err)
	}
	xlsxData := buff.Bytes()
	err = ioutil.WriteFile("/home/victor/tmp/1.xlsx", xlsxData, 0755)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_getColumnsNameWithLength(t *testing.T) {
	now := time.Now()
	res := getAllColNameWithLength(1000)
	println(time.Since(now).String())
	for k, v := range res {
		fmt.Printf("%5d : %s\n", k, v)
	}
}

func Test_getColumnNameWithLength(t *testing.T) {
	for i := 1; i <= 520; i++ {
		res := getColNameWithID(i)
		fmt.Println(i, res)
	}
}

func Benchmark_getColumnsNameWithLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getAllColNameWithLength(i)
	}
}

func Benchmark_getColumnNameWithLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getColNameWithID(i)
	}
}

func Benchmark_getIDWithTitle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		colNameToID("AZIURIWQUIQWEWQUOQUWA")
	}
}

func Test_titleToID(t *testing.T) {
	res := getAllColNameWithLength(888)
	for k, v := range res {
		res, err := colNameToID(v)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("%d == %d == %s\n", k+1, res, v)
	}
}

func Benchmark_pp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decimalToAny(i, 26)
	}
}

func TestPrint(t *testing.T) {
	x := decimalToAny(27, 26)
	flag := false
	var y []byte
	for i := range x {
		if !flag {
			if x[i] == '0' {
				y = append(y, 'Z')
				flag = true
			} else {
				if flag {
					if i == len(x)-1 && byte(x[i]) == '1' {
						break
					}
					if byte(x[i]) == '1' {
						y = append(y, 'Z')
					} else {
						y = append(y, twentySix[x[i]])
						flag = false
					}
				} else {
					y = append(y, twentySix[x[i]])
				}
			}
		} else {
			if i == len(x)-1 && byte(x[i]) == '1' {
				break
			}
			if byte(x[i]) == '1' {
				y = append(y, 'Z')
			} else {
				if byte(x[i]) == '0' {
					y = append(y, 'Y')
				} else {
					y = append(y, twentySix[x[i-1]])
					flag = false
				}
			}
		}
	}
	for i := len(y) - 1; i >= 0; i-- {
		fmt.Print(string(y[i]))
	}
	fmt.Println()
}

var tenToAny = map[int]byte{0: '0', 1: '1', 2: '2', 3: '3', 4: '4', 5: '5', 6: '6', 7: '7', 8: '8', 9: '9', 10: 'a', 11: 'b', 12: 'c', 13: 'd', 14: 'e', 15: 'f', 16: 'g', 17: 'h', 18: 'i', 19: 'j', 20: 'k', 21: 'l', 22: 'm', 23: 'n', 24: 'o', 25: 'p', 26: 'q', 27: 'r', 28: 's', 29: 't', 30: 'u', 31: 'v', 32: 'w', 33: 'x', 34: 'y', 35: 'z', 36: ':', 37: ';', 38: '<', 39: '=', 40: '>', 41: '?', 42: '@', 43: '[', 44: ']', 45: '^', 46: '_', 47: '{', 48: '|', 49: '}', 50: 'A', 51: 'B', 52: 'C', 53: 'D', 54: 'E', 55: 'F', 56: 'G', 57: 'H', 58: 'I', 59: 'J', 60: 'K', 61: 'L', 62: 'M', 63: 'N', 64: 'O', 65: 'P', 66: 'Q', 67: 'R', 68: 'S', 69: 'T', 70: 'U', 71: 'V', 72: 'W', 73: 'X', 74: 'Y', 75: 'Z'}
var twentySix = map[byte]byte{'1': 'A', '2': 'B', '3': 'C', '4': 'D', '5': 'E', '6': 'F', '7': 'G', '8': 'H', '9': 'I', 'a': 'J', 'b': 'K', 'c': 'L', 'd': 'M', 'e': 'N', 'f': 'O', 'g': 'P', 'h': 'Q', 'i': 'R', 'j': 'S', 'k': 'T', 'l': 'U', 'm': 'V', 'n': 'W', 'o': 'X', 'p': 'Y', 'q': 'Z'}

func decimalToAny(num, n int) []byte {
	var newNumStr []byte
	var remainder int
	var remainderString byte
	for num != 0 {
		remainder = num % n
		remainderString = tenToAny[remainder]
		newNumStr = append(newNumStr, remainderString)

		num = num / n
	}
	return newNumStr
}

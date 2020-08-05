package uid

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

//NewID --
func NewID() string {
	uid4, err := uuid.NewV4()
	if err != nil {
		return getRandomCode()
	}
	return strings.ReplaceAll(uid4.String(), "-", "")
}

//Md5 --
func Md5(text string) string {
	hashMd5 := md5.New()
	io.WriteString(hashMd5, text)
	return fmt.Sprintf("%x", hashMd5.Sum(nil))
}

func getRandomCode() string {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(time.Now().UnixNano()))
	b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7] = b[7], b[6], b[5], b[4], b[3], b[2], b[1], b[0]
	return uintToString(binary.BigEndian.Uint64(b), 62)
}

var uintMpChar map[uint64]string = map[uint64]string{
	0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "a",
	11: "b", 12: "c", 13: "d", 14: "e", 15: "f", 16: "g", 17: "h", 18: "i", 19: "j", 20: "k",
	21: "l", 22: "m", 23: "n", 24: "o", 25: "p", 26: "q", 27: "r", 28: "s", 29: "t", 30: "u",
	31: "v", 32: "w", 33: "x", 34: "y", 35: "z", 36: "A", 37: "B", 38: "C", 39: "D", 40: "E",
	41: "F", 42: "G", 43: "H", 44: "I", 45: "J", 46: "K", 47: "L", 48: "M", 49: "N", 50: "O",
	51: "P", 52: "Q", 53: "R", 54: "S", 55: "T", 56: "U", 57: "V", 58: "W", 59: "X", 60: "Y", 61: "Z",
}

// 10进制转任意进制
func uintToString(num, n uint64) string {
	newNumStr := ""
	var remainder uint64
	var remainderString string
	for num != 0 {
		remainder = num % n
		if 62 > remainder && remainder > 9 {
			remainderString = uintMpChar[remainder]
		} else {
			remainderString = strconv.Itoa(int(remainder))
		}
		newNumStr = remainderString + newNumStr
		num = num / n
	}
	return newNumStr
}

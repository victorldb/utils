package speedbufio

import (
	"bytes"
	"fmt"
	"io"
	"testing"
	"time"
)

//go test -v -test.run TestSpeedR

func TestSpeedR(t *testing.T) {
	//make 100MB data
	var err error
	var n int
	length := 100 * 1024 * 1024
	data := make([]byte, length)
	for i := 0; i < length; i++ {
		if i > length-10 {
			data[i] = 66
			continue
		}
		data[i] = 65
	}
	r := bytes.NewReader(data)

	sr, err := NewSpeedReader("20MB", r, 10)
	if err != nil {
		t.Fatal(err)
	}

	tmpData := make([]byte, 5*1024)
	add := 0
	count := 0
	lastTime := time.Now()
	for {
		count++
		n, err = sr.Read(tmpData)
		add += n
		// fmt.Printf("%s %2fMB %d\n", time.Now().Sub(lastTime).String(), float64(add)/(1024*1024), add)
		if err != nil {
			fmt.Println(err)
			break
		}
		if add >= length {
			break
		}

	}
	fmt.Println(time.Now().Sub(lastTime).String())
	fmt.Println(count)
	if n > 20 {
		dataTmp := tmpData[:n]
		fmt.Println(string(dataTmp[len(dataTmp)-20:]))
	}
	if time.Now().Sub(lastTime) < time.Duration(length/(20*1024*1024)) {
		t.Fatal("time is less")
	}

}

//go test -v -test.run TestSpeedW
func TestSpeedW(t *testing.T) {
	var err error
	buff := bytes.NewBuffer(make([]byte, 0))
	length := 100 * 1024 * 1024
	data := make([]byte, length)
	for i := 0; i < length; i++ {
		if i > length-10 {
			data[i] = 66
			continue
		}
		data[i] = 65
	}

	sr, err := NewSpeedWriter("20MB", buff, 10)
	if err != nil {
		t.Fatal(err)
	}

	lastTime := time.Now()
	n, err := sr.Write(data)
	fmt.Println(time.Now().Sub(lastTime).String())
	if err != nil {
		t.Fatal(err)
	}

	if n != length {
		t.Fatal("n not eq length:", n, length)
	}
	if buff.Len() != length {
		t.Fatal("n not eq length:", buff.Len(), length)
	}
	buff.Next(length - 20)
	tmpBuff := make([]byte, 100)

	fmt.Println(buff.Len())
	n, err = buff.Read(tmpBuff)
	if err != nil && err == io.EOF {
		t.Fatal(err)
	}
	fmt.Println(string(tmpBuff[:n]), n)

}

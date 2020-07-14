package readfileline

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/victorldb/utils/file"
)

func TestNewReadLog(t *testing.T) {
	var err error
	fileDir := fmt.Sprintf("/tmp/%d", time.Now().UnixNano())
	fileName := fmt.Sprintf("%s/test%d.data", fileDir, time.Now().UnixNano())
	if file.FileExists(fileDir) {
		err = os.RemoveAll(fileDir)
		if err != nil {
			t.Fatal(err)
		}
	}

	err = os.MkdirAll(fileDir, 0755)
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		err = os.RemoveAll(fileDir)
		if err != nil {
			t.Fatal(err)
		}
	}()

	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	go func() {
		for i := 1; i <= 100; i++ {
			fmt.Fprintf(f, "TEST READ LI")
			<-time.After(time.Second)
			fmt.Fprintf(f, "NE:%d\n", i)
		}
	}()

	cfg := DirFileConfig{
		Dir:                 fileDir,
		DataType:            1,
		FileLike:            "*.data",
		OffsetName:          "",
		OffsetSize:          0,
		RmLineFlag:          true,
		ContentBaseNameFlag: true,
	}
	dataChan := make(chan *ContentType, 100)
	rd, err := NewReadLine(cfg, dataChan)
	if err != nil {
		t.Fatal(err)
	}
	go rd.Start()
	defer rd.Close()

	timer := time.NewTimer(2 * time.Second)
	checkCount := 0
lableExit:
	for {
		timer.Reset(2 * time.Second)
		select {
		case data := <-dataChan:
			if data != nil {
				checkCount++
				fmt.Println(string(data.Data), checkCount, data.Name, data.Offset)
			}
		case <-timer.C:
			break lableExit
		}
	}
}

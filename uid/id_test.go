package uid

import (
	"testing"
	"time"

	"golang.org/x/exp/errors/fmt"
)

func TestNewID(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(NewID(), time.Now().UnixNano())
	}
}

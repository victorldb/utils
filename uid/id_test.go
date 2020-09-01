package uid

import (
	"fmt"
	"testing"
	"time"
)

func Test_NewID(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(NewID(), time.Now().UnixNano())
	}
}

func Test_uintToString(t *testing.T) {
	type args struct {
		num uint64
		n   uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uintToString(tt.args.num, tt.args.n); got != tt.want {
				t.Errorf("uintToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

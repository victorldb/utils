package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestNextStartDay(t *testing.T) {
	offsetTime := 3 * time.Hour
	now := time.Now().Add(0 * time.Hour)
	tt := NextStartDay(now.Add(offsetTime)).Add(-offsetTime)
	fmt.Printf("%s\n", tt.Format("2006-01-02 15:04:05"))
}

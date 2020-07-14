package speedbufio

import (
	"runtime"
	"testing"
	"time"
)

func TestRateLimitPerSec_Check(t *testing.T) {
	ra := NewRateLimitPerSec(1000000)
	cpuNu := runtime.NumCPU()
	for i := 0; i < cpuNu; i++ {
		go func() {
			for {
				ra.Check()
			}
		}()
	}

	time.Sleep(time.Hour)

}

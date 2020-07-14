package speedbufio

import (
	"sync"
	"time"
)

// RateLimitPerSec --
type RateLimitPerSec struct {
	unitTime  time.Duration
	rate      uint64
	countRate uint64
	syncRa    sync.Mutex
}

// NewRateLimitPerSec --
func NewRateLimitPerSec(rate uint64) (raLimit *RateLimitPerSec) {
	raLimit = &RateLimitPerSec{
		rate:     rate,
		unitTime: time.Second,
	}

	go raLimit.run()
	return raLimit
}

// NewRateLimit --
func NewRateLimit(unitTime time.Duration, rate uint64) (raLimit *RateLimitPerSec) {
	raLimit = &RateLimitPerSec{
		rate:     rate,
		unitTime: unitTime,
	}

	go raLimit.run()
	return raLimit
}

func (c *RateLimitPerSec) run() {
	ticker := time.NewTicker(c.unitTime)
	for {
		<-ticker.C
		c.syncRa.Lock()
		c.countRate = 0
		c.syncRa.Unlock()
	}
}

// func (c *RateLimitPerSec) run() {
// 	ticker := time.NewTicker(time.Second)
// 	for {
// 		<-ticker.C
// 		println(c.countRate)
// 		atomic.StoreUint64(&c.countRate, 0)
// 	}
// }

// Check --
func (c *RateLimitPerSec) Check() (ok bool) {
	c.syncRa.Lock()
	if c.countRate < c.rate {
		ok = true
		c.countRate++
	}
	c.syncRa.Unlock()
	return ok
}

// CheckWait --
func (c *RateLimitPerSec) CheckWait() {
	for {
		if c.Check() {
			break
		}
		<-time.After(50 * time.Microsecond)
	}
}

// CheckWaitTime --
func (c *RateLimitPerSec) CheckWaitTime(t time.Duration) {
	if t == 0 {
		t = 50 * time.Microsecond
	}
	for {
		if c.Check() {
			break
		}
		<-time.After(t)
	}
}

// Check --
// func (c *RateLimitPerSec) Check() (ok bool) {
// 	count := c.countRate
// 	if count < c.rate {
// 		ok = atomic.CompareAndSwapUint64(&c.countRate, count, count+1)
// 	}
// 	return ok
// }

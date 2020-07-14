package speedbufio

import (
	"fmt"
	"log"
	"math"
	"sync"
	"time"

	"github.com/victorldb/utils/callbackmsg"
)

//SpeedStatus --
type SpeedStatus struct {
	startTime time.Time
	flow      int64
	flowAll   int64
	syncFlow  sync.Mutex

	cbm callbackmsg.CallbackMsg
}

// NewSpeedCheck --
func NewSpeedCheck(startTime time.Time, cbm callbackmsg.CallbackMsg) (speed *SpeedStatus) {
	if startTime.IsZero() {
		startTime = time.Now()
	}
	speed = &SpeedStatus{}
	speed.startTime = startTime
	if cbm != nil {
		speed.cbm = cbm
	}

	go func() {
		timer := time.NewTimer(time.Second)
		for {
			timer.Reset(time.Second)
			select {
			case <-timer.C:
				data := speed.getSpeedStatus()
				if speed.cbm != nil {
					speed.cbm.RegInfo("%s", data)
				} else {
					log.Printf("%s\n", data)
				}
			}
		}
	}()

	return speed
}

//SetCBM --
func (c *SpeedStatus) SetCBM(cbm callbackmsg.CallbackMsg) {
	c.cbm = cbm
}

//SetFlow --
func (c *SpeedStatus) SetFlow(flow int64) {
	c.syncFlow.Lock()
	c.flow += flow
	c.flowAll += flow
	c.syncFlow.Unlock()
}

//GetSpeed --
func (c *SpeedStatus) getSpeedStatus() string {
	c.syncFlow.Lock()
	now := time.Now()
	flow := c.flow
	allFlow := c.flowAll
	startTime := c.startTime

	c.flow = 0
	c.startTime = now
	c.syncFlow.Unlock()

	subTime := int64(now.Sub(startTime) / 1e9)
	avpFlowSecond := flow / subTime
	value, suffix := conventFlowToVS(avpFlowSecond)
	allValue, allSuffix := conventFlowToVS(allFlow)
	return fmt.Sprintf("speed:%0.2f%s rec:%0.2f%s byte:%d", value, suffix, allValue, allSuffix, allFlow)
}

func conventFlowToVS(v int64) (value float64, suffix string) {
	switch {
	case v < 1<<20:
		percentFloat64 := float64(v) / float64(1<<10)
		value = math.Trunc(percentFloat64*1e2+0.5) / 1e2
		suffix = "KB"
	case v < 1<<30:
		percentFloat64 := float64(v) / float64(1<<20)
		value = math.Trunc(percentFloat64*1e2+0.5) / 1e2
		suffix = "MB"
	case v < 1<<40:
		percentFloat64 := float64(v) / float64(1<<30)
		value = math.Trunc(percentFloat64*1e2+0.5) / 1e2
		suffix = "GB"
	case v >= 1<<40:
		percentFloat64 := float64(v) / float64(1<<40)
		value = math.Trunc(percentFloat64*1e2+0.5) / 1e2
		suffix = "TB"
	default:
		value = float64(v)
		suffix = "Byte"
	}
	return value, suffix
}

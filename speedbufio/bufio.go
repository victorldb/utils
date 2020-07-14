package speedbufio

import (
	"errors"
	"io"
	"strconv"
	"strings"
	"time"
)

var (
	//ErrSpeedFormat --并发不安全
	ErrSpeedFormat = errors.New("speed format is error")
)

//SpeedRW --
type SpeedRW struct {
	//byte
	unit     int
	per      int
	r        io.Reader
	w        io.Writer
	count    int
	timer    *time.Timer
	buffData []byte
}

//NewSpeedWriter --
func NewSpeedWriter(speed string, w io.Writer, per int) (*SpeedRW, error) {
	sw := &SpeedRW{}
	sw.setPer(per)
	err := sw.setSpeed(speed)
	if err != nil {
		return nil, err
	}
	sw.setWriter(w)
	return sw, nil
}

//NewSpeedReader --
func NewSpeedReader(speed string, r io.Reader, per int) (*SpeedRW, error) {
	sw := &SpeedRW{}
	sw.setPer(per)
	err := sw.setSpeed(speed)
	if err != nil {
		return nil, err
	}
	sw.setReader(r)
	return sw, nil
}

//SetSpeed 设置传输速度
func (c *SpeedRW) setSpeed(speed string) error {
	speedLength := len(speed)
	if speedLength <= 2 {
		return errors.New("speed format is error:length")
	}
	unit, err := strconv.Atoi(speed[:speedLength-2])
	if err != nil {
		return errors.New("speed format is error:unit")
	}

	//MB KB GB /s
	unitStr := strings.ToUpper(speed[speedLength-2:])
	switch unitStr {
	case "KB":
		c.unit = unit * 1024
	case "MB":
		c.unit = unit * 1024 * 1024
	case "GB":
		c.unit = unit * 1024 * 1024 * 1024
	default:
		return errors.New("speed format is error:unitStr")
	}
	if c.per < 1 {
		c.per = 1
	}
	c.unit = c.unit / c.per
	c.buffData = make([]byte, 10*1024)
	c.timer = time.NewTimer(time.Second / time.Duration(c.per))
	return nil
}

//SetReader --
func (c *SpeedRW) setReader(r io.Reader) {
	c.r = r
}

//SetWriter --
func (c *SpeedRW) setWriter(w io.Writer) {
	c.w = w
}

//SetPer --
func (c *SpeedRW) setPer(per int) {
	c.per = per
	if c.per < 1 {
		c.per = 1
	}
	c.unit = c.unit / c.per
	c.timer = time.NewTimer(time.Second / time.Duration(c.per))
}

//Read --
func (c *SpeedRW) Read(data []byte) (int, error) {
	if c.r == nil {
		return 0, errors.New("reader is nil")
	}
	pd := false
	tmpData := make([]byte, 0)
	if len(c.buffData) >= len(data) {
		tmpData = data
		pd = true
	} else {
		tmpData = c.buffData
	}
	less := c.count + len(tmpData) - c.unit
	if less > 0 && c.count < c.unit {
		tmpData = tmpData[:len(tmpData)-less]
	}
	n, err := c.r.Read(tmpData)
	if n > 0 && !pd {
		copy(data, tmpData[:n])
	}
	c.count += n
	if c.count >= c.unit {
		<-c.timer.C
		c.count = 0
		c.timer.Reset(time.Second / time.Duration(c.per))
	}
	return n, err
}

func (c *SpeedRW) Write(data []byte) (int, error) {
	if c.w == nil {
		return 0, errors.New("writer is nil")
	}
	add := 0
	pd := false
	end := 0
	tmpData := make([]byte, 0)
	for {
		dataLength := len(data)
		if dataLength <= c.unit {
			end = dataLength
		} else {
			end = c.unit
		}
		less := c.count + end - c.unit
		if less > 0 {
			end = c.unit - c.count
		}

		tmpData = data[:end]
		if len(data) <= end {
			pd = true
		} else {
			data = data[end:]
		}

		n, err := c.w.Write(tmpData)
		add += n
		c.count += n
		if err != nil {
			return add, err
		}
		if c.count >= c.unit {
			<-c.timer.C
			c.count = 0
			c.timer.Reset(time.Second / time.Duration(c.per))
		}
		if pd {
			break
		}
	}
	return add, nil
}

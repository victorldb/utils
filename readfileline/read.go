package readfileline

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/victorldb/utils/callbackmsg"
	"github.com/victorldb/utils/directoryfile"
)

const (
	defaultEOFWaitTime  = 10 * time.Millisecond
	defaultMaxWaitCount = 80
)

// ReadFileLine --
type ReadFileLine struct {
	cfg            DirFileConfig
	dataChan       chan *ContentType
	dires          *directoryfile.Directs
	lastFileNameID directoryfile.FileNameID
	offset         int64

	finishTag bool
	finished  chan int

	isRmLine          bool
	isContentBaseName bool
	f                 *os.File
	br                *bufio.Reader
	cbm               callbackmsg.CallbackMsg
	sync.RWMutex
}

// ContentType --
type ContentType struct {
	Data     []byte
	DataType int
	Offset   int64
	Name     string
}

// DirFileConfig --
type DirFileConfig struct {
	//监控的目录
	Dir string

	//日志类型
	DataType int

	//匹配名称
	FileLike string

	// OffsetName为baseName
	OffsetName string
	OffsetSize int64

	RmLineFlag          bool
	ContentBaseNameFlag bool

	CallBM callbackmsg.CallbackMsg
}

// NewReadLine --
func NewReadLine(cfg DirFileConfig, chanData chan *ContentType) (rl *ReadFileLine, err error) {
	rl = &ReadFileLine{
		cfg:               cfg,
		dataChan:          chanData,
		isRmLine:          cfg.RmLineFlag,
		isContentBaseName: cfg.ContentBaseNameFlag,
		lastFileNameID:    directoryfile.FileNameID{Name: cfg.OffsetName},
		cbm:               cfg.CallBM,
	}
	rl.dires, err = directoryfile.NewDirects(cfg.Dir, cfg.FileLike, cfg.OffsetName)
	if err != nil {
		return nil, err
	}
	if cfg.OffsetName != "" {
		rl.lastFileNameID, err = directoryfile.GetNameIDByFile(filepath.Join(cfg.Dir, cfg.OffsetName))
		if err == nil {
			rl.offset = cfg.OffsetSize
		}
	}

	rl.finished = make(chan int)
	return rl, nil
}

//Start --
func (c *ReadFileLine) Start() {
	var err error

	sleepTag := false

	for {
		if c.checkFinish() {
			break
		}
		if sleepTag {
			time.Sleep(2 * time.Second)
		} else {
			sleepTag = true
		}

		_, err = c.setBufioReader()
		if err != nil {
			c.logError("setBufioReader:err=%s,dir=%s", err.Error(), c.cfg.Dir)
			continue
		}
		err = c.read()
		if err != nil {
			c.logError("read:err=%s,dir=%s", err.Error(), c.cfg.Dir)
			continue
		}
	}
	c.finished <- 1
	close(c.dataChan)
	return
}

func (c *ReadFileLine) setOffset(i int64) {
	c.Lock()
	c.offset = i
	c.Unlock()
}

func (c *ReadFileLine) getOffset() int64 {
	c.RLock()
	i := c.offset
	c.RUnlock()
	return i
}

func (c *ReadFileLine) setOffsetAdd(i int64) {
	c.Lock()
	c.offset += i
	c.Unlock()
}

func (c *ReadFileLine) getStatus() (string, int64) {
	c.RLock()
	name := c.lastFileNameID.Name
	offset := c.offset
	c.RUnlock()
	return name, offset
}

//Close --
func (c *ReadFileLine) Close() {
	c.Lock()
	c.finishTag = true
	if c.f != nil {
		c.f.Close()
	}
	c.Unlock()

	timer := time.NewTimer(1 * time.Second)
	select {
	case <-c.dataChan:
	case <-timer.C:
	}
	<-c.finished
}

func (c *ReadFileLine) checkFinish() (ok bool) {
	c.RLock()
	ok = c.finishTag
	c.RUnlock()
	return ok
}

func (c *ReadFileLine) getCurrentName() (name string) {
	c.RLock()
	name = c.lastFileNameID.Name
	c.RUnlock()
	return name
}

func (c *ReadFileLine) setBufioReader() (isNewFile bool, err error) {
	nameID, err := c.dires.GetNextFileName()
	if err != nil {
		return false, err
	}

	if c.lastFileNameID.Equle(nameID) && c.f != nil && c.br != nil {
		return false, nil
	}

	if !c.lastFileNameID.Equle(nameID) {
		if c.f != nil {
			c.f.Close()
		}
		c.f, err = os.Open(nameID.Name)
		if err != nil {
			return false, err
		}
		if c.br == nil {
			c.br = bufio.NewReader(c.f)
		} else {
			c.br.Reset(c.f)
		}
		c.lastFileNameID = nameID
		c.setOffset(0)
		isNewFile = true
	}

	// offset 处理
	if c.lastFileNameID.Equle(nameID) && c.f == nil && c.br == nil {
		c.f, err = os.Open(nameID.Name)
		if err != nil {
			return false, err
		}
		_, err := c.f.Seek(c.offset, 0)
		if err != nil {
			return false, err
		}
		c.br = bufio.NewReader(c.f)
	}
	return isNewFile, nil
}

func (c *ReadFileLine) read() (err error) {
	// panic检查
	defer func() {
		err := recover()
		if err != nil {
			c.logInfo("redline recover: %s\n", c.cfg.Dir)
		}
	}()

	// 每次读取内容的长度
	var lengthContent int
	// 读取的内容
	var data []byte
	// 文件结尾检查多少次后，再去调用下一层接口来获取是否有新文件。主要是为了避免产生过多的系统调用，尤其是文件数量非常多的时候
	var waitCount int
	// 是否是新文件
	var isNewFile bool
	// 不满足一行数据的缓存
	tmpData := make([]byte, 0)

	for {
		// 检查是否退出
		if c.checkFinish() {
			break
		}
		data, err = c.br.ReadBytes('\n')
		lengthContent = len(data)
		if lengthContent > 0 {
			waitCount = 0
			c.setOffsetAdd(int64(lengthContent))
		} else {
			waitCount++
		}

		if err != nil {
			// 文件异常退出
			if err != io.EOF {
				break
			}

			<-time.After(defaultEOFWaitTime)

			if lengthContent == 0 {
				if waitCount > defaultMaxWaitCount || !c.dires.CheckIsLastFile(filepath.Base(c.getCurrentName())) {
					isNewFile, err = c.setBufioReader()
					if err != nil {
						break
					}
					waitCount = 0

					// 有新文件产生了还未读到\n就丢弃tmpData里面的无效数据
					if isNewFile {
						tmpData = make([]byte, 0)
					}
				}
				continue
			} else if lengthContent > 0 {
				tmpData = append(tmpData, data...)
				continue
			}
		}

		if len(tmpData) > 0 {
			data = bytes.Join([][]byte{tmpData, data}, nil)
			tmpData = make([]byte, 0)
		}

		newLengthContent := len(data)
		if newLengthContent > 0 {
			name, offset := c.getStatus()

			contentName := ""
			if c.isContentBaseName {
				contentName = filepath.Base(name)
			} else {
				contentName = name
			}

			if c.isRmLine {
				for {
					if length := len(data); length > 0 && (data[length-1] == '\r' || data[length-1] == '\n') {
						data = data[:length-1]
						continue
					}
					break
				}
				c.dataChan <- &ContentType{Data: data[:newLengthContent-1], DataType: c.cfg.DataType, Name: contentName, Offset: offset}
			} else {
				c.dataChan <- &ContentType{Data: data, DataType: c.cfg.DataType, Name: contentName, Offset: offset}
			}
		}
	}
	if err == io.EOF {
		err = nil
	}
	return err
}

func (c *ReadFileLine) logError(format string, v ...interface{}) {
	if c.cbm != nil {
		c.cbm.RegError(format, v...)
	} else {
		log.Printf(format, v...)
	}
}

func (c *ReadFileLine) logInfo(format string, v ...interface{}) {
	if c.cbm != nil {
		c.cbm.RegInfo(format, v...)
	} else {
		log.Printf(format, v...)
	}
}

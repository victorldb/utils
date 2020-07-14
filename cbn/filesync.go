package cbn

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"sync"
	"time"
)

type chanName struct {
	ch     chan int
	name   string
	dir    string
	isWait bool
}

// SyncDirFile --
type SyncDirFile struct {
	cn         []*chanName
	length     int
	dirMpIndex map[string]int

	fileLikeSlice []string
	fileLike      string

	getCompareNameFunc func(name string) string

	syncCBN sync.RWMutex
}

// NewCBN --
func NewCBN(dirs []string, fileLike string, getCompareName func(name string) string) (newcb *SyncDirFile, err error) {
	if len(dirs) < 2 {
		return nil, fmt.Errorf("dirs length must then 2")
	}
	if fileLike == "" {
		return nil, fmt.Errorf("fileLike is empty")
	}
	if getCompareName == nil {
		return nil, fmt.Errorf("getCompareName is nil")
	}

	newcb = &SyncDirFile{
		dirMpIndex:         make(map[string]int, 0),
		fileLikeSlice:      make([]string, 0),
		getCompareNameFunc: getCompareName,
	}

	sp1 := strings.Split(fileLike, "*")
	for _, v := range sp1 {
		if v != "" {
			newcb.fileLikeSlice = append(newcb.fileLikeSlice, v)
		}
	}
	newcb.fileLike = fileLike

	newcb.cn = make([]*chanName, len(dirs))
	newcb.length = len(dirs)
	for k, dir := range dirs {
		newcb.dirMpIndex[dir] = k
		// init chanName
		newcb.cn[k] = &chanName{
			ch:  make(chan int, 0),
			dir: dir,
		}
	}
	go newcb.start()
	return newcb, nil
}

// Check --
func (c *SyncDirFile) Check(fileName string, dir string) (ok bool) {
	id, ok := c.dirMpIndex[dir]
	if !ok {
		return true
	}

	c.syncCBN.Lock()
	c.cn[id].name = fileName
	c.cn[id].isWait = true
	ch := c.cn[id].ch
	c.syncCBN.Unlock()
	<-ch
	return false
}

func (c *SyncDirFile) start() {
	var waitIndex, checkIndex int
	var waitName string
	var isCheck, runPD, checkDir bool
	var checkDirTime time.Time
	checkIndex = 0
	checkDirTime = time.Now()

	for {
		waitName = ""
		waitIndex = 0
		isCheck = false
		runPD = true
		if checkIndex >= c.length {
			checkIndex = 0
		}

		c.syncCBN.RLock()
		if c.cn[checkIndex].isWait {
			waitIndex = checkIndex
			waitName = c.cn[checkIndex].name
			isCheck = true
		}
		checkIndex++

		if isCheck && waitName != "" {
			for k, v := range c.cn {

				if k == waitIndex {
					continue
				}
				if v.name == "" {
					runPD = false
					break
				}
				if c.getCompareNameFunc(v.name) < c.getCompareNameFunc(waitName) {
					checkDir = true

					if !c.checkIslastFile(v.dir, v.name) {
						runPD = false
						break
					}
				}
			}
		} else {
			runPD = false
		}
		c.syncCBN.RUnlock()

		//check run
		if runPD {
			c.syncCBN.Lock()
			c.cn[waitIndex].ch <- 1
			c.cn[waitIndex].isWait = false
			c.syncCBN.Unlock()
		}

		if checkDir {
			if time.Since(checkDirTime) < 100*time.Millisecond {
				time.Sleep(500 * time.Millisecond)
			}
			checkDirTime = time.Now()
			checkDir = false
		}

		time.Sleep(10 * time.Millisecond)
	}
}

func (c *SyncDirFile) matchFileName(name string) bool {
	var pd bool
	if c.fileLike[0] != '*' && !strings.HasPrefix(name, c.fileLikeSlice[0]) {
		return false
	}

	if c.fileLike[len(c.fileLike)-1] != '*' && !strings.HasSuffix(name, c.fileLikeSlice[len(c.fileLikeSlice)-1]) {
		return false
	}

	for _, m := range c.fileLikeSlice {
		if strings.Index(name, m) == -1 {
			pd = false
			break
		}
		pd = true
	}
	return pd
}

func (c *SyncDirFile) checkIslastFile(dir, name string) bool {
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		return true
	}

	fileNames := make([]string, 0)
	for _, v := range infos {
		if v.IsDir() {
			continue
		}
		newName := v.Name()
		if c.matchFileName(newName) {
			fileNames = append(fileNames, newName)
		}
	}
	length := len(fileNames)
	if length == 0 {
		return true
	}
	sort.Slice(fileNames, func(i, j int) bool {
		return fileNames[i] < fileNames[j]
	})
	checkName := fileNames[length-1]
	if checkName > name {
		return false
	}
	return true
}

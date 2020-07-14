package directoryfile

import (
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"sync"
	"time"
)

/*
	除了FileNameID里面的Name为FullName,其他均为baseName
*/

// FileNameID --
type FileNameID struct {
	// full name
	Name string
	ID   uint
}

// Equle --
func (c *FileNameID) Equle(nameID FileNameID) (ok bool) {
	return c.Name == nameID.Name && c.ID == nameID.ID
}

// GetNameIDByFile --
func GetNameIDByFile(name string) (nameID FileNameID, err error) {
	stat, err := os.Stat(name)
	if err != nil {
		return nameID, err
	}
	nameID = FileNameID{
		Name: name,
		ID:   uint(reflect.ValueOf(stat.Sys()).Elem().FieldByName("Ino").Uint()),
	}
	return nameID, nil
}

//Directs 目录数据结构
type Directs struct {
	baseDir string

	// base name
	currentName string

	lastTime   time.Time
	offsetFlag bool

	fileLikeSlice []string
	fileLike      string

	fileList []os.FileInfo

	syncFileList sync.RWMutex
}

//NewDirects --
func NewDirects(baseDir, fileLike, offsetBaseName string) (ds *Directs, err error) {
	if fileLike == "" {
		return nil, errors.New("fileLike is empty")
	}

	if baseDir == "" {
		return nil, errors.New("baseDir is empty")
	}

	_, err = os.Stat(baseDir)
	if err != nil {
		return nil, err
	}

	ds = &Directs{
		baseDir:     baseDir,
		currentName: offsetBaseName,
		fileList:    make([]os.FileInfo, 0),
		lastTime:    time.Now(),
	}

	if offsetBaseName != "" {
		ds.offsetFlag = true
	}

	sp1 := strings.Split(fileLike, "*")
	for _, v := range sp1 {
		if v != "" {
			ds.fileLikeSlice = append(ds.fileLikeSlice, v)
		}
	}
	ds.fileLike = fileLike
	ds.setFileList()
	return ds, nil
}

func (c *Directs) setFileList() (err error) {
	infoList, err := c.getFileList()
	if err != nil {
		return err
	}
	c.syncFileList.Lock()
	c.fileList = infoList
	c.syncFileList.Unlock()
	return nil
}

//GetFileNameList 并发安全
func (c *Directs) GetFileNameList() (nameList []string) {
	nameList = make([]string, 0)
	c.syncFileList.RLock()
	for _, v := range c.fileList {
		nameList = append(nameList, v.Name())
	}
	c.syncFileList.RUnlock()
	return nameList
}

func (c *Directs) getCurrentName() (currentName string) {
	c.syncFileList.RLock()
	currentName = c.currentName
	c.syncFileList.RUnlock()
	return currentName
}

//GetNextFileName --
func (c *Directs) GetNextFileName() (fnd FileNameID, err error) {
	isFlush := true
	var name string
	nowTime := time.Now()
	if nowTime.Sub(c.lastTime) > time.Second {
		c.setFileList()
	}

	c.syncFileList.RLock()
	length := len(c.fileList)
	c.syncFileList.RUnlock()

	if length == 0 {
		if isFlush {
			return fnd, errors.New("Dir is empty")
		}
		err = c.setFileList()
		if err != nil {
			return fnd, err
		}
	}

	c.syncFileList.Lock()
	defer c.syncFileList.Unlock()

	if c.offsetFlag {
		c.offsetFlag = false
		name = filepath.Join(c.baseDir, c.currentName)
		if fileInfo, err := os.Stat(name); err == nil {
			fnd = FileNameID{
				Name: name,
				ID:   uint(reflect.ValueOf(fileInfo.Sys()).Elem().FieldByName("Ino").Uint()),
			}
			return fnd, nil
		}
	}
	if c.currentName == "" {
		c.currentName = c.fileList[0].Name()
		name = filepath.Join(c.baseDir, c.currentName)
		fnd = FileNameID{
			Name: name,
			ID:   uint(reflect.ValueOf(c.fileList[0].Sys()).Elem().FieldByName("Ino").Uint()),
		}
		return fnd, nil
	}

	length = len(c.fileList)
	index := sort.Search(length, func(i int) bool {
		return strings.Compare(c.fileList[i].Name(), c.currentName) >= 0
	})

	if index >= length {
		return fnd, errors.New("Get name failed")
	}

	nextNameIndex := index
	checkName := c.fileList[index].Name()
	if checkName == c.currentName {
		if index+1 < length {
			nextNameIndex = index + 1
		}
	}

	if nextNameIndex >= 0 {
		nextName := c.fileList[nextNameIndex].Name()
		name = filepath.Join(c.baseDir, nextName)
		c.currentName = nextName
		fnd = FileNameID{
			Name: name,
			ID:   uint(reflect.ValueOf(c.fileList[nextNameIndex].Sys()).Elem().FieldByName("Ino").Uint()),
		}
		return fnd, nil
	}
	return fnd, errors.New("Get file name failed")
}

//CheckIsLastFile baseName
func (c *Directs) CheckIsLastFile(baseName string) (ok bool) {
	c.syncFileList.RLock()
	length := len(c.fileList)
	if length > 0 && c.fileList[length-1].Name() == baseName {
		ok = true
	}
	c.syncFileList.RUnlock()
	return ok
}

func (c *Directs) getFileList() (infoList []os.FileInfo, err error) {
	var fileInfos []os.FileInfo
	fileInfos, err = c.readDir()
	if err != nil {
		return nil, err
	}
	infoList = make([]os.FileInfo, 0)
	for _, v := range fileInfos {
		if v.IsDir() {
			continue
		}
		pd := false
		if c.fileLike[0] != '*' && !strings.HasPrefix(v.Name(), c.fileLikeSlice[0]) {
			continue
		}
		if c.fileLike[len(c.fileLike)-1] != '*' && !strings.HasSuffix(v.Name(), c.fileLikeSlice[len(c.fileLikeSlice)-1]) {
			continue
		}
		for _, m := range c.fileLikeSlice {
			if strings.Index(v.Name(), m) == -1 {
				pd = false
				break
			}
			pd = true
		}
		if !pd {
			continue
		}

		infoList = append(infoList, v)
	}
	return infoList, nil
}

// 如果currentName存在的话，返回的fileInfo会以currentName开始
func (c *Directs) readDir() (listInfo []os.FileInfo, err error) {
	currentName := c.getCurrentName()
	df, err := os.Open(c.baseDir)
	if err != nil {
		return nil, err
	}
	defer df.Close()

	names, err := df.Readdirnames(-1)
	if err != nil {
		return nil, err
	}

	var index int
	lengthNames := len(names)

	if lengthNames == 0 {
		return nil, errors.New("Dir is empty")
	}

	sort.Strings(names)
	index = sort.Search(lengthNames, func(i int) bool {
		return strings.Compare(names[i], currentName) >= 0
	})
	if index >= lengthNames {
		return nil, errors.New("No suitable list found")
	}
	names = names[index:]
	listInfo = make([]os.FileInfo, len(names))
	for k, v := range names {
		listInfo[k], err = os.Stat(filepath.Join(c.baseDir, v))
		if err != nil {
			return nil, err
		}
	}
	return listInfo, nil
}

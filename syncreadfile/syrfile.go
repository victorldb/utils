package syncreadfile

import (
	"path/filepath"

	"github.com/victorldb/utils/callbackmsg"
	"github.com/victorldb/utils/cbn"
	"github.com/victorldb/utils/file"
	"github.com/victorldb/utils/readfileline"
)

// SyncfConfig --
type SyncfConfig struct {
	DataType            int
	Dir                 string
	FileMatch           string
	OffsetName          string
	Offset              int64
	CallBackData        chan *readfileline.ContentType
	RmLineFlag          bool
	ContentBaseNameFlag bool
	CBM                 callbackmsg.CallbackMsg
	Syncf               *cbn.SyncDirFile
}

// SyReadFile --
type SyReadFile struct {
	cfg          SyncfConfig
	currentName  string
	offset       int64
	callBackData chan *readfileline.ContentType

	lgf   readfileline.DirFileConfig
	syncf *cbn.SyncDirFile
}

// NewDataHander --
func NewDataHander(cfg SyncfConfig) (hander *SyReadFile) {
	hander = &SyReadFile{
		cfg:          cfg,
		callBackData: cfg.CallBackData,
		syncf:        cfg.Syncf,
		lgf: readfileline.DirFileConfig{
			Dir:                 cfg.Dir,
			DataType:            cfg.DataType,
			FileLike:            cfg.FileMatch,
			OffsetName:          cfg.OffsetName,
			OffsetSize:          cfg.Offset,
			RmLineFlag:          cfg.RmLineFlag,
			ContentBaseNameFlag: cfg.ContentBaseNameFlag,
			CallBM:              cfg.CBM,
		},
	}
	return hander
}

// StartRead --
func (c *SyReadFile) StartRead() (err error) {
	// 初始有偏移的文件，检查是否需要同步检查
	if c.cfg.OffsetName != "" && file.FileExists(filepath.Join(c.cfg.Dir, c.cfg.OffsetName)) {
		c.currentName = c.cfg.OffsetName
		if c.syncf != nil {
			c.cfg.CBM.RegInfo("cbn:<recover> wait fullName:%s", filepath.Join(c.cfg.Dir, c.currentName))
			c.syncf.Check(c.currentName, c.cfg.Dir)
			c.cfg.CBM.RegInfo("cbn:<recover> active fullName:%s", filepath.Join(c.cfg.Dir, c.currentName))
		}
	}

	// 配置启动
	contents := make(chan *readfileline.ContentType, 100)
	rd, err := readfileline.NewReadLine(c.lgf, contents)
	if err != nil {
		return err
	}
	go rd.Start()
	defer func() {
		rd.Close()
	}()

	// 开始读取数据
	for {
		select {
		case content := <-contents:
			if c.currentName != content.Name {
				c.currentName = content.Name
				if c.syncf != nil && c.currentName != "" {
					c.cfg.CBM.RegInfo("cbn:wait fullName:%s", filepath.Join(c.cfg.Dir, c.currentName))
					c.syncf.Check(c.currentName, c.cfg.Dir)
					c.cfg.CBM.RegInfo("cbn:active fullName:%s", filepath.Join(c.cfg.Dir, c.currentName))
				}
			}
			c.offset = content.Offset
			c.callBackData <- content
		}
	}
}

// func (c *SyReadFile)GetRecoverStatus()(name string,offset int64){
// 	return
// }

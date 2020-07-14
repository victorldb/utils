package pid

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/victorldb/utils/file"

	"github.com/shirou/gopsutil/process"
)

// WritePid --
func WritePid(filename string) (err error) {
	if filename == "" {
		return errors.New("fileName is empty")
	}

	var pid int32
	if file.FileExists(filename) {
		bpid, err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}
		spid := string(bpid)
		//match service sh script
		if spid != "" {
			temp, err := strconv.Atoi(spid)
			if err != nil {
				return err
			}
			pid = int32(temp)
			if pid > 0 {
				ok, err := process.PidExists(pid)
				if err != nil {
					return err
				}
				if ok {
					return errors.New("process is running")
				}
			}
		}
	}

	pidParentDir := filepath.Dir(filename)
	if !file.FileExists(pidParentDir) {
		err = os.MkdirAll(pidParentDir, 0666)
		if err != nil {
			return err
		}
	}
	return ioutil.WriteFile(filename, []byte(strconv.Itoa(os.Getpid())), 0755)
}

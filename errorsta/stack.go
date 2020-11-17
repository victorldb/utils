package errorsta

import (
	"fmt"
	"path/filepath"
	"runtime"
)

// LastStackError --
func LastStackError(err error) error {
	if err == nil {
		return nil
	}
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		err = fmt.Errorf("%+w\n%s:%d %s", err, filepath.Base(file), line, runtime.FuncForPC(pc).Name())
	}
	return err
}

package file

import (
	"os"
	"syscall"
)

// MkdirAll --
func MkdirAll(path string, perm os.FileMode) (err error) {
	oldMask := syscall.Umask(0)
	err = os.MkdirAll(path, perm)
	syscall.Umask(oldMask)
	return err
}

// Mkdir --
func Mkdir(path string, perm os.FileMode) (err error) {
	oldMask := syscall.Umask(0)
	err = os.Mkdir(path, perm)
	syscall.Umask(oldMask)
	return err
}

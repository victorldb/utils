package utils

import (
	"errors"
	"os"
	"path/filepath"
)

// GetCurrentPath --
func GetCurrentPath() (currentPath string, err error) {
	if len(os.Args) < 1 {
		return "", errors.New("args length is error")
	}
	currentPath, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	return currentPath, nil
}

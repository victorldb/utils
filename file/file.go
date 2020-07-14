// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package file

import (
	"bufio"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

// SelfPath gets compiled executable file absolute path
func SelfPath() string {
	path, _ := filepath.Abs(os.Args[0])
	return path
}

// SelfDir gets compiled executable file directory
func SelfDir() string {
	return filepath.Dir(SelfPath())
}

// FileExists reports whether the named file or directory exists.
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// Search a file in paths.
// this is often used in search config file in /etc ~/
func SearchFile(filename string, paths ...string) (fullpath string, err error) {
	for _, path := range paths {
		if fullpath = filepath.Join(path, filename); FileExists(fullpath) {
			return
		}
	}
	err = errors.New(fullpath + " not found in paths")
	return
}

// like command grep -E
// for example: GrepFile(`^hello`, "hello.txt")
// \n is striped while read
func GrepFile(patten string, filename string) (lines []string, err error) {
	re, err := regexp.Compile(patten)
	if err != nil {
		return
	}

	fd, err := os.Open(filename)
	if err != nil {
		return
	}
	lines = make([]string, 0)
	reader := bufio.NewReader(fd)
	prefix := ""
	isLongLine := false
	for {
		byteLine, isPrefix, er := reader.ReadLine()
		if er != nil && er != io.EOF {
			return nil, er
		}
		if er == io.EOF {
			break
		}
		line := string(byteLine)
		if isPrefix {
			prefix += line
			continue
		} else {
			isLongLine = true
		}

		line = prefix + line
		if isLongLine {
			prefix = ""
		}
		if re.MatchString(line) {
			lines = append(lines, line)
		}
	}
	return lines, nil
}

func CopyFile(src, dst string) (err error) {
	fsrc, err := os.Open(src)
	defer fsrc.Close()
	if err != nil {
		return err
	}
	if FileExists(dst) {
		err := os.Remove(dst)
		if err != nil {
			return err
		}
	}
	fdst, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY, 0755)
	defer fdst.Close()
	if err != nil {
		return err
	}
	_, err = io.Copy(fdst, fsrc)
	if err != nil {
		if FileExists(dst) {
			os.Remove(dst)
		}
		return err
	}
	return nil
}

func MoveFile(src, dst string, modePerm os.FileMode) (err error) {
	if FileExists(dst) {
		err = os.Remove(dst)
		if err != nil {
			return err
		}
		err = os.Rename(src, dst)
		if err != nil {
			return err
		}
		err = os.Chmod(dst, modePerm)
		if err != nil {
			return err
		}
	} else {
		err = os.Rename(src, dst)
		if err != nil {
			return err
		}
		err = os.Chmod(dst, modePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func CleanFile(dir string, unfileName []string) (err error) {
	var isUnFile bool
	dir_list, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	if len(dir_list) > 0 {
		for _, v := range dir_list {
			if !v.IsDir() {
				fName := v.Name()
				isUnFile = false
				for _, v := range unfileName {
					if v == fName {
						isUnFile = true
						break
					}
				}
				if !isUnFile {
					delFilePath := filepath.Join(dir, fName)
					os.Remove(delFilePath)
				}
			}
		}
	}
	return nil
}

// GetCurrentABSDir --
func GetCurrentABSDir() (name string, err error) {
	if len(os.Args) == 0 {
		return "", errors.New("get current dir failed")
	}
	absPath, err := filepath.Abs(os.Args[0])
	if err != nil {
		return "", err
	}
	name = filepath.Dir(absPath)
	if name == "" {
		return "", errors.New("get current dir failed")
	}
	return name, nil
}

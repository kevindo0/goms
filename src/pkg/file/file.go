package file

import (
	"fmt"
	"os"
	"path"
)

// check path or file is exist
// if not exist return true
func CheckNotExist(src string) bool {
	_, err := os.Stat(src)
	return os.IsNotExist(err)
}

func IfNotExistMkdir(src string) error {
	if res := CheckNotExist(src); res == true {
		if err := Mkdir(src); err != nil {
			return err
		}
	}
	return nil
}

// mkdir
func Mkdir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return  err
	}
	return nil
}

// open file with flag and perm
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func MustOpen(filePath, filename string) (*os.File, error) {
	err := IfNotExistMkdir(filePath)
	if err != nil {
		return nil, err
	}
	src := path.Join(filePath, filename)
	f, err := Open(src, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}
	return f, nil
}
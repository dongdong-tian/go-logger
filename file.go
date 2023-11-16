package logger

import (
	"fmt"
	"os"
)

/*
	fileIsExist 文件是否存在
*/
func fileIsExist(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

/*
	checkPermission 文件是否有权限
*/
func checkPermission(path string) bool {
	_, err := os.Stat(path)
	return os.IsPermission(err)

}

/*
	DirNotExistMkdir 目录是否存在
*/
func DirNotExistMkdir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := mkDir(path)
		return err
	}

	return nil

}

/*
	mkDir 创建目录
*/
func mkDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	return err
}

/*
	mustOpen 创建日志文件
*/
func mustOpen(name, dir string) (*os.File, error) {
	var f *os.File
	if perm := checkPermission(dir); perm {
		return nil, fmt.Errorf("%s directory not permission", dir)
	}
	err := DirNotExistMkdir(dir)
	if err != nil {
		return nil, fmt.Errorf("%s directory create fail", dir)
	}
	f, err = os.OpenFile(dir+string(os.PathSeparator)+name, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("%s log file open fail", name)
	}

	return f, nil

}

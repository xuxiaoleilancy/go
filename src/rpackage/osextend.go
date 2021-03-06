package rpackage

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	Separator     = string(os.PathSeparator)     // 路径分隔符（分隔路径元素）
	ListSeparator = string(os.PathListSeparator) // 路径列表分隔符（分隔多个路径）
)

//目录是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//拷贝文件，并自动创建目录
func CopyFileAdvanced(src, des string) (w int64, err error) {
	exist, err := PathExists(filepath.Dir(des))
	if err != nil {
		return 0, err
	}
	if !exist {
		err := os.MkdirAll(filepath.Dir(des), 0766)
		if err != nil {
			return 0, err
		}
	}
	return CopyFile(src, des)
}

//拷贝文件，调用此接口之前请确保目标目录存在
func CopyFile(src, des string) (w int64, err error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	desFile, err := os.Create(des)
	if err != nil {
		return 0, err
	}
	defer desFile.Close()

	return io.Copy(desFile, srcFile)
}

func GetFiles(dirPath, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix)

	err = filepath.Walk(dirPath, func(filename string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fi.IsDir() {
			return nil
		}

		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}

		return nil
	})

	return files, err
}

func Readline(fileName string) (ret []string, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return ret, err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		ret = append(ret, line)
		if err != nil {
			if err == io.EOF {
				return ret, nil
			}
			return ret, err
		}
	}
}

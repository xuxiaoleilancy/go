package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func readline(fileName string) (ret []string, err error) {
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
	return ret, nil
}

func Print(line string) {
	//fmt.Println(line)
}

func containsString(list []string, val string) bool {
	for _, str := range list {
		if val == str {
			return true
		}
	}
	return false
}
func scanDir(scanPath string) {

	//获取要删除的目录名列表
	removepathList, errRead := readline("removepath.txt")
	if errRead != nil {
		return
	}
	fmt.Println("removepathlist is :", removepathList)

	//获取要删除的文件扩展名列表
	removeFileList, errRead := readline("removefile.txt")
	if errRead != nil {
		return
	}
	fmt.Println("removefilelist is :", removeFileList)

	if errRead != nil {
		return
	}
	err := filepath.Walk(scanPath, func(scanPath string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			pathName := filepath.Base(scanPath)
			if containsString(removepathList, pathName) {
				println("remove path----" + scanPath)
				os.RemoveAll(scanPath)
			}
			return nil
		}

		var filenameWithSuffix string
		filenameWithSuffix = path.Base(scanPath) //获取文件名带后缀
		var fileSuffix string
		fileSuffix = path.Ext(filenameWithSuffix) //获取文件后缀
		if containsString(removeFileList, fileSuffix) {
			println("remove file****" + scanPath)
			os.Remove(scanPath)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func main() {
	flag.Parse()

	args := os.Args //获取用户输入的所有参数

	if args == nil || len(args) < 2 {
		root := "../../../"
		println(filepath.Dir(root))
		scanDir(root)
	} else {
		root := args[1]
		println(filepath.Dir(root))
		scanDir(root)
	}
}

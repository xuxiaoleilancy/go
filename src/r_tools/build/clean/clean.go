package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"rpackage"
)

var (
	//获取要删除的目录名列表
	removepathList, errRemovePathRead = rpackage.Readline("removepath.txt")

	//获取要删除的文件扩展名列表
	removeFileList, errRemoveFileRead = rpackage.Readline("removefile.txt")

	//获取例外目录名列表
	removepathExcludeList, errRemovePathExcludeRead = rpackage.Readline("removepathexclude.txt")
)

func cleanFunc(scanPath string, f os.FileInfo, err error) error {
	if f == nil {
		return err
	}
	println("path----" + scanPath)
	return nil
	if f.IsDir() {
		pathName := filepath.Base(scanPath)
		if rpackage.ContainsString(removepathExcludeList, pathName) {
			return nil
		}
		if rpackage.ContainsString(removepathList, pathName) {
			println("remove path----" + scanPath)
			os.RemoveAll(scanPath)
		}
		return nil
	}

	var filenameWithSuffix string
	filenameWithSuffix = path.Base(scanPath) //获取文件名带后缀
	var fileSuffix string
	fileSuffix = path.Ext(filenameWithSuffix) //获取文件后缀
	if rpackage.ContainsString(removeFileList, fileSuffix) {
		println("remove file****" + scanPath)
		os.Remove(scanPath)
	}
	return nil
}

func scanDir(scanPath string) {
	if errRemovePathRead != nil {
		return
	}
	fmt.Println("removepathlist is :", removepathList)
	if errRemoveFileRead != nil {
		return
	}
	fmt.Println("removefilelist is :", removeFileList)

	err := filepath.Walk(scanPath, cleanFunc)
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func findRPath(curDir string, rootDir string) string {

	parentPath := filepath.Dir(curDir)
	folderName := filepath.Base(parentPath)
	//fmt.Println("parent dir is : ", parentPath, " folder name is :", folderName)
	if folderName == "R" {
		relPath, _ := filepath.Rel(rootDir, parentPath)
		fmt.Println("parent dir is : ", parentPath, " rootDir is :", rootDir, " relPath is : ", relPath)
		return relPath
	} else if folderName == filepath.VolumeName(rootDir) ||
		folderName == rpackage.Separator {
		return ""
	} else {
		return findRPath(parentPath, rootDir)
	}
}
func parseArgs(args []string) string {
	var scanPath string
	if args == nil || len(args) < 2 {
		curDir, _ := os.Getwd()
		curDir, _ = filepath.Abs(curDir)
		fmt.Println(curDir)
		scanPath = findRPath(curDir, curDir)

	} else {
		scanPath = args[1]
	}
	return scanPath
}
func main() {
	flag.Parse()

	args := os.Args //获取用户输入的所有参数
	scanPath := parseArgs(args)

	if scanPath != "" {
		scanPath, _ = filepath.Abs(scanPath)
		fmt.Println(" this program will clean the path: ", scanPath)
		scanDir(scanPath)
	} else {
		fmt.Println(" this program will not clean any path ")
	}
}

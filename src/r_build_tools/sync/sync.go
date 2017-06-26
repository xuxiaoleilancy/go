package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"rpackage"
	"strings"
	//"syscall"
)

var (
	syncRootDir string
)

func parseSyncFileList2Map(mapStrings []string) (retMap map[string]string) {

	retmap := make(map[string]string)
	for _, str := range mapStrings {

		fileMap := strings.Split(str, "=>")
		if len(fileMap) == 2 {
			srcHeadFileName := strings.TrimSpace(fileMap[0])
			destHeadFileName := strings.TrimSpace(fileMap[1])

			srcHeadFileName = strings.Trim(srcHeadFileName, "\"")
			destHeadFileName = strings.Trim(destHeadFileName, "\"")

			retmap[srcHeadFileName] = destHeadFileName
		} else {
			fmt.Println(fileMap)
		}
	}

	return retmap
}

func copyFile2Include(filesMap map[string]string) {
	for fileFullName, targetFileName := range filesMap {
		os.Chdir(syncRootDir)

		relFilePath := strings.TrimPrefix(fileFullName, syncRootDir)
		destFileName := rpackage.Separator + "include" + relFilePath
		targetFileFullName := syncRootDir + rpackage.Separator + "include" + rpackage.Separator + targetFileName
		tmpStrList := strings.Split(destFileName, rpackage.Separator)

		fmt.Println("**********************************")
		fmt.Println("fileFullName:", fileFullName)
		fmt.Println("syncRootDir:", syncRootDir)
		fmt.Println("relfilePath:", relFilePath)
		fmt.Println("destFileName:", destFileName)
		fmt.Println("targetFileFullName:", targetFileFullName)
		fmt.Println("tmpStrList:", tmpStrList)

		if len(tmpStrList) > 1 {
			for _, subDir := range tmpStrList {
				if subDir != tmpStrList[len(tmpStrList)-1] &&
					subDir != "" {
					curPath, _ := os.Getwd()
					ok, _ := rpackage.PathExists(curPath + rpackage.Separator + subDir)
					if !ok {
						//oldMask := syscall.Umask(0)
						os.Mkdir(subDir, 0777)
						//syscall.Umask(oldMask)
						//os.Chmod(subDir, os.FileMode)
						fmt.Println("mkdir: ", subDir)
					}
					os.Chdir(subDir)
				} else {

				}
			}
		}

		//rpackage.CopyFile(fileFullName, syncRootDir+destFileName)
		fmt.Println(fileFullName, " ====>>>> ", syncRootDir+destFileName)
		desFile, err := os.Create(targetFileFullName)
		if err != nil {
		}

		fmt.Println(fileFullName, " ====>>>> ", syncRootDir+destFileName)
		relFilePath = strings.TrimPrefix(relFilePath, rpackage.Separator)
		fmt.Println(" relFilePath after TrimPrefix is : ", relFilePath)
		relFilePath = strings.Replace(relFilePath, rpackage.Separator, "/", -1)
		desFile.WriteString("#include " + "\"../" + relFilePath + "\"")
		defer desFile.Close()
		fmt.Println("**********************************")
		continue
	}
}
func getSyncFileList(curDir string) (filesMap map[string]string, err error) {

	filesMap = make(map[string]string)

	//解析要同步的头文件列表
	syncFileList, err := rpackage.Readline(syncRootDir + rpackage.Separator + "sync.profile")
	if err != nil {
		return filesMap, err
	}

	//	获取当前目录下的所有.h文件
	files, err := rpackage.GetFiles(curDir, ".h")
	syncFileMaps := parseSyncFileList2Map(syncFileList)
	fmt.Println("sync.profile-------------")
	for fileFullName, fileMap2BaseName := range syncFileMaps {
		fmt.Println(fileFullName, "=>", fileMap2BaseName)
	}
	fmt.Println("sync.profile-------------")

	for _, fileFullName := range files {
		fileBaseName := strings.TrimPrefix(fileFullName, curDir+rpackage.Separator)
		fileBaseName = strings.Replace(fileBaseName, "\\", "/", -1)
		fmt.Println("--------------", fileBaseName)
		fileMapName, ok := syncFileMaps[fileBaseName]
		if ok {
			filesMap[fileFullName] = fileMapName
		}
	}

	return filesMap, err
}

func sync(curDir string) {
	includePath := curDir + rpackage.Separator + "include"

	//删除原来的include目录
	ok, _ := rpackage.PathExists(includePath)
	if ok {
		os.RemoveAll(includePath)
	}

	syncFileMaps, _ := getSyncFileList(curDir)
	for fileFullName, fileMap2BaseName := range syncFileMaps {
		fmt.Println(fileFullName, "=>", fileMap2BaseName)
	}

	//编译时，指向源文件目录里的头文件
	copyFile2Include(syncFileMaps)
}

func parseArgs(args []string) string {
	var scanPath string
	if args == nil || len(args) < 2 {
		scanPath, _ = os.Getwd()
	} else {
		scanPath = args[1]

		//strTestPath := "D:/code/suirui/svn/1_CodeLib/05.PC/SRQT/R/src/srfoundation/src"
		fmt.Println("args[1] is :", scanPath)
		scanPath, _ = filepath.Abs(scanPath)
	}
	return scanPath
}

func main() {
	flag.Parse()
	args := os.Args //获取用户输入的所有参数
	syncRootDir = parseArgs(args)

	fmt.Println(syncRootDir)
	sync(syncRootDir)
}

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"rpackage"
	"strconv"
	"strings"
)

var (
	syncRootDir        string
	syncDestIncludeDir string
	syncSrcDir         string
	syncProfile        string
	bPrintLog          bool

	MD5_LINE          int
	INCLUDE_FILE_LINE int
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
			printLogln(fileMap)
		}
	}

	return retmap
}

func copyFile2Include(filesMap map[string]string) {
	for fileFullName, targetFileName := range filesMap {
		os.Chdir(syncRootDir)

		relFilePath := strings.TrimPrefix(fileFullName, syncSrcDir)
		destFileName := rpackage.Separator + "include" + relFilePath
		targetFileFullName := syncDestIncludeDir + rpackage.Separator + targetFileName
		tmpStrList := strings.Split(destFileName, rpackage.Separator)

		printLogln("**********************************")
		printLogln("fileFullName:", fileFullName)
		printLogln("syncRootDir:", syncRootDir)
		printLogln("relfilePath:", relFilePath)
		printLogln("destFileName:", destFileName)
		printLogln("targetFileFullName:", targetFileFullName)
		printLogln("tmpStrList:", tmpStrList)

		if len(tmpStrList) > 1 {
			for _, subDir := range tmpStrList {
				if subDir != tmpStrList[len(tmpStrList)-1] &&
					subDir != "" && subDir == "include" {
					curPath, _ := os.Getwd()
					ok, _ := rpackage.PathExists(curPath + rpackage.Separator + subDir)
					if !ok {
						//oldMask := syscall.Umask(0)
						os.Mkdir(subDir, 0777)
						//syscall.Umask(oldMask)
						//os.Chmod(subDir, os.FileMode)
						printLogln("mkdir: ", subDir)
					}
					os.Chdir(subDir)
				} else {

				}
			}
		}

		//rpackage.CopyFile(fileFullName, syncRootDir+destFileName)
		printLogln(fileFullName, " ====>>>> ", syncRootDir+destFileName)
		newMd5Value, err := rpackage.Md5SumFile(fileFullName)
		//newMd5String := fmt.Sprintf("%x\n", newMd5Value)

		var oldIncludeString string

		if err != nil {

		}

		ok, _ := rpackage.PathExists(targetFileFullName)

		bUpdate := true
		relFilePath = strings.TrimPrefix(relFilePath, rpackage.Separator)
		printLogln(" relFilePath after TrimPrefix is : ", relFilePath)
		relFilePath = strings.Replace(relFilePath, rpackage.Separator, "/", -1)

		if ok {
			printLogln("!!!!!!!!!!!!!!!!!!!", targetFileFullName)
			f, err := os.Open(targetFileFullName)
			if err != nil {
				panic(err)
			}
			defer f.Close()

			rd := bufio.NewReader(f)
			nLen := 0
			for {
				line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
				if err != nil || io.EOF == err {
					break
				}
				nLen++
				//				if nLen == MD5_LINE {
				//					var oldMd5String = line
				//				}
				//				fmt.Println("~~~~~~~~~~~~~~~~~~~")
				//				fmt.Printf("%d,%s", nLen, line)
				//				fmt.Println("~~~~~~~~~~~~~~~~~~~")
				if nLen == INCLUDE_FILE_LINE {
					oldIncludeString = line

					newIncludeString := fmt.Sprintf("#include " + "\"../src/" + relFilePath + "\"\n")
					if strings.EqualFold(newIncludeString, oldIncludeString) {
						bUpdate = false
					}
				}
			}
		}

		//		if strings.EqualFold(newMd5String, line) {
		//			bUpdate = false
		//		} else {
		//			fmt.Println("-------------------")
		//			fmt.Println(newMd5String)
		//			fmt.Println(line)
		//			fmt.Println(strings.EqualFold(newMd5String, line))
		//			fmt.Println("-------------------")
		//		}

		if bUpdate {
			desFile, err := os.Create(targetFileFullName)
			if err != nil {
			}

			printLogln(fileFullName, " ====>>>> ", syncRootDir+destFileName)

			desFile.WriteString("/*" + "\n")
			fmt.Fprintf(desFile, "%x\n", newMd5Value)
			desFile.WriteString(fileFullName + "\n")
			//desFile.WriteString(str + "\n")
			desFile.WriteString("*/" + "\n")
			desFile.WriteString("#include " + "\"../src/" + relFilePath + "\"\n")
			defer desFile.Close()
		}
		printLogln("**********************************")
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
	printLogln("sync.profile-------------")
	for fileFullName, fileMap2BaseName := range syncFileMaps {
		printLogln(fileFullName, "=>", fileMap2BaseName)
	}
	printLogln("sync.profile-------------")

	for _, fileFullName := range files {
		fileBaseName := strings.TrimPrefix(fileFullName, curDir+rpackage.Separator)
		fileBaseName = strings.Replace(fileBaseName, "\\", "/", -1)
		printLogln("--------------", fileBaseName)
		fileMapName, ok := syncFileMaps[fileBaseName]
		if ok {
			filesMap[fileFullName] = fileMapName
		}
	}

	return filesMap, err
}

func sync() {

	//删除原来的include目录
	ok, _ := rpackage.PathExists(syncDestIncludeDir)
	if ok {
		//为了解决每次include文件夹都更新，影响编译速度的问题，这里不自动删除，可选择手动删除
		//os.RemoveAll(syncDestIncludeDir)
	} else {

	}

	syncFileMaps, _ := getSyncFileList(syncSrcDir)
	for fileFullName, fileMap2BaseName := range syncFileMaps {
		printLogln(fileFullName, "=>", fileMap2BaseName)
	}

	//编译时，指向源文件目录里的头文件
	copyFile2Include(syncFileMaps)
}

func parseArgs(args []string) string {
	var scanPath string
	if args == nil || len(args) < 2 {
		scanPath, _ = os.Getwd()
		//bPrintLog = false
	} else if len(args) < 3 {
		scanPath = args[1]
		//bPrintLog = false

		//strTestPath := "D:/code/suirui/svn/1_CodeLib/05.PC/SRQT/R/src/srfoundation/src"
		printLogln("args[1] is :", scanPath)
		scanPath, _ = filepath.Abs(scanPath)
	} else {
		scanPath = args[1]
		bPrintLog, _ = strconv.ParseBool(args[2])

		//strTestPath := "D:/code/suirui/svn/1_CodeLib/05.PC/SRQT/R/src/srfoundation/src"
		printLogln("args[1] is :", scanPath)
		scanPath, _ = filepath.Abs(scanPath)
	}
	return scanPath
}

func printLogln(a ...interface{}) (n int, err error) {
	if bPrintLog {
		return fmt.Println(a...)
	}

	return 0, nil
}
func initConfig() {
	syncDestIncludeDir = syncRootDir + rpackage.Separator + "include"
	syncSrcDir = syncRootDir + rpackage.Separator + "src"
	syncProfile = syncRootDir + rpackage.Separator + "sync.profile"

	printLogln("~~~~~~~~syncDestIncludeDir(", syncDestIncludeDir, ")")
	printLogln("~~~~~~~~syncSrcDir(", syncSrcDir, ")")
	printLogln("~~~~~~~~syncProfile(", syncProfile, ")")

	MD5_LINE = 2
	INCLUDE_FILE_LINE = 5
}
func main() {
	bPrintLog = false

	flag.Parse()
	args := os.Args //获取用户输入的所有参数
	syncRootDir = parseArgs(args)
	printLogln(syncRootDir)

	initConfig()

	sync()

	fmt.Println("~~~~~~~~Sync Finished!!(", syncRootDir, ")")
}

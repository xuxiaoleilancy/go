package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"rpackage"
	"strings"
)

func syncMap(mapStrings []string) (retMap map[string]string) {

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

func createSyncFile(srcHeadFileName string, destHeadFileName string) {

}
func scanDir(scanPath string) {

	//解析要同步的头文件列表
	syncFileList, errRead := rpackage.Readline("sync.profile")
	if errRead != nil {
		return
	}

	retmap := syncMap(syncFileList)
	for k := range retmap {
		fmt.Println("head file is", k, " map to ", retmap[k])
	}

	err := filepath.Walk(scanPath, func(scanPath string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		if f.IsDir() {
			fmt.Println(" subdir is : " + f.Name())
			if f.Name() != "inclue" &&
				f.Name() != scanPath {
				os.Mkdir("include"+"//"+f.Name(), os.ModeDir)
			}
			return nil
		}

		var filenameWithSuffix string

		filenameWithSuffix = path.Base(f.Name()) //获取文件名带后缀

		var fileSuffix string
		fileSuffix = path.Ext(filenameWithSuffix) //获取文件后缀

		fmt.Println("filenameWithSuffix is ", filenameWithSuffix, " fileSuffix is ", fileSuffix)

		if fileSuffix == ".h" {
			var ok bool
			_, ok = retmap[filenameWithSuffix]

			if ok {
				rpackage.CopyFile(filenameWithSuffix, "include"+"//"+filenameWithSuffix)
				desFile, err := os.Create("include" + "//" + retmap[filenameWithSuffix])
				if err != nil {
				}
				desFile.WriteString("#include " + "\"" + filenameWithSuffix + "\"")
				defer desFile.Close()
				//os.OpenFile(retmap[filenameWithSuffix])
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func main() {
	flag.Parse()

	curDir, _ := os.Getwd()
	os.Chdir(curDir)
	fmt.Println(curDir)
	os.RemoveAll("include")
	os.Mkdir("include", os.ModeDir)
	scanDir(curDir)
}

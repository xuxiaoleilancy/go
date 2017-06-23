// myapp1 project main.go
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"rpackage"
	"strings"
)

func main() {
	fmt.Println("Hello World!")

	curDir, _ := os.Getwd()

	os.RemoveAll("include")
	//	return
	files, err := rpackage.WalkDir(curDir, ".h")

	fmt.Println(files, err)

	for _, str := range files {
		os.Chdir(curDir)

		tmpStr := strings.TrimPrefix(str, curDir)
		fmt.Println("******:----", tmpStr)
		tmpStr = "include" + tmpStr
		tmpStrList := strings.Split(tmpStr, "\\")

		fmt.Println("----------list:----", tmpStrList)

		if len(tmpStrList) > 1 {
			for _, subDir := range tmpStrList {
				if subDir != tmpStrList[len(tmpStrList)-1] {
					fmt.Println("mkdir: ", subDir)
					os.Mkdir(subDir, os.ModeDir)
					os.Chdir(subDir)
				} else {

				}
			}
		}

		rpackage.CopyFile(str, curDir+filepath.ToSlash(rpackage.Separator)+tmpStr)
		desFile, err := os.Create(curDir + "\\" + "include" + "\\" + strings.ToUpper(tmpStrList[len(tmpStrList)-1]))
		if err != nil {
		}

		relPath, _ := filepath.Rel(curDir, str)
		relPath = strings.Replace(relPath, "\\", "/", -1)

		fmt.Println("##########relPath:", relPath)
		desFile.WriteString("#include " + "\"" + relPath + "\"")
		defer desFile.Close()
	}

}

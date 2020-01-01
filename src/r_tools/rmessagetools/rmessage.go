package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"rpackage"
	"strings"
	"time"
)

type RMsgStruct struct {
	Type    string          `json:"type"`
	Members [][]interface{} `json:"members"`
}

type RMsgStructs struct {
	Headers []struct {
		Header string `json:"header"`
	} `json:"headers"`

	Structs []RMsgStruct `json:"structs"`
}

func name_ver(cName string) string {
	datatype := map[string]string{}
	datatype["type"] = "G_type"
	_, ok := datatype[cName]
	if ok {
		return datatype[cName]
	}
	return "G_" + cName
}
func isVector(s string) bool {
	return strings.HasPrefix(s, "[") && strings.HasSuffix(s, "]")
}
func isMap(s string) bool {
	return strings.HasPrefix(s, "{") && strings.HasSuffix(s, "}")
}
func scanJsonFiles() {
	datatype := map[string]string{}

	datatype["int32_t"] = "int32"
	datatype["bool"] = "bool"
	datatype["string"] = "string"
	datatype["uint32_t"] = "int32"
	datatype["int8_t"] = "int8"
	datatype["double"] = "float64"
	datatype["float"] = "float32"
	datatype["uint64_t"] = "int64"

	all_json_files, _ := rpackage.GetFiles(".", ".json")
	b, _ := rpackage.PathExists("../../rmessageforgolang")
	if !b {
		os.Mkdir("../../rmessageforgolang", os.ModePerm)
	}

	for _, tmpFile := range all_json_files {

		fmt.Println("\n-------------------", tmpFile)
		b, err := ioutil.ReadFile(tmpFile)
		if err != nil {
			fmt.Print(err)
		}
		str := string(b)
		structs := RMsgStructs{}

		fmt.Println(str)
		body := bytes.TrimPrefix([]byte(str), []byte("\xef\xbb\xbf"))
		err = json.Unmarshal(body, &structs)

		if err != nil {
			b, _ := rpackage.PathExists("../../rmessageforgolang/error_bak")
			if !b {
				os.Mkdir("../../rmessageforgolang/error_bak", os.ModePerm)
			}
			rpackage.CopyFile(tmpFile, "../../rmessageforgolang/error_bak/"+filepath.Base(tmpFile))
			fmt.Println("json unmarshal error:", err)
			continue
		}

		g_code := "/*Copyright (C) 2020 随锐集团科技有限公司\n"
		g_code += " *版权所有。\n"
		g_code += " *时间："
		g_code += time.Now().Format("2006-01-02 15:04:05")
		g_code += "\n"
		g_code += " *创建者： 徐晓磊\n"
		g_code += " *版本：V1.0.0\n"
		g_code += "*/\n"
		g_code += "package rmessage\n\n"

		for i := range structs.Headers {
			g_code += "//import \" "
			g_code += structs.Headers[i].Header
			g_code += " \"\n"
		}

		for i := range structs.Structs {
			g_code += "type " + structs.Structs[i].Type + " struct {\n"
			for j := range structs.Structs[i].Members {
				g_code += "\t"
				g_code += name_ver(structs.Structs[i].Members[j][1].(string))
				g_code += " "
				_, ok := datatype[structs.Structs[i].Members[j][0].(string)]
				if ok {
					g_code += datatype[structs.Structs[i].Members[j][0].(string)]
				} else {
					fmt.Println(structs.Structs[i].Members[j][0])
					go_type_des := structs.Structs[i].Members[j][0].(string)
					if isMap(go_type_des) {

						fmt.Println("*******************this is a map")
						go_type_des = strings.TrimPrefix(go_type_des, "{")
						go_type_des = strings.TrimSuffix(go_type_des, "}")
						_, ok := datatype[go_type_des]
						if ok {
							go_type_des = datatype[go_type_des]
						}
						go_type_des = "map[string]" + go_type_des
					}
					if isVector(go_type_des) {
						fmt.Println("*******************this is a vector")
						go_type_des = strings.TrimPrefix(go_type_des, "[")
						go_type_des = strings.TrimSuffix(go_type_des, "]")
						_, ok := datatype[go_type_des]
						if ok {
							go_type_des = datatype[go_type_des]
						}
						go_type_des = "[]" + go_type_des
					}
					g_code += go_type_des
				}

				g_code += " `json:\"" + structs.Structs[i].Members[j][1].(string) + "\"`"
				g_code += "\n"

			}
			g_code += "}"
			g_code += "\n"
			g_code += "\n"
		}
		go_file_name := "../../rmessageforgolang/"
		go_file_name += strings.TrimSuffix(filepath.Base(tmpFile), filepath.Ext(tmpFile))
		go_file_name += ".go"
		err = ioutil.WriteFile(go_file_name, []byte(g_code), 0644)
		if err != nil {
			fmt.Print(err)
		}
	}
}
func main() {

	datatype := map[string]string{}

	datatype["int32_t"] = "int"
	datatype["bool"] = "bool"
	datatype["string"] = "string"

	scanJsonFiles()
	return

}

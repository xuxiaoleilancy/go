package main

import (
	/*"bufio"
	"io"
	"strconv"
	*/
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"rpackage"
	"runtime"

	"strings"

	"github.com/astaxie/beego/config"
)

var (
	av_engine_sdk_dir    string
	release_profile_name string
	release_dest_path    string
	enable_log           bool

	lib_file_extension string

	release_result_errs []string
)

func add_result(str string) {
	release_result_errs = append(release_result_errs, str+"\n")
}

func copy_release_files_new() {

	printLogln("清理旧文件")
	err := os.RemoveAll(release_dest_path)
	if err != nil {
		add_result("清理旧文件 error:")
		fmt.Errorf("清理旧文件 error:", err)
	}
	printLogln("读取release_profile,", release_profile_name)
	releaseFileList, err := rpackage.Readline(release_profile_name)
	if err != nil {
		add_result("读取 release_profile error:")
		fmt.Errorf("读取 release_profile error:", err)
	}
	printLogln("发布文件列表=", releaseFileList)

	head_file_reg, _ := regexp.Compile(".h")
	lib_file_reg, _ := regexp.Compile(lib_file_extension)

	for _, str := range releaseFileList {
		printLogln("~~~~~~~~~~~~~~~~~~")
		if len(str) == 0 {
			printLogln("这是一个空行，忽略\n")
			continue
		}
		if strings.HasPrefix(str, "##") {
			printLogln("这是一个注释，忽略\n")
			continue
		}
		printLogln("file =", str)
		path_format := av_engine_sdk_dir + rpackage.Separator + strings.Replace(str, "/", rpackage.Separator, 0)
		file_path, _ := filepath.Abs(path_format)
		file_base_name := filepath.Base(file_path)
		file_rel_path := filepath.Dir(str)
		printLogln("file_abs_path =", file_path)
		printLogln("file_base_name =", file_base_name)
		printLogln("file_rel_path =", file_rel_path)

		file_release_abs_path := release_dest_path + rpackage.Separator
		file_release_full_name := release_dest_path
		head_match_result := head_file_reg.MatchString(str)

		if head_match_result {
			file_release_abs_path += "include"
			if file_rel_path != "." {
				file_release_abs_path += rpackage.Separator + file_rel_path
			}
		}
		lib_match_result := lib_file_reg.MatchString(str)
		if lib_match_result {
			file_release_abs_path += "lib"
			if file_rel_path != "." {
				file_release_abs_path += rpackage.Separator + file_rel_path
			}
		}

		//通配符，拷贝所有扩展名一致的文件
		if strings.HasPrefix(file_base_name, "*.") {
			printLogln("通配符，拷贝所有扩展名一致的文件,目录:", file_path, " 扩展名:", filepath.Ext(file_base_name))

			files, _ := rpackage.GetFiles(filepath.Dir(file_path), filepath.Ext(file_base_name))

			for _, tmpFile := range files {
				fmt.Println("需要拷贝的文件:", tmpFile)
				tmp_file_rel_path, _ := filepath.Rel(filepath.Dir(file_path), filepath.Dir(tmpFile))
				tmp_file_base_name := filepath.Base(tmpFile)
				tmp_release_full_name := file_release_abs_path + rpackage.Separator + tmp_file_rel_path + rpackage.Separator + tmp_file_base_name

				printLogln("复制文件", tmpFile, "===>", tmp_release_full_name)
				_, err = rpackage.CopyFileAdvanced(tmpFile, tmp_release_full_name)
				if err != nil {
					add_result("copy failed!" + tmpFile + "===>" + tmp_release_full_name)
					fmt.Printf("copy failed![%v]\n", err)
				}
			}
		} else { //指定文件拷贝
			file_release_full_name = file_release_abs_path + rpackage.Separator + file_base_name
			printLogln("复制文件", file_path, "===>", file_release_full_name)
			_, err = rpackage.CopyFileAdvanced(file_path, file_release_full_name)
			if err != nil {
				add_result("copy failed!" + file_path + "===>" + file_release_full_name)
				fmt.Printf("copy failed![%v]\n", err)
			}
		}

		printLogln("~~~~~~~~~~~~~~~~~~\n")
	}
}

func copy_release_files() {

	printLogln("清理旧文件")
	err := os.RemoveAll(release_dest_path)
	if err != nil {
		add_result("清理旧文件 error:")
		fmt.Errorf("清理旧文件 error:", err)
	}
	printLogln("读取release_profile,", release_profile_name)
	releaseFileList, err := rpackage.Readline(release_profile_name)
	if err != nil {
		add_result("读取 release_profile error:")
		fmt.Errorf("读取 release_profile error:", err)
	}
	printLogln("发布文件列表=", releaseFileList)

	head_file_reg, _ := regexp.Compile(".h")
	lib_file_reg, _ := regexp.Compile(lib_file_extension)

	for _, str := range releaseFileList {
		printLogln("~~~~~~~~~~~~~~~~~~")
		if len(str) == 0 {
			printLogln("这是一个空行，忽略\n")
			continue
		}
		if strings.HasPrefix(str, "##") {
			printLogln("这是一个注释，忽略\n")
			continue
		}
		printLogln("file =", str)
		path_format := av_engine_sdk_dir + rpackage.Separator + strings.Replace(str, "/", rpackage.Separator, 0)
		file_path, _ := filepath.Abs(path_format)
		file_base_name := filepath.Base(file_path)
		file_rel_path := filepath.Dir(str)
		printLogln("file_abs_path =", file_path)
		printLogln("file_base_name =", file_base_name)
		printLogln("file_rel_path =", file_rel_path)

		file_release_abs_path := release_dest_path + rpackage.Separator
		file_release_full_name := release_dest_path
		head_match_result := head_file_reg.MatchString(str)

		if head_match_result {
			file_release_abs_path += "include"
			if file_rel_path != "." {
				file_release_abs_path += rpackage.Separator + file_rel_path
			}
		}
		lib_match_result := lib_file_reg.MatchString(str)
		if lib_match_result {
			file_release_abs_path += "lib"
			if file_rel_path != "." {
				file_release_abs_path += rpackage.Separator + file_rel_path
			}
		}

		exist, err := rpackage.PathExists(file_release_abs_path)
		if err != nil {
			add_result("get dir error!" + file_release_abs_path)
			fmt.Printf("get dir error![%v]\n", err)
		}
		if !exist {
			printLogln("创建目录", file_release_abs_path)
			err := os.MkdirAll(file_release_abs_path, 0766)
			if err != nil {
				add_result("mkdir failed!" + file_release_abs_path)
				fmt.Printf("mkdir failed![%v]\n", err)
			}
		}

		//通配符，拷贝所有扩展名一致的文件
		if strings.HasPrefix(file_base_name, "*.") {
			printLogln("通配符，拷贝所有扩展名一致的文件,目录:", file_path, " 扩展名:", filepath.Ext(file_base_name))

			files, _ := rpackage.GetFiles(filepath.Dir(file_path), filepath.Ext(file_base_name))

			for _, tmpFile := range files {
				fmt.Println("需要拷贝的文件:", tmpFile)
				tmp_file_rel_path, _ := filepath.Rel(filepath.Dir(file_path), filepath.Dir(tmpFile))
				tmp_file_base_name := filepath.Base(tmpFile)
				tmp_release_full_name := file_release_abs_path + rpackage.Separator + tmp_file_rel_path + rpackage.Separator + tmp_file_base_name

				exist, err := rpackage.PathExists(filepath.Dir(tmp_release_full_name))
				if err != nil {
					add_result("get dir error!" + filepath.Dir(tmp_release_full_name))
					fmt.Printf("get dir error![%v]\n", err)
				}
				if !exist {
					printLogln("创建目录", filepath.Dir(tmp_release_full_name))
					err := os.MkdirAll(filepath.Dir(tmp_release_full_name), 0766)
					if err != nil {
						add_result("mkdir failed!" + filepath.Dir(tmp_release_full_name))
						fmt.Printf("mkdir failed![%v]\n", err)
					}
				}
				printLogln("复制文件", tmpFile, "===>", tmp_release_full_name)
				_, err = rpackage.CopyFile(tmpFile, tmp_release_full_name)
				if err != nil {
					add_result("copy failed!" + tmpFile + "===>" + tmp_release_full_name)
					fmt.Printf("copy failed![%v]\n", err)
				}
			}
		} else { //指定文件拷贝
			file_release_full_name = file_release_abs_path + rpackage.Separator + file_base_name
			printLogln("复制文件", file_path, "===>", file_release_full_name)
			_, err = rpackage.CopyFile(file_path, file_release_full_name)
			if err != nil {
				add_result("copy failed!" + file_path + "===>" + file_release_full_name)
				fmt.Printf("copy failed![%v]\n", err)
			}
		}

		printLogln("~~~~~~~~~~~~~~~~~~\n")
	}
}

func printLogln(a ...interface{}) (n int, err error) {
	if enable_log {
		return fmt.Println(a...)
	}

	return 0, nil
}

func initConfig() {

	os_verify()

	conf, err := config.NewConfig("ini", "release.ini")
	if err != nil {
		fmt.Println("release.ini new config failed, err:", err)
		add_result("release.ini new config failed")
		return
	}

	enable_log, err = conf.Bool("log::enable_log")
	av_engine_sdk_dir_tmp := conf.String("src::av_engine_sdk_path")
	release_profile_name_tmp := conf.String("src::release_profile")
	release_dest_path_tmp := conf.String("dest::release_dest_path")

	av_engine_sdk_dir, err = filepath.Abs(av_engine_sdk_dir_tmp)
	release_profile_name, err = filepath.Abs(release_profile_name_tmp)
	release_dest_path, err = filepath.Abs(release_dest_path_tmp)

	printLogln("日志开关=", enable_log)
	printLogln("sdk源码根目录=", av_engine_sdk_dir)
	printLogln("发布文件profile=", release_profile_name)
	printLogln("发布目录=", release_dest_path)
	printLogln("库文件扩展名=", lib_file_extension)
}
func os_verify() {
	printLogln("操作系统:", runtime.GOOS)
	switch runtime.GOOS {
	case "darwin":
		lib_file_extension = ".a"
	case "windows":
		lib_file_extension = ".lib"
	case "linux":
		lib_file_extension = ".so"
	default:
		break
	}
}

func result_varify() {
	if len(release_result_errs) == 0 {
		printLogln("(*￣︶￣)(*￣︶￣)(*￣︶￣)(*￣︶￣)(*￣︶￣)(*￣︶￣)")
		printLogln("(*￣︶￣)(*￣︶￣)(*￣︶￣)(*￣︶￣)(*￣︶￣)(*￣︶￣)")
		printLogln("(*￣︶￣)(*￣︶￣)(*￣︶￣)(*￣︶￣)(*￣︶￣)(*￣︶￣)")
		printLogln("可以去喝杯咖啡了")
	} else {
		printLogln("难过(ಥ﹏ಥ) 难过(ಥ﹏ಥ) 难过(ಥ﹏ಥ) 难过(ಥ﹏ಥ) 难过(ಥ﹏ಥ)")
		printLogln("难过(ಥ﹏ಥ) 难过(ಥ﹏ಥ) 难过(ಥ﹏ಥ) 难过(ಥ﹏ಥ) 难过(ಥ﹏ಥ)")
		printLogln("难过(ಥ﹏ಥ) 难过(ಥ﹏ಥ) 难过(ಥ﹏ಥ) 难过(ಥ﹏ಥ) 难过(ಥ﹏ಥ)")
		printLogln("哭也没用啊，赶紧解决问题吧")
		printLogln("↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓")

		fmt.Println(release_result_errs)
	}
}
func main() {
	enable_log = true
	flag.Parse()
	initConfig()
	copy_release_files_new()
	result_varify()
}

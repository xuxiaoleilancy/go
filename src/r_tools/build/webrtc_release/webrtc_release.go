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

	printLogln("Clean OLD FILES")
	err := os.RemoveAll(release_dest_path)
	if err != nil {
		add_result("Clean OLD FILES error: " + err.Error())
		fmt.Errorf("Clean OLD FILES error:", err)
	}
	printLogln("READ release_profile,", release_profile_name)
	releaseFileList, err := rpackage.Readline(release_profile_name)
	if err != nil {
		add_result("READ release_profile error:")
		fmt.Errorf("READ release_profile error:", err)
	}
	printLogln("RELEASE FILE LIST=", releaseFileList)

	for _, str := range releaseFileList {
		printLogln("~~~~~~~~~~~~~~~~~~")
		if len(str) == 0 {
			printLogln("THIS IS A BLANK LINE , IGNORE IT\n")
			continue
		}
		if strings.HasPrefix(str, "##") {
			printLogln("THIS IS A COMMETN , IGNORE IT\n")
			continue
		}
		printLogln("file =", str)
		path_format := av_engine_sdk_dir + rpackage.Separator + strings.Replace(str, "/", rpackage.Separator, 0)
		file_path, _ := filepath.Abs(path_format)
		file_base_name := filepath.Base(file_path)
		file_ext_name := filepath.Ext(file_path)
		file_rel_path := filepath.Dir(str)
		printLogln("file_abs_path =", file_path)
		printLogln("file_base_name =", file_base_name)
		printLogln("file_rel_path =", file_rel_path)

		file_release_abs_path := release_dest_path + rpackage.Separator
		file_release_full_name := release_dest_path
		head_match_result := strings.EqualFold(file_ext_name, ".h")
		printLogln("head_match_result =", head_match_result)

		if head_match_result {
			file_release_abs_path += "include"
			if file_rel_path != "." {
				file_release_abs_path += rpackage.Separator + file_rel_path
			}
		}
		lib_match_result := strings.EqualFold(file_ext_name, lib_file_extension)
		printLogln("lib_match_result =", lib_match_result)
		if lib_match_result {
			file_release_abs_path += "lib"
			//			if file_rel_path != "." {
			//				file_release_abs_path += rpackage.Separator + file_rel_path
			//			}
		}

		//通配符，拷贝所有扩展名一致的文件
		if strings.HasPrefix(file_base_name, "*.") {
			printLogln("Wildcard, Copies all files with the same extension, Path:", file_path, " Extension:", filepath.Ext(file_base_name))

			files, _ := rpackage.GetFiles(filepath.Dir(file_path), filepath.Ext(file_base_name))

			for _, tmpFile := range files {
				fmt.Println("Files need to be copied:", tmpFile)
				tmp_file_rel_path, _ := filepath.Rel(filepath.Dir(file_path), filepath.Dir(tmpFile))
				tmp_file_base_name := filepath.Base(tmpFile)
				tmp_release_full_name := file_release_abs_path + rpackage.Separator + tmp_file_rel_path + rpackage.Separator + tmp_file_base_name
				if lib_match_result {
					tmp_release_full_name = file_release_abs_path + rpackage.Separator + tmp_file_base_name
				}
				printLogln("COPY FILE", tmpFile, "===>", tmp_release_full_name)
				_, err = rpackage.CopyFileAdvanced(tmpFile, tmp_release_full_name)
				if err != nil {
					add_result("copy failed!" + tmpFile + "===>" + tmp_release_full_name)
					fmt.Printf("copy failed![%v]\n", err)
				}
			}
		} else { //指定文件拷贝
			file_release_full_name = file_release_abs_path + rpackage.Separator + file_base_name
			printLogln("COPY FILE", file_path, "===>", file_release_full_name)
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

	printLogln("Clean OLD FILES")
	err := os.RemoveAll(release_dest_path)
	if err != nil {
		add_result("REMOVE PATH " + release_dest_path + " error:" + err.Error())
		fmt.Errorf("CLEAN OLD FILES error:", err)
	}
	printLogln("READ release_profile,", release_profile_name)
	releaseFileList, err := rpackage.Readline(release_profile_name)
	if err != nil {
		add_result("READ release_profile error:")
		fmt.Errorf("READ release_profile error:", err)
	}
	printLogln("RELEASE FILE LIST=", releaseFileList)

	head_file_reg, _ := regexp.Compile(".h")
	lib_file_reg, _ := regexp.Compile(lib_file_extension)

	for _, str := range releaseFileList {
		printLogln("~~~~~~~~~~~~~~~~~~")
		if len(str) == 0 {
			printLogln("THIS IS A BLANK LINE , IGNORE IT\n")
			continue
		}
		if strings.HasPrefix(str, "##") {
			printLogln("THIS IS A COMMETN , IGNORE IT\n")
			continue
		}
		printLogln("file =", str)
		path_format := av_engine_sdk_dir + rpackage.Separator + strings.Replace(str, "/", rpackage.Separator, 0)
		file_path, _ := filepath.Abs(path_format)
		file_base_name := filepath.Base(file_path)
		file_rel_path := filepath.Dir(str)
		printLogln("FILE_ABS_PATH =", file_path)
		printLogln("FILE_BASE_NAME =", file_base_name)
		printLogln("FILE_REL_PATH =", file_rel_path)

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
			//			if file_rel_path != "." {
			//				file_release_abs_path += rpackage.Separator + file_rel_path
			//			}
		}

		exist, err := rpackage.PathExists(file_release_abs_path)
		if err != nil {
			add_result("get dir error!" + file_release_abs_path)
			fmt.Printf("get dir error![%v]\n", err)
		}
		if !exist {
			printLogln("mkdir ", file_release_abs_path)
			err := os.MkdirAll(file_release_abs_path, 0766)
			if err != nil {
				add_result("mkdir failed!" + file_release_abs_path + " error:" + err.Error())
				fmt.Printf("mkdir failed![%v]\n", err)
			}
		}

		//通配符，拷贝所有扩展名一致的文件
		if strings.HasPrefix(file_base_name, "*.") {
			printLogln("Wildcard, Copies all files with the same extension, Path:", file_path, " Extension:", filepath.Ext(file_base_name))

			files, _ := rpackage.GetFiles(filepath.Dir(file_path), filepath.Ext(file_base_name))

			for _, tmpFile := range files {
				fmt.Println("Files need to be copied:", tmpFile)
				tmp_file_rel_path, _ := filepath.Rel(filepath.Dir(file_path), filepath.Dir(tmpFile))
				tmp_file_base_name := filepath.Base(tmpFile)
				tmp_release_full_name := file_release_abs_path + rpackage.Separator + tmp_file_rel_path + rpackage.Separator + tmp_file_base_name

				if lib_match_result {
					tmp_release_full_name = file_release_abs_path + rpackage.Separator + tmp_file_base_name
				}

				exist, err := rpackage.PathExists(filepath.Dir(tmp_release_full_name))
				if err != nil {
					add_result("get dir error!" + filepath.Dir(tmp_release_full_name))
					fmt.Printf("get dir error![%v]\n", err)
				}
				if !exist {
					printLogln("MKDIR ", filepath.Dir(tmp_release_full_name))
					err := os.MkdirAll(filepath.Dir(tmp_release_full_name), 0766)
					if err != nil {
						add_result("mkdir failed!" + filepath.Dir(tmp_release_full_name))
						fmt.Printf("mkdir failed![%v]\n", err)
					}
				}
				printLogln("COPY FILE", tmpFile, "===>", tmp_release_full_name)
				_, err = rpackage.CopyFile(tmpFile, tmp_release_full_name)
				if err != nil {
					add_result("copy failed!" + tmpFile + "===>" + tmp_release_full_name)
					fmt.Printf("copy failed![%v]\n", err)
				}
			}
		} else { //指定文件拷贝
			file_release_full_name = file_release_abs_path + rpackage.Separator + file_base_name
			printLogln("COPY FILE ", file_path, "===>", file_release_full_name)
			_, err = rpackage.CopyFile(file_path, file_release_full_name)
			if err != nil {
				add_result("COPY FAILED!" + file_path + "===>" + file_release_full_name)
				fmt.Printf("COPY FAILED![%v]\n", err)
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

	printLogln("enable_log=", enable_log)
	printLogln("av_engine_sdk_dir=", av_engine_sdk_dir)
	printLogln("release_profile_name=", release_profile_name)
	printLogln("release_dest_path=", release_dest_path)
	printLogln("lib_file_extension=", lib_file_extension)
}
func os_verify() {
	printLogln("OS:", runtime.GOOS)
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

func result_varify() int {
	if len(release_result_errs) == 0 {
		printLogln("(*￣︶￣)(*￣︶￣)(*￣︶￣)(*￣︶￣)(*￣︶￣)(*￣︶￣)")
		printLogln("(*￣︶￣)(*￣︶￣)(*￣︶￣)(*￣︶￣)(*￣︶￣)(*￣︶￣)")
		printLogln("(*￣︶￣)(*￣︶￣)(*￣︶￣)(*￣︶￣)(*￣︶￣)(*￣︶￣)")
		printLogln("Drink a cup of tea")

	} else {
		printLogln("SAD(ಥ﹏ಥ) SAD(ಥ﹏ಥ) SAD(ಥ﹏ಥ) SAD(ಥ﹏ಥ) SAD(ಥ﹏ಥ)")
		printLogln("SAD(ಥ﹏ಥ) SAD(ಥ﹏ಥ) SAD(ಥ﹏ಥ) SAD(ಥ﹏ಥ) SAD(ಥ﹏ಥ)")
		printLogln("SAD(ಥ﹏ಥ) SAD(ಥ﹏ಥ) SAD(ಥ﹏ಥ) SAD(ಥ﹏ಥ) SAD(ಥ﹏ಥ)")
		printLogln("iT'S NO USE CRYING, JUST SOLVE THE PROBLEM")
		printLogln("↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓")

		fmt.Println(release_result_errs)
	}

	return len(release_result_errs)
}
func main() {
	enable_log = true
	flag.Parse()
	initConfig()
	copy_release_files_new()
	os.Exit(result_varify())
}

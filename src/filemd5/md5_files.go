package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
)

func Md5SumFile(file string) (value [md5.Size]byte, err error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	value = md5.Sum(data)
	return
}
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./md5file yourfile")
		return
	}
	md5Value, err := Md5SumFile(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%x %s\n", md5Value, os.Args[1])
}

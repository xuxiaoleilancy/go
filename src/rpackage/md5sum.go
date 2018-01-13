package rpackage

import (
	"crypto/md5"
	"io/ioutil"
)

func Md5SumFile(file string) (value [md5.Size]byte, err error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	value = md5.Sum(data)
	return
}

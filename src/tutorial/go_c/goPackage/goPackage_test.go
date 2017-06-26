package goPackage

import (
	"testing"
)

func Test_Add(t *testing.T) {
	if i := Add(6, 2); i != 8 { //try a unit test on function
		t.Error("add函数测试没通过") // 如果不是如预期的那么就报错
	} else {
		t.Log("第一个测试通过了") //记录一些你期望记录的信息
	}
}

func Test_Add_2(t *testing.T) {
	t.Error("就是不通过")
}

package test

import (
	"testing"
)

/**
*单元测试文件名必须以test.go结尾，例如calc_test.go
*单元测试函数名签名必须是func Testxxx(t *testing.T){}
*跳转到单元测试所在path，使用go test -v执行单元测试
*/

func TestAdd(t *testing.T){
	r:=add(2, 3)
	if r!=5{
		t.Fatalf("add(2,3) error, expect:%d,actual:%d",5, r)
	}
	t.Logf("test add success")
}
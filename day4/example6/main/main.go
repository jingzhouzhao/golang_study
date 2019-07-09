package main

import (
	"fmt"
)

func test() {
	//string是不可变的，不能直接修改
	var str string = "hello world"
	//直接编译错误
	//str[0]='a'
	//可以通过rune来修改
	s := []rune(str)
	s[0] = 'a'
	str = string(s)
	fmt.Println(str)

	//也可以通过byte数组来修改，但是一个中文字符是占三个字节的，所以修改可能会有问题。
}

func main() {
	test()
}

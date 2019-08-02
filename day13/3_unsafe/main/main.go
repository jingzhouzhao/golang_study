package main

import (
	"fmt"
	"unsafe"
)

type Program struct{
	name string
	language string
}

type Program2 struct{
	name string
	size int
	language string
}

func main() {
	//使用go底层unsafe来操作结构体
	//因为go底层结构体分配内存是连续的
	//unsafe 包定义了 Pointer 和三个函数：
	/*
	type ArbitraryType int
	type Pointer *ArbitraryType

	func Sizeof(x ArbitraryType) uintptr
	func Offsetof(x ArbitraryType) uintptr
	func Alignof(x ArbitraryType) uintptr
	*/

	p:=Program{name:"utest",language:"go"}
	fmt.Println(p)
	name:=(*string)(unsafe.Pointer(&p))
	*name="test"
	//Offsetof 获取地址偏移量
	lang:=(*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p))+unsafe.Offsetof(p.language)))
	*lang="Golang"
	fmt.Println(p)

	p2:=Program2{name:"haha",size:18,language:"xx"}
	fmt.Println(p2)
	name2:=(*string)(unsafe.Pointer(&p2))
	*name2="test2"
	//Sizeof 获得变量地址大小，int 与 size 变量一致即可，因为int不管值多少大小是一致的
	lang2:=(*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p2))+unsafe.Sizeof(int(0))+unsafe.Sizeof(string(""))))
	*lang2="Golang2"
	fmt.Println(p2)

}
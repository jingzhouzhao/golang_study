package main

import (
	"fmt"
)

func test(){
	//切片是数组的一个引用（底层实现还是数组），长度是可变的
	//从数组切片获得一个slice
	var a[5]int = [...]int{1,2,3,4,5}
	//[start:end]，包含start下标，不包含end下标
	s:=a[1:3]
	fmt.Println(s)
	//如果start为0，end为数组len-1，那么可以省略
	s1:=a[:3]
	fmt.Println(s1)
	s1=a[3:]
	fmt.Println(s1)

	//slice 可以继续缩小(resize)
	s1=s1[:1]
	fmt.Println(s1)

	//slice ，cap容量，len当前长度。cap>=len
	fmt.Printf("cap:%d,len:%d \n", cap(s1),len(s1))
	fmt.Printf("扩容之前s1[0]:%p\n", &s1[0])
	s1 = append(s1,1)
	s1 = append(s1,1)
	s1 = append(s1,1)
	s1 = append(s1,1)
	s1 = append(s1,1)
	s1 = append(s1,1)
	fmt.Printf("扩容之后s1[0]:%p\n", &s1[0])
	//当切片容量不足时会自动扩容，copy原始的值，到一块新的数组中，默认容量*2
	fmt.Printf("cap:%d,len:%d \n", cap(s1),len(s1))

	a1:=[]int{1,2,3}
	a2:=[]int{4,5,6}
	//合并两个切片
	a1 =append(a1, a2...)
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Printf("a2[0]:%p,%d\n",&a2[0],a2[0])
	fmt.Printf("a1[3]:%p,%d\n",&a1[3],a1[3])

	var b1 []int = []int{1,2,3,4,5,6}
	b2:=make([]int,1)
	//copy 不会扩容
	copy(b2,b1)
	fmt.Println(b2)
}

func test1(){
	//字符串底层是一个byte数组，所以也可以用来切片
	var str string = "hello world"
	str1:=str[:5]
	str2:=str[6:]
	fmt.Println(str1)
	fmt.Println(str2)
}

func main(){
	//test()
	test1()
}
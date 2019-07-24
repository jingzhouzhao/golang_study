package main

import (
	"fmt"
)


func test(){
	//slice 与array 在形式上的区别
	//[]int slice
	//[5]int array

	//new 创建一个指针
	s1:=new([]int)
	//make 用在引用类型初始化创建
	s2:=make([]int,10)
	fmt.Println(s1)
	fmt.Println(s2)
	//如果给没有初始化的slice指针赋值会报错
	//(*s1)[0]=100
	//通过make给引用类型初始化(指针是引用类型)
	*s1=make([]int,5)
	(*s1)[0]=100
	fmt.Println(s1)
}

func main(){
	test()
}
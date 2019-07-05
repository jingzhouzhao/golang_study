package main

import (
	"fmt"
)

func main()  {
	
	i:=100
	//匿名函数，可以对比example11，看看defer一个语句和一个匿名函数的差异。
	defer func(){
		fmt.Println("退出时i=", i)
	}()
	//匿名函数
	fn:=func(){
		i++
	}
	fn()
	i++
	fmt.Println(i)
}
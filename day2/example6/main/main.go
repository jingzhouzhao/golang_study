package main

import (
	"fmt"
)

func modify(a int){
	a = 5
}
//定义参数为指针
func modify1 (a *int){
	//修改指针值
	*a = 5
}

func main()  {
	var a int = 100;
	var b chan int = make(chan int ,1)
	fmt.Println("值类型：", a)
	fmt.Println("引用类型：", b)

	modify(a)
	fmt.Println("值类型：", a)
	//传递指针
	modify1(&a)
	fmt.Println("引用类型：", a)
	
}
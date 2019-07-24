package main

import (
	"fmt"
)

func main()  {
	var a int = 1
	//& 取址符
	fmt.Println("变量a的指向的内存地址是：",&a)
	var b int = 2
	fmt.Println("变量b的指向的内存地址是：",&b)
	//指针
	var c *int=&a
	//改变指针指向的值
	*c = *&b
	fmt.Println("*c=",*c)
	fmt.Println("a=",a)
	fmt.Println("变量a的指向的内存地址是：",&a)

	/*总结：
	&变量 ，获取的是变量的内存地址
	*type ，定义这是一个指针（地址变量）
	*type = ?  改变地址所指向的值
	*/

	
}
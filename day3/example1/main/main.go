package main

import (
	"fmt"
)

func main() {
	//Scanf 用法
	// var a int
	// fmt.Scanf("%d", a)
	// fmt.Println("a=", a)
	//Scanf需要传递引用
	var b int
	fmt.Scanf("%d", &b)
	fmt.Println("b=", b)

	
}

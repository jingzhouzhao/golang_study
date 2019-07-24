package main

import (
	"fmt"
)

func main()  {
	//defer示例
	var i = 0
	defer fmt.Println(i)
	defer fmt.Print("退出时i为")
	i++
	fmt.Println(i)
	return
}
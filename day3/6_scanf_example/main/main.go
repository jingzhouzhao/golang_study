package main

import (
	"time"
	"math/rand"
	"fmt"
)

func main()  {
	
	var input int
	rand.Seed(time.Now().UnixNano())
	var rand int = rand.Intn(100)
	for{
		fmt.Println("猜猜是哪个数字[0-100]：")
		//输入数字然后敲下回车，缓冲区中其实同时进入了数字和\n所以此处是%d\n
		fmt.Scanf("%d\n", &input)
		fmt.Println("你输入的是：",input)
		switch {
		case rand==input:
			fmt.Println("猜对了！")
			return
		case rand>input:
			fmt.Println("小了一点！")
		case rand<input:
			fmt.Println("大了一点！")
		}
	}
}
package main

import (
	"fmt"
)
/*数据类型及转换*/
func main()  {
	
	var str = "hello  \n golang"
	fmt.Println("str：",str)
	//使用``的字符串中的特殊符号不会转义，直接原样输出
	str = `hello \n world`
	fmt.Println("str:",str)

	var a uint8=255;
	var b int8=127
	var c int8=-128
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	var a1 int8 = 100;
	//int8与int16是不同的类型，无法直接赋值
	//var a2 int16 = a1;
	var a2 int16 = int16(a1)
	fmt.Println(a2)
	//
	var a3 int32 = 10;
	//不同的类型不能直接操作
	var a4 int32 = a3+int32(a1);
	fmt.Println(a4)

	//可以直接操作常数，此处5自动推导为a4的类型。
	a4= a4+5;
	fmt.Println(a4)

}
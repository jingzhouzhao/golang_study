package main

import (
	"strings"
	"fmt"
)
/*闭包：一个函数与其相关的引用环境组合而成的实体*/
func test() func(int)int{
	var x int
	return func (d int) int{
		x+=d
		return x
	}
}

func makeSuffix(suffix string)func(string) string{
	return func(name string)string{
		if !strings.HasSuffix(name, suffix) {
			name+=("."+suffix)
		}
		return name
	}
}

func main()  {
	fn:=test()
	fmt.Println(fn(1))
	fmt.Println(fn(10))
	fmt.Println(fn(100))

	/*output:
	1
	11
	111
	说明变量x并不会在每次调用初始化。
	x是与fn绑定的
	*/

	fn_jpg := makeSuffix("jpg")
	fn_bmp := makeSuffix("bmp")
	fmt.Println(fn_jpg("123"))
	fmt.Println(fn_jpg("456"))
	fmt.Println(fn_bmp("abc"))
	fmt.Println(fn_bmp("dfds"))
	fmt.Println(fn_jpg("789"))
	fmt.Println(fn_jpg("000.jpg"))
	
}
package main

import (
	"fmt"
)
/*可变参数示例*/
func add(a, b int, args ...int) (sum int) {
	sum = a + b
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return
}

func concat(a,b string,args...string)(result string){
	result=a+b
	for i := 0; i < len(args); i++ {
		result+=args[i]
	}
	return
}

func main() {
	fmt.Printf("a+b=%d\n", add(1, 2))
	fmt.Printf("a+b+c=%d\n", add(1, 2, 3))
	fmt.Printf("a+b+c+d=%d\n", add(1, 2, 3, 4))
	fmt.Printf("s1+s2=%s\n", concat("hello", "golang"))
	fmt.Printf("s1+s2+s3=%s\n", concat("hello", "golang","你好"))
	fmt.Printf("s1+s2+s3+s4=%s\n", concat("hello", "golang","你好","够浪"))
}

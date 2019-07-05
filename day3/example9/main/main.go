package main

import (
	"fmt"
)

type operator func(a,b int)int

func add(a,b int) int{
	return a+b
}
func sub(a,b int) int{
	return a-b
}

func calc(op operator,a int ,b int)int{
	return op(a,b)
}

func main()  {
	fmt.Printf("a+b=%d\n", calc(add, 1, 3))
	fmt.Printf("a-b=%d\n", calc(sub, 3, 1))
	
}
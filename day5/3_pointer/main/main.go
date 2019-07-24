package main

import (
	"fmt"
)

func change(a *int){
	fmt.Println(a)
	t:=30
	b:=&t
	//这里只是将变量a指向了另外一个地址b，而b指向的是地址&t，不会改变函数外的a
	a = b
	fmt.Println("inner a",*a)
}

func change2(a **int){
	fmt.Println(a)
	t:=30
	b:=&t
	*a = b
	fmt.Println("inner a",*a)
}

func main(){
	// var a int = 20
	// var ptr = &a
	// var ptrr = &ptr

	// fmt.Println(ptr)
	// fmt.Println(ptrr)

	// fmt.Println(*ptr)
	// fmt.Println(*ptrr)
	// fmt.Println(**ptrr)

	// a=30
    // fmt.Println("a changed--------------")
	// fmt.Println(*ptr)
	// fmt.Println(*ptrr)
	// fmt.Println(**ptrr)

	// **ptrr = 40 

	// fmt.Println("**ptrr changed--------------")
	// fmt.Println(*ptr)
	// fmt.Println(*ptrr)
	// fmt.Println(**ptrr)

	var a int = 20
	change(&a)
	fmt.Println(a)

	fmt.Println("change2")
	var aptr = &a
	aptrr := &aptr
	change2(aptrr)
	fmt.Println(**aptrr)
	fmt.Println(a)

	
}
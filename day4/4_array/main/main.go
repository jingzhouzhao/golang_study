package main

import (
	"fmt"
)

func test(){
	//定义一个数组
	var a[10]int
	b:=a
	b[0]=100
	fmt.Println(b)
	//a的值未发生改变，说明数组是值类型
	fmt.Println(a)
}

func test1(arr [10]int){
	arr[0]=100
}

func test2(arr *[10]int){
	//更改数组指针的值
	(*arr)[0]=100
}

func test3(){
	//数组初始化的几种方式
	var a1 [5]int =[5]int{1,2,3,4,5}
	fmt.Println(a1)
	var a2 = [6]int{1,2,3,4,5,6}
	fmt.Println(a2)
	var a3 = [...]int{1,2,3}
	fmt.Println(a3)
	var a4 = [5]string{2:"hello",4:"world"}
	fmt.Println(a4)
	var a5 = [...]string{2:"hello",40:"golang"}
	fmt.Println(a5)
}

func test4(){
	var a [2][5]int = [...][5]int{{1,2,3,4,5},{6,7,8,9,10}}

	for row,v :=range a{
		for col,v1 :=range v{
			fmt.Printf("%d-%d=%d \t",row,col,v1)
		}
		fmt.Println()
	}

}

func main()  {

	test()

	/*
	长度是数组类型的一部分，长度为5的与长度为10的不是一个类型，所以不能传递
	var arr [5] int
	*/
	var arr [10] int
	test1(arr)
	//arr的值不会改变，数组是值传递
	fmt.Println(arr)
	//如果需要在函数中改变数组的值，需要传递地址，并在函数中操作指针的值
	test2(&arr)
	fmt.Println(arr)
	fmt.Println("--------")
	test3()

	test4()
}
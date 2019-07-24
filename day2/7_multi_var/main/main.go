package main

import (
	"fmt"
)

/*交换变量值1*/
func swap(a int ,b int) (int ,int){
	return b,a
}
/*交换变量值2*/
func swap1(a *int ,b *int){
	*a = *a+*b
	*b = *a-*b
	*a = *a-*b
}
/*交换变量值3*/
func swap2(a *int ,b*int){
	*a = *a ^ *b
	*b = *a ^ *b 
	*a = *a ^ *b
}


func main()  {
	first:=100
	second:=200
	//直接交换
	first,second = second,first
	fmt.Println("first:",first,"second:",second)
	first,second = swap(first, second);
	fmt.Println("first:",first,"second:",second)
	swap1(&first, &second);
	fmt.Println("first:",first,"second:",second)
	swap2(&first, &second);
	fmt.Println("first:",first,"second:",second)
}

package main

import (
	"fmt"
)

func main(){
	str:="字符串"
	fmt.Println(str)
	fmt.Printf("%T",str)
	fmt.Println("")
	fmt.Printf("%b",2)
	fmt.Println("")
	fmt.Printf("%o",10)
	fmt.Println("")
	fmt.Printf("%X",15)
	fmt.Println("")
	fmt.Printf("%+d",-15)
	fmt.Println("")
	fmt.Printf("%5d",1234567)
}
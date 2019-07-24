package main

import (
	"fmt"
	"go_dev/day1/4_func/calc"
)

func main(){
	a:=calc.Add(1,3)
	b:=calc.Sub(1,3)
	fmt.Println("add:",a,"sub:",b)
}
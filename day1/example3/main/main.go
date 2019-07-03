package main

import (
	"fmt"
	"go_dev/day1/goroute_example/goroute"
)

func main(){
	pipe:=make(chan int,1)
	goroute.Add(1,2,pipe)
	sum:=<- pipe
	fmt.Println("sum=",sum)
}
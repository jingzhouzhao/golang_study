package main

import (
	"time"
	"fmt"
)
//只写chan<-
func Write(ch chan<- int){
	for i := 0; i < 100; i++ {
		ch <- i
	}
}

//只读<-chan
func Read(ch <-chan int){
	var b int
	for{
		b= <-ch
		fmt.Println(b)
	}
}

func main(){
	var ch chan int = make(chan int ,10)
	go Write(ch)
	go Read(ch)
	time.Sleep(5*time.Second)
}
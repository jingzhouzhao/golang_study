package main

import (
	"fmt"
	"time"
)

func test() {
	//如果不捕获异常整个程序会挂掉。正确做法是在每个goroute中都捕获异常
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic:", err)
		}
	}()
	var m map[string]int
	m["test"] = 1
}

func calc(ch chan<- int) {
	for {
		fmt.Println("calc...")
		time.Sleep(time.Second)
	}
}

func main() {
	ch:=make(chan int,1)
	for i := 0; i < 2; i++ {
		go calc(ch)
	}
	go test()

	<-ch
}

package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int = make(chan int, 10)
	ch2 := make(chan int, 10)
	i := 1
	go func() {
		for {
			ch <- i
			time.Sleep(time.Second)
			ch2 <- i*i
			time.Sleep(time.Second)
			i++
		}
	}()
	
	//多个chan用for
	for {
		select {
		case v := <-ch:
			fmt.Println("ch:", v)
		case v := <-ch2:
			fmt.Println("ch2:", v)
		default:
			fmt.Println("do something others....")
			time.Sleep(time.Second)
		}
	}
}

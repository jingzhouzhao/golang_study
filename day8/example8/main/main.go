package main

import (
	"fmt"
	"time"
)

//定时器
func testTicker() {
	//每隔一秒执行一次
	t := time.NewTicker(time.Second)
	for v := range t.C {
		fmt.Println(v)
	}
}

func queryDB(ch chan<- int) {
	//模拟查询两秒
	time.Sleep(2 * time.Second)
	ch <- 1
}
func testTimeout() {
	ch := make(chan int, 1)
	go queryDB(ch)
	t := time.NewTicker(time.Second) //推荐使用，需要及时Stop，否则可能内存泄露
	defer t.Stop()
	select {
	case v := <-ch:
		fmt.Println("result:", v)
		return
	case <-t.C:
		fmt.Println("timeout")
		return
	}
}

func testTimeout2() {
	ch := make(chan int, 1)
	go queryDB(ch)
	t := time.After(time.Second) //不推荐使用
	select {
	case v := <-ch:
		fmt.Println("result:", v)
		return
	case <-t:
		fmt.Println("timeout")
		return
	}
}

func main() {
	//testTicker()
	testTimeout()
}

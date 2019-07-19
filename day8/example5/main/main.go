package main

import (
	"fmt"
	"math"
	"runtime"
)

func Calc(taskChan chan int, retChan chan int, exitChan chan bool) {
	for v := range taskChan {
		flag := true
		//判断素数
		for i := 2; i <= int(math.Ceil(math.Sqrt(float64(v)))); i++ {
			if v%i == 0 && v != i {
				flag = false
				break
			}
		}
		if flag {
			retChan <- v
		}
	}
	//退出
	exitChan <- true
}

func main() {
	var numCpu int = runtime.NumCPU()
	var taskChan chan int = make(chan int, 1000)
	var retChan chan int = make(chan int, 1000)
	var exitChan chan bool = make(chan bool, numCpu)
	go func() {
		for i := 0; i < 1000; i++ {
			taskChan <- i
		}
		close(taskChan)
	}()
	
	for i := 0; i < numCpu; i++ {
		go Calc(taskChan, retChan, exitChan)
	}
	go func() {
		//通过chan来判断所有goroutine是否结束
		for i := 0; i < numCpu; i++ {
			<-exitChan
		}
		fmt.Println("ready to close chan!")
		close(retChan)
		fmt.Println("retChan closed!")
	}()

	for v := range retChan {
		fmt.Println(v)
	}

}

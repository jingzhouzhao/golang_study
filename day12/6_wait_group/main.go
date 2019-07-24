package main

import (
	"math/rand"
	"time"
	"fmt"
	"sync"
)

func main(){

	var num_goroutine int = 10

	wg:=sync.WaitGroup{}
	wg.Add(num_goroutine)

	for i := 0; i < num_goroutine; i++ {
		go func (i int)  {
			time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond)
			wg.Done()
			fmt.Printf("goroute[%d] done.\n",i)
		}(i)
	}
	wg.Wait()
	fmt.Println("main exit!")
}
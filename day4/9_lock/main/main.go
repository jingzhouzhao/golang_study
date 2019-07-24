package main

import (
	"sync/atomic"
	"sync"
	"math/rand"
	"time"
	"fmt"
)

var lock sync.Mutex
var rwLock sync.RWMutex

func test(){
	var a map[int]int= make(map[int]int)
	a[0]=1
	a[1]=1
	a[2]=1

	for i := 0; i < 2; i++ {
		go func (){
			rand.Seed(time.Now().UnixNano())
			lock.Lock()
			a[0]=rand.Intn(100)
			lock.Unlock()
		}()
	}
	time.Sleep(time.Second)
	lock.Lock()
	fmt.Println(a)
	lock.Unlock()
}

func test2(){
	var a map[int]int= make(map[int]int)
	var count int32 
	a[0]=1
	a[1]=1
	a[2]=1

	for i := 0; i < 2; i++ {
		go func (){
			rwLock.Lock()
			rand.Seed(time.Now().UnixNano())
			a[0]=rand.Intn(100)
			rwLock.Unlock()
		}()
	}
	for i := 0; i < 100; i++ {
		go func (){
			rwLock.RLock()
			fmt.Println(a)
			rwLock.RUnlock()
			atomic.AddInt32(&count, 1)
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(atomic.LoadInt32(&count))

}

func main(){
	//test()
	test2()
}
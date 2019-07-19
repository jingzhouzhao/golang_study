package main

import (
	"time"
	"sync"
	"sort"
	"fmt"
)

type Task struct{
	n int
}

var (
	m = make(map[int]uint64)
	lock sync.Mutex
)

func calc(t *Task){
	sum:=1
	for i := 1; i <=t.n; i++ {
		sum*=i
	}
	lock.Lock()
	m[t.n]=uint64(sum)
	defer lock.Unlock()
}

func main(){
	for i := 1; i <=1000; i++ {
		t:=&Task{n:i,}
		go calc(t)
	}
	time.Sleep(5*time.Second)
	var keys []int
	lock.Lock()
	for k,_:=range m{
		keys = append(keys,k)
	}
	sort.Ints(keys)
	for _,k:=range keys{
		fmt.Printf("%d:%d\n", k,m[k])
	}
	defer lock.Unlock()
}
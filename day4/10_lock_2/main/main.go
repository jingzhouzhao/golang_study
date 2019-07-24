package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex

//锁不可重复加锁
func test1(){
	//死锁
	lock.Lock()
	lock.Lock()
}

//出现异常锁不会解锁
func test2(){
	defer func(){
		if err:=recover();err!=nil{
			//再次加锁，死锁
			//lock.Lock()
			lock.Unlock()
			fmt.Println(err)
		}
	}()
	lock.Lock()
	var m map[string]int
	m["test"]=1
	lock.Unlock()
	fmt.Println("已经出现异常")
}

const (
	a = 1+iota
	b
	c
)

const (
	d = 1<<iota
	e
	f
)


func main(){
	//test1()
	//test2()
	fmt.Println(a,b,c,d,e,f)
}
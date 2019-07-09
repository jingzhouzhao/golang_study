package main

import (
	"errors"
	"runtime/debug"
	"time"
	"fmt"
)

func initConfig() error {
	return errors.New("init config failed")
}

func test(){
	/**recover 恢复，捕获异常*/
	defer func(){
		if err:=recover(); err!=nil{
			//打印错误信息
			fmt.Println(err)
			//方便调试可以答应错误堆栈
			debug.PrintStack()
			//也可获取堆栈，打印到日志，或者存储到elk等日志平台
			//debug.Stack()
		}
	}()
	b:=0
	a:=100/b
	fmt.Println(a)
}

func main()  {
	err:=initConfig()
	if err!=nil{
		//抛出异常，可结合recover做捕获
		panic(err)
	}
	for{
		test()
		time.Sleep(time.Second)
	}
}
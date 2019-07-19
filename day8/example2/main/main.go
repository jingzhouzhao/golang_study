package main

import (
	"fmt"
)

type User struct{
	UserName string
}

func main(){

	var intChan chan int = make(chan int,5)
	intChan <- 10

	var a int 
	a = <-intChan
	fmt.Println(a)

	var structChan chan *User = make(chan *User,5)
	var user = User{UserName:"张三"} 
	structChan <- &user
	var user1 *User = <-structChan
	fmt.Println(user1)
	user1.UserName = "李四"
	fmt.Print(user)
	
}
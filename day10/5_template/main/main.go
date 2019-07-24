package main

import (
	"fmt"
	"os"
	"text/template"
)

type User struct{
	UserName string
	age int
}

func main(){
	r,err:=template.ParseFiles("C:\\Users\\zhaojz\\go\\src\\go_dev\\day10\\example5\\main\\index.html")
	if err!=nil{
		fmt.Println(err)
		return
	}
	u:=User{
		UserName:"张三",
		age:18,
	}
	r.Execute(os.Stdout, u)
}

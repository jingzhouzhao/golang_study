package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct{
	UserName string
	Age int
}

func Index(w http.ResponseWriter ,r *http.Request){
	t,err:=template.ParseFiles("C:\\Users\\zhaojz\\go\\src\\go_dev\\day10\\example5\\main\\index.html")
	if err!=nil{
		fmt.Println(err)
		return
	}
	u:=User{
		UserName:"张三",
		Age:28,
	}
	t.Execute(w, u)
}

func main(){
	http.HandleFunc("/", Index)
	http.ListenAndServe("0.0.0.0:8088", nil)
}
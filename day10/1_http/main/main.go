package main

import (
	"fmt"
	"net/http"
)

func Hello(res http.ResponseWriter,req *http.Request){
	fmt.Println("hello web")
	fmt.Fprintf(res, "hello web")
}

func Login(res http.ResponseWriter,req *http.Request){
	fmt.Println("hello login")
	fmt.Fprintf(res, "hello login")
}

func main(){
	http.HandleFunc("/user/login", Login)
	http.HandleFunc("/", Hello)
	err:=http.ListenAndServe("0.0.0.0:8000",nil)
	if err!=nil{
		fmt.Println("http listen failed")
	}
}
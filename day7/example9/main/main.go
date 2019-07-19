package main

import (
	"fmt"
	"encoding/json"
)

type User struct{
	UserName string  `json:"username"`
	NickName string  `json:"nickname"`
	Age int `json:"age"`
	Gender string `json:"gender"`
	Email string `json:"email"`
	MobileNo string `json:"mobileno"`

}

func testStruct(){
	user := &User{
		UserName:"张三", 
		NickName:"张三丰",
		Age:18,
		Gender:"man",
		Email:"zsf@163.com",
		MobileNo:"13999999999",
	}
	data,err:=json.Marshal(user)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

}

func testMap(){
	m:=make(map[string]interface{})
	m["username"]="张三"
	m["age"]=18
	m["gender"]="man"
	data,err:=json.Marshal(m)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

}



func testSlice(){
	var s [] map[string]interface{}
	m:=make(map[string]interface{})
	m["username"]="张三"
	m["age"]=18
	m["gender"]="man"
	s = append(s, m)
	m=make(map[string]interface{})
	m["username"]="李四"
	m["age"]=19
	m["gender"]="female"
	s = append(s, m)
	data,err:=json.Marshal(s)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

}

func main(){
	//testStruct()
	//testMap()
	testSlice()
}
package main

import (
	"fmt"
	"encoding/json"
)

type User struct{
	UserName string `json:"username"`
	Age int `json:"age"`
	Gender string `json:"gender"`
}

func testStruct(){
	jsonstr:="{\"age\":19,\"gender\":\"female\",\"username\":\"李四\"}"
	var user User
	json.Unmarshal([]byte(jsonstr),&user)
	fmt.Println(user)
}

func testMap(){
	jsonstr:="{\"age\":29,\"gender\":\"female\",\"username\":\"李1四\"}"
	var m map[string]interface{}
	json.Unmarshal([]byte(jsonstr),&m)
	fmt.Println(m)
}

func testSlice(){
	jsonstr:="[{\"age\":18,\"gender\":\"man\",\"username\":\"张三\"},{\"age\":19,\"gender\":\"female\",\"username\":\"李四\"}]"
	var s []map[string]interface{}
	json.Unmarshal([]byte(jsonstr),&s)
	fmt.Println(s)
}

func main(){
	testStruct()
	testMap()
	testSlice()
}
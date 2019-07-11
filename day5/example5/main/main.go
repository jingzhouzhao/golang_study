package main

import (
	"fmt"
	"encoding/json"
)

type Student struct{
	Name string `json:"name___"`
	Age int `json:"age"`
	Score float32 `json:"score"`
}

func main(){
	var stu = Student{
		Name:"zhangsan",
		Age:18,
		Score:88.8,
		
	}
	data,err:=json.Marshal(stu)
	if err!=nil{
		fmt.Print("convert to json failed！")
	}
	fmt.Print(string(data))
}
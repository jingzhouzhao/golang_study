package main

import (
	"io/ioutil"
	"fmt"
	"net/http"
)

func main(){
	res,err:=http.Get("https://www.baidu.com")
	if err!=nil{
		fmt.Println("http get failed:",err)
	}
	body,err:=ioutil.ReadAll(res.Body)
	if err!=nil{
		fmt.Println("http get body failed:",err)
	}
	fmt.Println(string(body))
}
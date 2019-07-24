package main

import (
	"fmt"
)

func main(){
	var str string = "day day 向上"
	for i,v:= range str {
		fmt.Printf("index:[%d],value:[%c],len:[%d]\n", i,v,len([]byte(string(v))))
	}
}
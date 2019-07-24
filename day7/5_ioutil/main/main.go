package main

import (
	"os"
	"fmt"
	"io/ioutil"
)

func main(){
	buf,err:=ioutil.ReadFile("D:/test.log")
	if err!=nil{
		fmt.Fprintf(os.Stderr,"File Error:%s\n", err)
		return
	}
	fmt.Printf("read content:%s\n",string(buf))
	err=ioutil.WriteFile("D:/test.log.bak", buf, 0664)
	if err!=nil{
		//fmt.Fprintf(os.Stderr, "Write File Error:%s\n", err)
		panic(err.Error())
	}
}
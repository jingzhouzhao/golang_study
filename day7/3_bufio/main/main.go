package main

import (
	"fmt"
	"os"
	"bufio"
)

var reader *bufio.Reader

func main(){
	file,err:=os.Open("D:/test.log")
	if err!= nil{
		fmt.Println("open file failed:",err)
		return
	}
	defer file.Close()
	reader = bufio.NewReader(file)
	str,err:=reader.ReadString('\n')
	if err !=nil {
		fmt.Println("read failed!")
		return
	}
	fmt.Printf("read string:%s\n", str)




}
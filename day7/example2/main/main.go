package main

import (
	"fmt"
	"os"
	"bufio"
)

var reader *bufio.Reader

func main(){

	reader = bufio.NewReader(os.Stdin)
	str,err:=reader.ReadString('\n')
	if err !=nil {
		fmt.Println("read failed!")
		return
	}
	fmt.Printf("read string:%s\n", str)




}
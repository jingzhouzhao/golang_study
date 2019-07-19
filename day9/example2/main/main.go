package main

import (
	"time"
	"strings"
	"os"
	"bufio"
	"fmt"
	"net"
)


func main() {
	conn,err:=net.DialTimeout("tcp", "localhost:50000",2*time.Second)
	if err!=nil{
		fmt.Println("Error Dialing:",err)
		return
	}
	defer conn.Close()
	reader:=bufio.NewReader(os.Stdin)
	for{
		input,_:=reader.ReadString('\n')
		input = strings.Trim(input,"\t\n")
		if input=="Q"{
			return
		}
		_,err:=conn.Write([]byte(input))
		if err!=nil{
			fmt.Println("write error:",err)
			return
		}
	}

}



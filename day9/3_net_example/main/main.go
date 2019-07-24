package main

import (
	"io"
	"fmt"
	"net"
)

func main() {
	conn,err:=net.Dial("tcp", "www.baidu.com:80")
	if err!=nil{
		fmt.Println("Error dialing:",err)
		return
	}
	defer conn.Close()
	header:="GET / HTTP/1.1\r\n"
	header+="Host:www.baidu.com\r\n"
	//请求-响应后关闭连接
	//header+="Connection: close\r\n"
	//请求-响应后保持连接
	header+="Connection: keep-alive"
	header+="\r\n\r\n"

	n,err:=io.WriteString(conn, header)
	if err!=nil{
		fmt.Println("write string failed:",err)
		return 
	}
	fmt.Println("sent to bytes:",n)
	buf:=make([]byte,4096)
	for{
		n,err=conn.Read(buf)
		if err!=nil{
			if err==io.EOF{
				fmt.Println("连接关闭")
				break
			}
			fmt.Println("read res error:",err)
			return
		}
		fmt.Println(string(buf[0:n]))

	}
}



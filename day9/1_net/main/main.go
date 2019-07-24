package main

import (
	"io"
	"fmt"
	"net"
)

func main() {
	fmt.Println("server starting")
	addr := "0.0.0.0:50000"
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("listen [%s] failed err:%s", addr, err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed err:%s\n", err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 512)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err==io.EOF{
				conn.Close()
				fmt.Println("链接关闭")
				return
			}
			fmt.Println("read data error:", err)
			return
		}
		fmt.Printf(string(buf[0:n]))
	}
}

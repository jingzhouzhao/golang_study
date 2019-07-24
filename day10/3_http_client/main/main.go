package main

import (
	"time"
	// "net"
	"fmt"
	"net/http"
)

func main(){
	var urls = [] string{
		"https://www.baidu.com",
		"https://www.google.com",
		"https://www.taobao.com",
	}
	c:=http.Client{
		// Transport:&http.Transport{
		// 	Dial:func(network, addr string) (net.Conn, error){
		// 		return net.DialTimeout(network, addr, 2*time.Second)
		// 	},
		// },
		Timeout:2*time.Second,
	}
	for _,v:=range urls{
		r,err:=c.Head(v)
		if err!=nil{
			fmt.Println("http head failed:",err)
			continue
		}
		fmt.Println(r.Status)
	}

}
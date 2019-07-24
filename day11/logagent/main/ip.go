package main

import (
	"fmt"
	"net"
)

var (
	localIPs []string
)

func init(){
	addrs,err:=net.InterfaceAddrs()
	if err!=nil || len(addrs) == 0{
		panic(fmt.Sprintf("Failed to get IP:%s", err))
	}
	for _,addr:=range addrs{
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			localIPs = append(localIPs, ipnet.IP.String())
		}
	}
	fmt.Println("localIPs:",localIPs)
}
package main

import (
	"fmt"
	"github.com/astaxie/beego/config"
)
func main(){
	conf,err:=config.NewConfig("ini", "./logagent.conf")
	if err !=nil{
		fmt.Println("new config error:",err)
		return
	}
	addr:=conf.String("server::addr")
	fmt.Println("server.addr:",addr)
	port,err:=conf.Int("server::port")
	if err!=nil{
		fmt.Println("read server port failed,err:",err)
		return
	}
	fmt.Println("server.port:",port)
	logLevel:=conf.String("log::level")
	fmt.Println("log.level:",logLevel)

}
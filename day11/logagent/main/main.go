package main

import (
	"go_dev/day11/logagent/kafka"
	"go_dev/day11/logagent/tailf"
	"github.com/astaxie/beego/logs"
	"fmt"
)


func main(){
	err:=loadConfig("ini","agentconfig.ini")
	if err!=nil {
		fmt.Println("load config failed:",err)
		panic(err)
	}
	//fmt.Printf("%#v\n",appConfig)
	err = initLogger()
	if err!=nil {
		fmt.Println("init logger failed:",err)
		panic(err)
	}

	err = tailf.InitTail(appConfig.collectConfigs)
	if err!=nil {
		logs.Error("init tail failed:",err)
		panic(err)
	}
	err = kafka.InitKafka(appConfig.kafkaAddr)
	if err!=nil {
		logs.Error("init kafka failed:",err)
		panic(err)
	}
	logs.Debug("initialize Success.")
	logs.Debug("load config Success,config:%v", appConfig)

	err=runServer()
	if err!=nil {
		logs.Error("run server failed:",err)
		panic(err)
	}
	logs.Info("program exited")
	
}
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
	fmt.Println("load config successful")
	//fmt.Printf("%#v\n",appConfig)
	err = initLogger()
	if err!=nil {
		fmt.Println("init logger failed:",err)
		panic(err)
	}

	cli,err := initEtcd(appConfig.etcdAddr, appConfig.etcdTimeout)
	if err!=nil {
		logs.Error("init etcd failed:",err)
		panic(err)
	}

	collectConfs,err:=loadCollectConfig(cli, appConfig.etcdKey, appConfig.etcdTimeout)
	if err!=nil {
		logs.Error("load collect configs failed:",err)
		panic(err)
	}
	err = tailf.InitTail(collectConfs)
	if err!=nil {
		logs.Error("init tail failed:",err)
		panic(err)
	}
	err = kafka.InitKafka(appConfig.kafkaAddr)
	if err!=nil {
		logs.Error("init kafka failed:",err)
		panic(err)
	}


	logs.Debug("Initialization all successful")
	

	err=runServer()
	if err!=nil {
		logs.Error("run server failed:",err)
		panic(err)
	}
	logs.Info("program exited")
	
}
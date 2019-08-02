package main

import (
	"fmt"
)

func main(){

	err:=loadConfig("ini", "transferconfig.conf")
	if err != nil {
		fmt.Println("load config failed:",err)
		panic(err)
	}
	err=initLogger()
	if err != nil {
		fmt.Println("init logger failed:",err)
		panic(err)
	}
	err =initEtcd(appConfig.etcdAddr, appConfig.etcdTimeout)
	if err != nil {
		fmt.Println("init etcd failed:",err)
		panic(err)
	}
	err=initKafka(appConfig.kafkaAddr)
	if err != nil {
		fmt.Println("init kafka failed:",err)
		panic(err)
	}

	err=initElasticSearch(appConfig.esAddr)

	runServer()

}
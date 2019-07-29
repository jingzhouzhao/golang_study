package main

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/config"
)

type AppConfig struct{
	logLevel string
	logPath string
	kafkaAddr string
	etcdAddr string 
	etcdTimeout int
	etcdKey string
	//collectConfigs []tailf.CollectConfig
}
var (
	appConfig = defaultAppConfig()
)

func defaultAppConfig() *AppConfig{
	return &AppConfig{
		logLevel:"debug",
		logPath:"/logs/agent.log",
	}
}

//Deprecated  use etcd
// func loadCollectConfig(conf config.Configer)(err error){
// 	var collectConfig tailf.CollectConfig
// 	collectConfig.LogPath = conf.String("collect::logPath")
// 	if len(collectConfig.LogPath)==0{
// 		err = errors.New("invalid collect::logPath")
// 		return
// 	}
// 	collectConfig.Topic = conf.String("collect::topic")
// 	if len(collectConfig.Topic)==0{
// 		err = errors.New("invalid collect::topic")
// 		return
// 	}
// 	collectConfig.ChanSize,err = conf.Int("collect::chanSize")
// 	if err!=nil{
// 		err = errors.New("invalid collect::chanSize")
// 		return
// 	}
// 	appConfig.collectConfigs = append(appConfig.collectConfigs, collectConfig)
// 	return 

// }

func loadConfig(confType,filename string) (err error){
	conf,err:=config.NewConfig(confType, filename)
	if err !=nil{
		fmt.Println("new config failed:",err)
		return
	}
	logLevel:=conf.String("agent::logLevel")
	if len(logLevel) > 0{
		appConfig.logLevel = logLevel
	}
	logPath:=conf.String("agent::logPath")
	if len(logPath) > 0{
		appConfig.logPath = logPath
	}
	kafkaAddr:=conf.String("kafka::addr")
	if len(kafkaAddr) == 0{
		err = errors.New("kafka::addr must not be empty")
		return
	}
	appConfig.kafkaAddr = kafkaAddr

	etcdAddr:=conf.String("etcd::addr")
	if len(etcdAddr) == 0{
		err = errors.New("etcd::addr must not be empty")
		return
	}
	appConfig.etcdAddr = etcdAddr

	etcdTimeout,err := conf.Int("etcd::timeout")
	if err!=nil{
		err = errors.New("invalid etcd::addr")
		return
	}
	appConfig.etcdTimeout = etcdTimeout

	etcdKey:=conf.String("etcd::key")
	if len(etcdKey) == 0{
		err = errors.New("etcd::key must not be empty")
		return
	}
	appConfig.etcdKey = etcdKey
	
	//err=loadCollectConfig(conf)
	
	return
}
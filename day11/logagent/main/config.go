package main

import (
	"go_dev/day11/logagent/tailf"
	"errors"
	"fmt"
	"github.com/astaxie/beego/config"
)

type AppConfig struct{
	logLevel string
	logPath string
	kafkaAddr string
	collectConfigs []tailf.CollectConfig
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

func loadCollectConfig(conf config.Configer)(err error){
	var collectConfig tailf.CollectConfig
	collectConfig.LogPath = conf.String("collect::logPath")
	if len(collectConfig.LogPath)==0{
		err = errors.New("invalid collect::logPath")
		return
	}
	collectConfig.Topic = conf.String("collect::topic")
	if len(collectConfig.Topic)==0{
		err = errors.New("invalid collect::topic")
		return
	}
	collectConfig.ChanSize,err = conf.Int("collect::chanSize")
	if err!=nil{
		err = errors.New("invalid collect::chanSize")
		return
	}
	appConfig.collectConfigs = append(appConfig.collectConfigs, collectConfig)
	return 

}

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

	err=loadCollectConfig(conf)
	
	return
}
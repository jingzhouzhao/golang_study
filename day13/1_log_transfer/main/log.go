package main

import (
	"fmt"
	"encoding/json"
	"github.com/astaxie/beego/logs"
)

func convertLogLevel(level string) int{
	switch level {
	case "trace":
		return logs.LevelTrace
	case "debug":
		return logs.LevelDebug
	case "info":
		return logs.LevelInfo
	case "warn":
		return logs.LevelWarn
	case "error":
		return logs.LevelError
	}
	return logs.LevelDebug
}

func initLogger()(err error){
	m:=make(map[string]interface{},3)
	m["filename"] = appConfig.logPath
	m["level"] = convertLogLevel(appConfig.logLevel)
	config,err := json.Marshal(m)
	if err!=nil{
		fmt.Println("init Logger failed,marshal config error:",err)
		return
	}
	logs.SetLogger(logs.AdapterFile, string(config))
	//打印函数及行数
	logs.EnableFuncCallDepth(true)
	//2打印的是log.go
	logs.SetLogFuncCallDepth(3)
	return 
}
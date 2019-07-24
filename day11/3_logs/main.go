package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func main() {

	m := make(map[string]interface{}, 3)
	m["filename"] = "./test1.log"
	m["level"] = logs.LevelTrace
	//m["perm"] = "0666" //默认为0660
	config, err := json.Marshal(m)
	if err != nil {
		fmt.Println("marshal failed ,err:", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(config))
	//打印函数及行数
	logs.EnableFuncCallDepth(true)
	//2打印的是log.go
	logs.SetLogFuncCallDepth(3)

	logs.Debug("this is a %s test ", "debug")
	logs.Info("this is a %s test ", "info")
	logs.Warn("this is a %s test ", "warn")
	logs.Error("this is a %s test ", "error")
}

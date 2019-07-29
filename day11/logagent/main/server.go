package main

import (
	"go_dev/day11/logagent/kafka"
	"time"
	"github.com/astaxie/beego/logs"
	"go_dev/day11/logagent/tailf"
)

func runServer()(err error){
	for{
		msg:=tailf.GetMsg()
		if msg==nil{
			continue
		}
		err=sendToKafka(msg)
		if err!=nil{
			logs.Error("send msg to kafka failed:",err)
			time.Sleep(time.Second)
			continue
		}

	}
}

func sendToKafka(msg *tailf.TextMsg)(err error){
	//logs.Debug("msg:%s,topic:%s", msg.Msg,msg.Topic)
	//fmt.Printf("msg:%s,topic:%s\n", msg.Msg,msg.Topic)
	err = kafka.SendToKafka(msg.Msg, msg.Msg, msg.Topic)
	return
}
package kafka

import (
	"github.com/astaxie/beego/logs"
	"github.com/Shopify/sarama"
)

var (
	producer sarama.SyncProducer
)

func InitKafka(addr string)(err error){
	config:=sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	config.Producer.Return.Successes = true
	pd,err:=sarama.NewSyncProducer([]string{addr}, config)
	if err!=nil{
		logs.Error("init kafka[addr:%s] producer failed:%s",addr, err)
		return 
	}
	producer = pd
	logs.Debug("initialize kafka Success")
	return
}

func SendToKafka(key,msg,topic string)(err error){
	m:=&sarama.ProducerMessage{}
	m.Key = sarama.StringEncoder(key)
	m.Value = sarama.StringEncoder(msg)
	m.Topic = topic
	pid,offset,err:=producer.SendMessage(m)
	if err!=nil{
		logs.Error("send message to kafka failed,data:%s,topic:%s", m.Value,m.Topic)
		return
	}
	logs.Debug("send message to kafka success,pid:%v,offset:%v,topic:%v\n",pid,offset,topic)
	return 
}
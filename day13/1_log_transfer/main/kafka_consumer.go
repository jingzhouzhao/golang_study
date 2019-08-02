package main

import (
	"sync"
	"github.com/astaxie/beego/logs"
	"strings"
	"github.com/Shopify/sarama"
)

type KafkaClient struct{
	cli sarama.Consumer
}

var (
	kafkaClient KafkaClient
)

func initKafka(addr string)(err error){
	consumer,err:=sarama.NewConsumer(strings.Split(addr, ","), nil)
	if err!=nil{
		logs.Error("Failed connect to kafka:", err)
		return
	}
	kafkaClient.cli = consumer
	logs.Debug("init kafka success")
	return
}

var (
	wg sync.WaitGroup
)

func consumerTopic() (err error){
	topics,err:=kafkaClient.cli.Topics()
	logs.Debug("topics:", topics)
	if err!=nil{
		logs.Error("get topics failed")
		return
	}
	for _,topic:=range topics{
		pids,errP:=kafkaClient.cli.Partitions(topic)
		if err!=nil{
			logs.Error("get Partitions failed")
			err = errP
			return
		}
		logs.Debug("pids:", pids)
		for _,pid:=range pids{
			pc,err:=kafkaClient.cli.ConsumePartition(topic, pid, sarama.OffsetNewest)
			if err!=nil{
				continue
			}
			defer pc.AsyncClose()
			wg.Add(1)
			go func (partitionConsumer sarama.PartitionConsumer){
				for msg:=range 	pc.Messages(){
					message := IndexMessage{
						Message:string(msg.Value),
						Topic:topic,
					}
					sendToElasticSearch("logs", message)
				}
				wg.Done()
			}(pc)
			
		}
	
	}
	wg.Wait()
	return
}
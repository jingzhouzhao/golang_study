package main

import (
	"time"
	"fmt"
	"strings"
	"github.com/Shopify/sarama"
)


func process(){
	consumer,err:=sarama.NewConsumer(strings.Split("127.0.0.1:9092", ","), nil)
	if err!=nil{
		fmt.Println("Failed connect to kafka:",err)
		return
	}
	pids,err:=consumer.Partitions("test")
	if err!=nil{
		fmt.Println("get pids failed:",err)
		return
	}
	for _,pid:=range pids{
		pc,err:=consumer.ConsumePartition("test", pid, sarama.OffsetNewest)
		if err!=nil{
			fmt.Println("Failed to start consumer partition:",err)
			return
		}
		defer pc.AsyncClose()
		go func (partitionConsumer sarama.PartitionConsumer){
			for msg:=range 	pc.Messages(){
				fmt.Printf("partition:%d,key:%s,value:%s,offset:%d\n",msg.Partition,msg.Key,msg.Value,msg.Offset)
			}
		}(pc)
	}
	time.Sleep(time.Hour)
	consumer.Close()
}

func main(){
	process()
}
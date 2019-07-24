package main

import (
	"time"
	"fmt"
	"github.com/Shopify/sarama"
)

func main(){

	config:=sarama.NewConfig()
	config.Producer.RequiredAcks=sarama.WaitForAll
	//sarama.NewHashPartitioner(topic)
		config.Producer.Partitioner = sarama.NewHashPartitioner
	config.Producer.Return.Successes = true
	client,err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err !=nil {
		fmt.Println("producer closed,err:",err)
	}
	defer client.Close()
	i:=0
	for{
		msg:=&sarama.ProducerMessage{}
		msg.Key = sarama.StringEncoder(fmt.Sprintf("%d",i))
		msg.Topic = "test"
		msg.Value = sarama.StringEncoder(fmt.Sprintf("%d",i))
		pid,offset,err:=client.SendMessage(msg)
		if err !=nil{
			fmt.Println("send message failed:",err)
			return
		}
		fmt.Printf("pid:%d,offset:%d\n", pid,offset)
		i++
		time.Sleep(10*time.Millisecond)
	}
}
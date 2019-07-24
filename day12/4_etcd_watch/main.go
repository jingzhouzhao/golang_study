package main

import (
	"context"
	"fmt"
	"time"
	"go.etcd.io/etcd/client"
)

func process(){
	cfg:=client.Config{
		Endpoints:[]string{"http://127.0.0.1:2379"},
		Transport: client.DefaultTransport,
		HeaderTimeoutPerRequest: 3 * time.Second,
	}
	cli,err:=client.New(cfg)
	if err!=nil{
		fmt.Println("Failed to connect to etcd:",err)
		return
	}
	kapi:=client.NewKeysAPI(cli)
	w:=kapi.Watcher("/logagent/conf",&client.WatcherOptions{
		Recursive:true,
	})
	for{
		res,err:=w.Next(context.Background())
		if err!=nil{
			fmt.Println("watcher failed:",err)
			break
		}
		fmt.Println("watcher action :",res.Action)
		fmt.Printf("key:%s,value:%s\n",res.Node.Key,res.Node.Value)

	}
	
}

func main(){
	process()
}
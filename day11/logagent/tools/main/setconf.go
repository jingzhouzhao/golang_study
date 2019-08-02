package main

import (
	"context"
	"encoding/json"
	"go_dev/day11/logagent/tailf"
	"fmt"
	"time"
	"go.etcd.io/etcd/client"
)

const (
	etcdKey = "/logagent/conf/172.23.154.1/"
)

func process(){
	cli,err:=client.New(client.Config{
		Endpoints:[]string{"http://127.0.0.1:2379"},
		Transport: client.DefaultTransport,
		HeaderTimeoutPerRequest: 3 * time.Second,
	})
	if err!=nil{
		panic(err)
	}
	kapi:=client.NewKeysAPI(cli)

	var collectConfs []tailf.CollectConfig
	collectConfs = append(collectConfs, tailf.CollectConfig{
		LogPath:"D:/nginx/logs/access.log",
		Topic:"nginx_log",
		ChanSize:100,
	})
	collectConfs = append(collectConfs, tailf.CollectConfig{
		LogPath:"D:/nginx/logs/error.log",
		Topic:"nginx_error_log",
		ChanSize:100,
	})

	collectConfs = append(collectConfs, tailf.CollectConfig{
		LogPath:"D:/nginx/logs/error1.log",
		Topic:"nginx_error_log",
		ChanSize:100,
	})


	val,err:=json.Marshal(collectConfs)
	if err!=nil{
		panic(err)
	}
	// res,err:=kapi.Delete(context.Background(), etcdKey, &client.DeleteOptions{
	// 	Recursive:true,
	// 	Dir:true,
	// })
	// if err!=nil{
	// 	panic(err)
	// }
	// fmt.Println("del_key:",res.Node.Key)
	// return
	res,err:=kapi.Set(context.Background(), etcdKey, string(val), nil)
	if err!=nil{
		panic(err)
	}
	fmt.Printf("key:%s,value:%s\n", res.Node.Key,res.Node.Value)
}

func main(){
	process()
}
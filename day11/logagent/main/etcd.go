package main

import (
	"fmt"
	"context"
	"github.com/astaxie/beego/logs"
	"time"
	"strings"
	"go.etcd.io/etcd/client"
)

func initEtcd(addr string,timeout int,key string)(err error){
	cli,err:=client.New(client.Config{
		Endpoints:strings.Split(addr, ","),
		HeaderTimeoutPerRequest:time.Duration(timeout)*time.Second,
	})
	if err!=nil{
		logs.Error("Failed to connect to etcd:%s",err)
		return 
	}
	logs.Debug("initialize etcd Success")
	kapi:=client.NewKeysAPI(cli)
	for _,ip:=range localIPs{
		if !strings.HasSuffix(key, "/"){
			key = key + "/"
		}
		etcdKey:= fmt.Sprintf("%s%s", key,ip)
		ctx,cancle:=context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
		res,err:=kapi.Get(ctx, etcdKey, nil)
		cancle()
		if err != nil{
			continue
		}
		for _,kv:=range res.Node.Nodes{
			fmt.Printf("key:%s,value:%s\n",kv.Key,kv.Value)
		}
	}	
	
	return
}
package main

import (
	"context"
	"time"
	"fmt"
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
	fmt.Print("Successful connection to etcd")
	kapi:=client.NewKeysAPI(cli)
	res,err:=kapi.Set(context.Background(), "user_id", "123123", nil)
	if err!=nil{
		fmt.Println("set key to etcd failed:",err)
		return
	}
	fmt.Println("set key to etcd successed:",res)
	res,err=kapi.Get(context.Background(), "user_id", nil)
	if err!=nil{
		fmt.Println("get key from etcd failed:",err)
		return
	}
	fmt.Printf("get key from etcd successed:%#v\n",res)
	fmt.Printf("key:%s,value:%s\n",res.Node.Key,res.Node.Value)
	//fmt.Printf("prevkey:%s,prevvalue:%s", res.PrevNode.Key,res.PrevNode.Value)
	fmt.Println("####################3")
	ctx,cancle:=context.WithTimeout(context.Background(), 2*time.Second)
	_,err=kapi.Set(ctx, "/logagent/conf/logPath", "logagent.log",nil)
	cancle()
	if err!=nil{
		fmt.Println("set node failed:",err)
	}

	ctx,cancle=context.WithTimeout(context.Background(), 2*time.Second)
	_,err=kapi.Set(ctx, "/logagent/conf/logPath2", "logagent.log2",nil)
	cancle()
	if err!=nil{
		fmt.Println("set node failed:",err)
	}

	ctx,cancle =context.WithTimeout(context.Background(), time.Nanosecond)
	res,err=kapi.Get(ctx, "/logagent/conf",nil)
	cancle()
	if err!=nil{
		fmt.Println("set node failed:",err)
	}
	for _,kv:=range res.Node.Nodes{
		fmt.Printf("key:%s,value:%s\n",kv.Key,kv.Value)
	}

}

func main(){
	process()
}
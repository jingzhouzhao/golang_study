package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"go.etcd.io/etcd/client"
	"go_dev/day11/logagent/tailf"
	"strings"
	"time"
)

func initEtcd(addr string, timeout int) (cli client.Client, err error) {
	cli, err = client.New(client.Config{
		Endpoints:               strings.Split(addr, ","),
		HeaderTimeoutPerRequest: time.Duration(timeout) * time.Second,
	})
	if err != nil {
		logs.Error("Connection etcd failed:%s", err)
		return
	}
	logs.Debug("initialize etcd Success")
	return
}

func loadCollectConfig(cli client.Client, key string, timeout int) (collectConfs []tailf.CollectConfig, err error) {
	kapi := client.NewKeysAPI(cli)
	for _, ip := range localIPs {
		if !strings.HasSuffix(key, "/") {
			key = key + "/"
		}
		etcdKey := fmt.Sprintf("%s%s", key, ip)
		go watchKey(cli, etcdKey)
		ctx, cancle := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
		res, err := kapi.Get(ctx, etcdKey, nil)
		cancle()
		if err != nil {
			continue
		}
		if res.Node.Key == etcdKey {
			err = json.Unmarshal([]byte(res.Node.Value), &collectConfs)
			break
		}
	}
	return
}

func watchKey(cli client.Client, key string) {
	kapi := client.NewKeysAPI(cli)
	logs.Debug("watcher key:", key)
	for {
		w := kapi.Watcher(key, nil)
		res,err:=w.Next(context.Background())
		if err!=nil{
			logs.Error("watcher key error:", err)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		//include get, set, delete, update, create, compareAndSwap,
		switch res.Action {
		case "delete":
			logs.Debug("delete collect configs")
			var collectConfs []tailf.CollectConfig
			tailf.UpdateCollectConfigs(collectConfs)
		case "update":
			fallthrough
		case "set":
			logs.Debug("update collect configs")
			var collectConfs []tailf.CollectConfig
			err = json.Unmarshal([]byte(res.Node.Value), &collectConfs)
			if err !=nil{
				logs.Error("Unmarshal new collectConfigs failed:", err)
			}else{
				tailf.UpdateCollectConfigs(collectConfs)
			}
		default:
			logs.Info("ignore watcher action:", res.Action)
		}
	}
}

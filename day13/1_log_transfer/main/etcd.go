package main

import (
	"strings"
	"github.com/astaxie/beego/logs"
	"go.etcd.io/etcd/client"
	"time"
)

type EtcdClient struct{
	cli client.Client
}

var (
	etcdClient EtcdClient
)

func initEtcd(addr string, timeout int) (err error) {
	cli, err := client.New(client.Config{
		Endpoints:               strings.Split(addr, ","),
		HeaderTimeoutPerRequest: time.Duration(timeout) * time.Second,
	})
	if err != nil {
		logs.Error("Connection etcd failed:%s", err)
		return
	}
	etcdClient.cli = cli
	logs.Debug("initialize etcd Success")
	return
}
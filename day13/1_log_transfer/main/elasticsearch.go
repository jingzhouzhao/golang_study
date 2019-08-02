package main

import (
	"context"
	"github.com/astaxie/beego/logs"
	"gopkg.in/olivere/elastic.v5"
)

type ElasticSearchClient struct {
	cli *elastic.Client
}

var (
	es ElasticSearchClient
)

func initElasticSearch(esAddr string) (err error) {
	client, err := elastic.NewClient(elastic.SetURL(esAddr))
	if err != nil {
		logs.Error("new elasticsearch client failed:", err)
		return
	}
	es.cli = client
	logs.Debug("init elasticsearch success")
	return
}

type Tweet struct {
	User     string
	Message  string
	Retweets int
}

type IndexMessage struct{
	Message string `json:"message"`
	Topic string `json:"topic"`
}

func sendToElasticSearch(index string, msg IndexMessage) {
	put,err:=es.cli.Index().
	Index(index).
	Type(index).
	//Id("1").
	BodyJson(msg).
	Do(context.Background())
	if err!=nil{
		logs.Error("Index failed:", err)
		return
	}
	logs.Debug("Indexed %s to index %s, type %s\n", put.Id, put.Index, put.Type)
}

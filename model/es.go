package model

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v7"
	"log"
)

var EsClient *elastic.Client

const (
	esHost  string = "127.0.0.1"
	esPort  string = "9200"
	esIndex        = "mylog" //使用的索引名称
)

func init() {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://"+esHost+":"+esPort))
	if err != nil {
		log.Panic(err)
	}
	EsClient = client
}

func EsHookLog() *elogrus.ElasticHook {
	fmt.Println(EsClient)
	hook, err := elogrus.NewElasticHook(EsClient, esHost, logrus.DebugLevel, esIndex)
	fmt.Println("hook", hook)
	if err != nil {
		log.Panic(err)
	}
	return hook
}

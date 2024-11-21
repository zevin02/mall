package mq

import (
	"github.com/streadway/amqp"
	"strings"
)

// 创建一个rabbitmq的单例模式
var RabbitMQ *amqp.Connection

const (
	RabbitMq         string = "amqp"
	RabbitMQUser     string = "guest"
	RabbitMQPassWord string = "guest"
	RabbitMQHost     string = "localhost"
	RabbitMQPort     string = "5672"
)

// 初始化rabbitmq的连接
func InitRabbitMQ() {
	pathRabbitMQ := strings.Join([]string{RabbitMq, "://", RabbitMQUser, ":", RabbitMQPassWord, "@", RabbitMQHost, ":", RabbitMQPort, "/"}, "")
	conn, err := amqp.Dial(pathRabbitMQ) //连接上mq
	if err != nil {
		panic(err)
	}
	RabbitMQ = conn

}

package main

import (
	"fmt"
	"mall/conf"
	"mall/routes"
	"mall/service"
)

func main() {
	conf.Init()
	fmt.Println("finish")
	// 启动 MQ 消费者
	go func() {
		if err := service.MQ2MySQL(); err != nil {
			fmt.Printf("Failed to start MQ consumer: %v\n", err)
		}
	}()
	r := routes.NewRouter() //前端的api路由
	r.Run(conf.HttpPort)

}

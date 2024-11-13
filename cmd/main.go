package main

import (
	"fmt"
	"mall/conf"
	"mall/routes"
)

func main() {
	conf.Init()
	fmt.Println("finish")
	r := routes.NewRouter() //前端的api路由
	r.Run(conf.HttpPort)

}

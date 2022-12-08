package main

import (
	"hi-trip/initialize"
)

func main() {
	//初始化数据库
	initialize.InitializeDB()

	//初始化日志
	initialize.InitLogger()

	//初始化路由
	router := initialize.Routers()
	router.Run(":8088")
}

package main

import (
	"my-backend/initialize"
)

func main() {

	// 初始化
	initialize.Init()

	// 初始化路由
	r := initialize.InitRoutes()

	// 启动服务
	r.Run("127.0.0.1:9000")
}

package main

import (
	"go-sec-code/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 初始化路由
	routers.InitRoutes(r)

	// 启动服务器
	r.Run()
}

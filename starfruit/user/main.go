package main

import (
	"github.com/gin-gonic/gin"
	srv "starfruit.top/common"
	_ "starfruit.top/user/api"
	"starfruit.top/user/router"
)

func main() {
	r := gin.Default()

	// 初始化路由
	router.InitRouter(r)

	// 服务运行
	srv.Run(r, ":8080", "user")

}

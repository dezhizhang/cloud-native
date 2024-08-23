package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
)

func handleExit(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"name": "tom",
	})
}

func main() {
	r := gin.Default()
	r.GET("/", handleExit)

	go func() {
		_ = r.Run(":8080")
	}()

	// 接收信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("服务关闭中....")

}

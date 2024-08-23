package examples

import (
	"gin/middleware"
	"github.com/gin-gonic/gin"
)

func handleUserInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"name": "tom",
	})
}

func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1").Use(middleware.Logger())
	{
		v1.GET("/info", handleUserInfo)
	}

	_ = r.Run(":8080")
}

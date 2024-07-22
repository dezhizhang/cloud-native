package examples

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
	log.Fatalf("服务运行在%d端口上", 8080)
}

package examples

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandlerPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func main() {
	r := gin.Default()
	r.GET("/ping", HandlerPong)
	_ = r.Run(":8080")
}

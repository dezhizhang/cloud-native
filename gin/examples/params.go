package examples

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleGetInfo(c *gin.Context) {
	id := c.Param("id")
	action := c.Param("action")
	c.JSON(http.StatusOK, gin.H{
		"id":     id,
		"action": action,
	})
}

func main() {
	r := gin.Default()
	r.GET("/info/:id/:action", handleGetInfo)
	_ = r.Run(":8080")
}

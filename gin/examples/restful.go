package examples

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandlerGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get",
	})
}

func HandlerPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "post",
	})
}

func HandlerPut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "put",
	})
}

func HandlerDelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "delete",
	})
}

func main() {
	r := gin.Default()
	r.GET("/get", HandlerGet)
	r.POST("/post", HandlerPost)
	r.PUT("/put", HandlerPut)
	r.DELETE("/delete", HandlerDelete)
	_ = r.Run(":8080")
}

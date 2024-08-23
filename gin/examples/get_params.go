package examples

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func handlerGetParams(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})
}

func main() {
	r := gin.Default()
	r.GET("/get-params", handlerGetParams)
	_ = r.Run(":8080")
}

package examples

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	Id   int    `uri:"id" binding:"required" json:"id"`
	Name string `uri:"name" binding:"required" json:"name"`
}

func handleParams(c *gin.Context) {
	var person Person
	if err := c.ShouldBindUri(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, person)
}

func main() {
	r := gin.Default()
	r.GET("/ping", handleParams)
	_ = r.Run(":8080")
}

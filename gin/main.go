package main

import (
	"github.com/gin-gonic/gin"
)

func handlePostParams(c *gin.Context) {
	name := c.PostForm("name")
	age := c.PostForm("age")
	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func main() {
	r := gin.Default()
	r.POST("/post-params", handlePostParams)
	_ = r.Run(":8080")
}

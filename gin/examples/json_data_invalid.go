package examples

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserLogin struct {
	Name     string `json:"name" binding:"required,min=2,max=20"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}

func handleJsonLogin(c *gin.Context) {
	var user UserLogin
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, &user)
}

func main() {
	r := gin.Default()
	r.POST("/login", handleJsonLogin)
	_ = r.Run(":8080")
}

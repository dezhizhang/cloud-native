package examples

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginForm struct {
	Name     string `form:"name" binding:"required,min=2,max=10"`
	Password string `form:"password" binding:"required,min=8,max=20"`
}

func handleLogin(c *gin.Context) {
	var loginForm LoginForm
	if err := c.ShouldBind(&loginForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, &loginForm)
}

func main() {
	r := gin.Default()
	r.POST("/login", handleLogin)
	_ = r.Run(":8080")
}

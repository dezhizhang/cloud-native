package examples

import "github.com/gin-gonic/gin"

func handleLogin(c *gin.Context) {

}

func handleSubmit(c *gin.Context) {

}

func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.POST("/login", handleLogin)
		v1.POST("/submit", handleSubmit)
	}

	v2 := r.Group("/api/v2")
	{
		v2.POST("/login", handleLogin)
		v2.POST("/submit", handleSubmit)
	}
}

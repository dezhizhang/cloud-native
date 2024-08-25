package main

import (
	"github.com/gin-gonic/gin"
	"starfruit.top/user_web/router"
)

func main() {
	r := gin.Default()
	router.RegisterUserRouter(r)
	err := r.Run(":8088")
	if err != nil {
		panic(err)
	}
}

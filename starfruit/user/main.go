package main

import (
	"github.com/gin-gonic/gin"
	srv "starfruit.top/common"
)

func main() {
	r := gin.Default()

	srv.Run(r, ":8080", "user")

}

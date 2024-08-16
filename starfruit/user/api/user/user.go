package user

import "github.com/gin-gonic/gin"

type HandlerUser struct {
}

func (*HandlerUser) GetCaptcha(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"msg": "hello world"})
}

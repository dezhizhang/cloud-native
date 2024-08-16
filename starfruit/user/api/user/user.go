package user

import (
	"github.com/gin-gonic/gin"
	"starfruit.top/common"
)

type HandlerUser struct {
}

func (*HandlerUser) GetCaptcha(ctx *gin.Context) {

	rsp := common.Result{}

	ctx.JSON(200, rsp.Success("123456"))
}

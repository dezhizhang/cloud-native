package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"starfruit.top/common"
	"starfruit.top/user/pkg/model"
	"time"
)

type HandlerUser struct {
}

// GetCaptcha 获取手机验证码
func (*HandlerUser) GetCaptcha(ctx *gin.Context) {

	rsp := common.Result{}
	//1. 获取参数
	mobile := ctx.PostForm("mobile")
	//2. 校验参数
	if common.VerifyMobile(mobile) {
		ctx.JSON(http.StatusOK, rsp.Fail(model.NoLegalMobil, "手机号不合法"))
		return
	}
	//3. 生成验证码
	code := "123456"

	//4. 调用短信平台
	go func() {
		time.Sleep(2 * time.Second)
		log.Println("短信平台调用成功，发送短信")
		//5.存储验证码过期时间为15分钟
		log.Printf("验证码存入redis成功:REGISTER_%s:%s", mobile, code)
	}()

	ctx.JSON(200, rsp.Success(code))
}

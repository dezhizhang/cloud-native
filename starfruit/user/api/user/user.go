package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"starfruit.top/common"
	"starfruit.top/user/pkg/dao"
	"starfruit.top/user/pkg/model"
	"starfruit.top/user/pkg/repo"
	"time"
)

type HandlerUser struct {
	cache repo.Cache
}

func New() *HandlerUser {
	return &HandlerUser{cache: dao.RC}
}

// GetCaptcha 获取手机验证码
func (h *HandlerUser) GetCaptcha(ctx *gin.Context) {

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
		c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		err := h.cache.Put(c, "REGISTER_"+mobile, code, 15*time.Minute)
		if err != nil {
			log.Printf("注册验证码存入redis出差%v\n", err)
		}
	}()

	ctx.JSON(200, rsp.Success(code))
}

package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"starfruit.top/user/router"
)

type RouterUser struct {
}

func init() {
	log.Println("init user router")
	router.Register(&RouterUser{})
}

func (*RouterUser) Route(r *gin.Engine) {
	h := &HandlerUser{}
	r.POST("/api/v1/login/captcha", h.GetCaptcha)
}

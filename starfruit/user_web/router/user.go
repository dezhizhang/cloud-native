package router

import (
	"github.com/gin-gonic/gin"
	"starfruit.top/user_web/api"
)

func RegisterUserRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1/user")
	{
		// 用户登录
		v1.POST("/login", api.HandleUserLogin)
		// 用户注册
		v1.POST("/register", api.HandleUserRegister)
		// 获取用户列表
		v1.GET("/list", api.HandleGetUserList)
		// 更新用户信息
		v1.POST("/update", api.HandleUpdateUser)
	}
}

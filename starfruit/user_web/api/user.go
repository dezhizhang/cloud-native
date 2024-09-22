package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net/http"
	"starfruit.top/user_web/proto"
	"starfruit.top/user_web/utils"
)

// HandleUserLogin 用户登录
func HandleUserLogin(c *gin.Context) {

}

// HandleUserRegister 用户注册
func HandleUserRegister(c *gin.Context) {

}

// HandleGetUserList 获取用户列表
func HandleGetUserList(c *gin.Context) {
	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	utils.HandlerGrpcErrorToHttp(c, err)

	defer conn.Close()

	client := proto.NewUserClient(conn)
	list, err := client.GetUserList(context.Background(), &proto.RequestUser{
		Page: 1,
		Size: 20,
	})
	utils.HandlerGrpcErrorToHttp(c, err)

	c.JSON(200, &list)
}

// HandleUpdateUser 更新用户信息
func HandleUpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

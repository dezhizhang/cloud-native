package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"starfruit.top/user_web/proto"
)

// 将grpc转换成http
func handlerGrpcErrorToHttp(c *gin.Context, err error) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"message": e.Message(),
				})
				return
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "服务器出错了",
				})
				return
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "其它错训",
				})

			}

		}
	}
}

// HandleUserLogin 用户登录
func HandleUserLogin(c *gin.Context) {

}

// HandleUserRegister 用户注册
func HandleUserRegister(c *gin.Context) {

}

// HandleGetUserList 获取用户列表
func HandleGetUserList(c *gin.Context) {
	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	handlerGrpcErrorToHttp(c, err)

	defer conn.Close()

	client := proto.NewUserClient(conn)
	list, err := client.GetUserList(context.Background(), &proto.RequestUser{
		Page: 1,
		Size: 20,
	})
	handlerGrpcErrorToHttp(c, err)

	c.JSON(200, &list)
}

// HandleUpdateUser 更新用户信息
func HandleUpdateUser(c *gin.Context) {

}

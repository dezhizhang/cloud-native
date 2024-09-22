package utils

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

// HandlerGrpcErrorToHttp 将grpc错误转换成http错误
func HandlerGrpcErrorToHttp(c *gin.Context, err error) {
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

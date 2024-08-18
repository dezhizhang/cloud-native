package router

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// Router 接口
type Router interface {
	Route(r *gin.Engine)
}

// RegisterRouter 注册路由
type RegisterRouter struct {
}

func New() *RegisterRouter {
	return &RegisterRouter{}
}

func (*RegisterRouter) Route(ro Router, r *gin.Engine) {
	ro.Route(r)
}

var routers []Router

// Register 注册路由
func Register(router ...Router) {
	routers = append(routers, router...)
}

func InitRouter(r *gin.Engine) {
	for _, router := range routers {
		router.Route(r)
	}
}

// GrpcConfig grpc配置
type GrpcConfig struct {
	Addr         string
	RegisterFunc func(server *grpc.Server)
}

// RegisterGrpc 注册gpc
func RegisterGrpc() {

}

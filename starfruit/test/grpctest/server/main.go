package main

import (
	"context"
	"google.golang.org/grpc"
	"net"
	"starfruit.top/test/grpctest/proto"
)

type Server struct {
	proto.UnimplementedGreeterServer
}

// SayHello rpc服务调用
func (s *Server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: req.Name + "hello"}, nil
}

func main() {
	//1. 实例化grpc
	g := grpc.NewServer()
	// 注册服务
	proto.RegisterGreeterServer(g, &Server{})
	// 启动服务
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic("启动服务失败" + err.Error())
	}

	err = g.Serve(l)
	if err != nil {
		panic(err)
	}
}

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
func (c *Server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	//1. 实例化grpc
	g := grpc.NewServer()
	//2. 注册服务
	proto.RegisterGreeterServer(g, &Server{})
	// 3.启动服务
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	err = g.Serve(l)
	if err != nil {
		panic(err)
	}
}

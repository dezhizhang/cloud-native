package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"starfruit.top/test/interceptor/proto"
)

type Server struct {
	proto.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	//拦戴器处理逻辑
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("--------interceptor")
		return handler(ctx, req)
	}

	opt := grpc.UnaryInterceptor(interceptor)
	g := grpc.NewServer(opt)
	proto.RegisterGreeterServer(g, &Server{})
	listen, err := net.Listen("tcp", ":8082")
	if err != nil {
		panic(err)
	}
	err = g.Serve(listen)
	if err != nil {
		panic(err)
	}
}

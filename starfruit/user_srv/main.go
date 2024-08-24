package main

import (
	"google.golang.org/grpc"
	"net"
	"starfruit.top/user_srv/handler"
	"starfruit.top/user_srv/proto"
)

func main() {

	server := grpc.NewServer()
	// 注册服务
	proto.RegisterUserServer(server, &handler.UserService{})
	// 启动服务
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}

	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}

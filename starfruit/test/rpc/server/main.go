package main

import (
	"net"
	"net/rpc"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello " + request
	return nil
}

func main() {
	//1. 实例化server
	listen, _ := net.Listen("tcp", ":1234")

	//2.注册服务名称
	_ = rpc.RegisterName("HelloService", &HelloService{})
	//3.启动服务
	conn, _ := listen.Accept()
	rpc.ServeConn(conn)
}

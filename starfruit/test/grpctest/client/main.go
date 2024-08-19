package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"starfruit.top/test/grpctest/proto"
)

func main() {
	// 建立链接
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{
		Name: "tom",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}

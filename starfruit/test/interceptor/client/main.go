package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"starfruit.top/test/grpctest/proto"
)

func main() {
	conn, err := grpc.Dial(":8082", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	h, err := c.SayHello(context.Background(), &proto.HelloRequest{
		Name: "tom",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(h.Message)
}

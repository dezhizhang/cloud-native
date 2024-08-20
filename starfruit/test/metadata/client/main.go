package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"starfruit.top/test/metadata/proto"
	"time"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)

	md := metadata.Pairs("timestamp", time.Now().Format("2006-01-02 15:04:05"))
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	h, err := c.SayHello(ctx, &proto.HelloRequest{
		Name: "tom",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(h.GetMessage())
}

package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"starfruit.top/test/grpcstream/proto"
)

func main() {
	conn, err := grpc.Dial(":8082", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	s, err := c.GetStream(context.Background(), &proto.StreamReqData{
		Data: "hello world",
	})
	for {
		r, _ := s.Recv()
		fmt.Println("recv:", r.Data)
	}
}

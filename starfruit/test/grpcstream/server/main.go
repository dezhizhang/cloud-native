package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"starfruit.top/test/grpcstream/proto"
	"time"
)

const PORT = ":8082"

type Server struct {
	proto.UnimplementedGreeterServer
}

func (s *Server) GetStream(in *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		err := res.Send(&proto.StreamResData{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		})
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
	return nil
}

func (s *Server) PutStream(in proto.Greeter_PutStreamServer) error {

	for {
		a, err := in.Recv()
		if err != nil {
			panic(err)
		}
		fmt.Println(a.Data)
	}
	return nil
}

func (s *Server) AllStream(in proto.Greeter_AllStreamServer) error {
	return nil
}

func main() {
	listen, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &Server{})
	err = s.Serve(listen)
	if err != nil {
		panic(err)
	}
}

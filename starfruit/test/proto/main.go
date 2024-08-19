package main

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	buf "starfruit.top/test/proto/proto"
)

func main() {
	req := buf.HelloRequest{
		Name: "John",
	}
	marshal, _ := proto.Marshal(&req)
	fmt.Println(string(marshal))
}

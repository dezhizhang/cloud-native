package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 建立链接
	conn, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		panic("链接失败")
	}

	var reply string

	err = conn.Call("HelloService.Hello", "world", &reply)
	if err != nil {
		panic("调用失败")
	}

	fmt.Println(reply)

}

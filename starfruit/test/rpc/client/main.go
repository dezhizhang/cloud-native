package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
)

type ResponseData struct {
	Data int `json:"data"`
}

func main() {
	req := HttpRequest.NewRequest()
	rsp, _ := req.Get("http://127.0.0.1:8080/add?a=1&b=2")
	body, _ := rsp.Body()

	rspDta := ResponseData{}
	_ = json.Unmarshal(body, &rspDta)
	fmt.Println("-----", rspDta)
}

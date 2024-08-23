package examples

import (
	"gin/proto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleProtobuf(c *gin.Context) {
	teacher := proto.Teacher{
		Name:   "tom",
		Course: []string{"python", "java", "go"},
	}
	c.ProtoBuf(http.StatusOK, &teacher)
}

func main() {
	r := gin.Default()
	r.GET("/protobuf", handleProtobuf)
	_ = r.Run(":8080")
}

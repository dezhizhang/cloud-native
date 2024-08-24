package test

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"starfruit.top/user_srv/driver"
	"starfruit.top/user_srv/model"
	"starfruit.top/user_srv/proto"
	"testing"
)

func TestUser(t *testing.T) {
	for i := 0; i < 11; i++ {
		user := model.User{
			Username: fmt.Sprintf("tom%d", i),
			Mobile:   fmt.Sprintf("1599247844%d", i),
			Password: "123456",
		}
		driver.DB.Save(&user)
	}
}

func TestGetUserList(t *testing.T) {
	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewUserClient(conn)
	users, err := c.GetUserList(context.Background(), &proto.RequestUser{
		Page: 1,
		Size: 20,
	})
	if err != nil {
		panic(err)
	}

	for _, user := range users.User {
		fmt.Println(user)
	}

}

func TestGetUserById(t *testing.T) {
	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewUserClient(conn)
	rsp, _ := c.GetUserById(context.Background(), &proto.IdRequest{
		Id: 20,
	})

	fmt.Println(rsp, err)
}

func TestCreateUser(t *testing.T) {
	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewUserClient(conn)
	user, err := c.CreateUser(context.Background(), &proto.CreateRequest{
		Username: "tom",
		Mobile:   "15992478441",
		Password: "123456",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}

// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userclient

import (
	"context"

	"go-zero/user/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	IdRequest      = user.IdRequest
	UserOkResponse = user.UserOkResponse
	UserRequest    = user.UserRequest
	UserResponse   = user.UserResponse

	User interface {
		// 创建用户
		CreateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserOkResponse, error)
		// 获取用户信息
		GetUserInfo(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserResponse, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

// 创建用户
func (m *defaultUser) CreateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserOkResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.CreateUser(ctx, in, opts...)
}

// 获取用户信息
func (m *defaultUser) GetUserInfo(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}

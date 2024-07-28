package logic

import (
	"context"
	"go-zero/user/models"

	"go-zero/user/internal/svc"
	"go-zero/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateUser 创建用户
func (l *CreateUserLogic) CreateUser(in *user.UserRequest) (*user.UserOkResponse, error) {
	_, err := l.svcCtx.UserModel.Insert(l.ctx, &models.Users{
		Name:     in.Name,
		Password: in.Password,
		Mobile:   in.Mobile,
	})
	if err != nil {
		panic(err)
	}
	return &user.UserOkResponse{Message: "创建用户成功"}, nil
}

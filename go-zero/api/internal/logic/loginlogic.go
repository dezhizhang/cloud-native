package logic

import (
	"context"
	"go-zero/user/userclient"

	"go-zero/api/internal/svc"
	"go-zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	info, err := l.svcCtx.User.GetUserInfo(l.ctx, &userclient.IdRequest{
		Id: "123",
	})
	if err != nil {
		panic(err)
	}
	return &types.Response{Id: info.Id, Name: info.Name, Gender: info.Gender}, nil
}

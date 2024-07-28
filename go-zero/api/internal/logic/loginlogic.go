package logic

import (
	"context"

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

// Login 用户登录
func (l *LoginLogic) Login(req *types.URequest) (resp *types.ULoginResponse, err error) {
	return &types.ULoginResponse{
		Token:   "xiaozhi",
		Message: "登录成功",
	}, nil
}

package logic

import (
	"context"
	"go-zero/user/userclient"

	"go-zero/api/internal/svc"
	"go-zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.URequest) (resp *types.UInfoResponse, err error) {
	// todo: add your logic here and delete this line
	info, err := l.svcCtx.User.GetUserInfo(l.ctx, &userclient.IdRequest{
		Id: req.Id,
	})
	if err != nil {
		panic(err)
	}

	return &types.UInfoResponse{
		Id:     info.Id,
		Name:   info.Name,
		Mobile: "15992478448",
		Gender: "mal",
	}, nil
}

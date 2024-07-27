package logic

import (
	"context"
	"fmt"
	"go-zero/user/pb/user"

	"go-zero/video/internal/svc"
	"go-zero/video/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoLogic {
	return &VideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VideoLogic) Video(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	info, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &user.IdRequest{
		Id: "1",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(info)

	return &types.Response{Message: info.Name}, nil
}

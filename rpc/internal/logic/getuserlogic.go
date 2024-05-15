package logic

import (
	"context"
	"errors"

	"go-zero-prac/rpc/internal/svc"
	"go-zero-prac/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.GetUserReq) (*user.GetUserResp, error) {
	// todo: add your logic here and delete this line
	u, ok := users[in.Id]
	if !ok {
		return nil, errors.New("查无此人")
	}

	return &user.GetUserResp{
		Id:    u.Id,
		Name:  u.Name,
		Phone: u.Phone,
	}, nil

}

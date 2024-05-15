package logic

import (
	"context"
	"go-zero-prac/rpc/userclient"

	"go-zero-prac/api/internal/svc"
	"go-zero-prac/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.GetUserReq) (resp *types.GetUserResp, err error) {
	// todo: add your logic here and delete this line

	//resp = &types.GetUserResp{
	//	Id:    "666",
	//	Name:  "",
	//	Phone: "",
	//}
	getUserResp, err := l.svcCtx.User.GetUser(l.ctx, &userclient.GetUserReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.GetUserResp{
		Id:    getUserResp.Id,
		Name:  getUserResp.Name,
		Phone: getUserResp.Phone,
	}, nil
}

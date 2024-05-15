package logic

import (
	"context"
	"go-zero-prac/models"

	"go-zero-prac/rpc/internal/svc"
	"go-zero-prac/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUSerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUSerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUSerLogic {
	return &CreateUSerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUSerLogic) CreateUSer(in *user.CreateUserReq) (*user.CreateUSerResp, error) {
	// todo: add your logic here and delete this line

	_, err := l.svcCtx.UserModel.Insert(l.ctx, &models.Users{
		Id:    in.Id,
		Name:  in.Name,
		Phone: in.Phone,
	})

	return &user.CreateUSerResp{}, err
}

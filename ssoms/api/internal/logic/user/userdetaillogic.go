package user

import (
	"context"

	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.UserDetailReq) (resp *types.UserDetailReply, err error) {
	user, err := l.svcCtx.UserModel.FindOneByUuid(l.ctx, req.UUID)
	if err != nil {
		return
	}
	resp = &types.UserDetailReply{
		UUID:     user.Uuid,
		Name:     user.Name,
		Avatar:   user.Avatar,
		Mobile:   user.Mobile,
		Status:   user.Status,
		Gender:   user.Gender,
		Password: user.Password,
	}

	return
}

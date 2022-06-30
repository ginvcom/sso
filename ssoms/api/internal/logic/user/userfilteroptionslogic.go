package user

import (
	"context"

	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFilterOptionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFilterOptionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFilterOptionsLogic {
	return &UserFilterOptionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFilterOptionsLogic) UserFilterOptions(req *types.UserFilterOptionsReq) (resp *types.UserFilterOptionsReply, err error) {
	// todo: add your logic here and delete this line

	return
}

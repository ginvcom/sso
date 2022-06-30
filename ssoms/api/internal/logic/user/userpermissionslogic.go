package user

import (
	"context"

	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserPermissionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserPermissionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserPermissionsLogic {
	return &UserPermissionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserPermissionsLogic) UserPermissions(req *types.UserPermissionsReq) (resp *types.UserPermissionsReply, err error) {
	// todo: add your logic here and delete this line

	return
}

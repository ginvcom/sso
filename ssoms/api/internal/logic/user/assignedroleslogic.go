package user

import (
	"context"

	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssignedRolesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssignedRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssignedRolesLogic {
	return &AssignedRolesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssignedRolesLogic) AssignedRoles(req *types.AssignedRolesReq) (resp *types.AssignedRolesReply, err error) {
	// todo: add your logic here and delete this line

	return
}

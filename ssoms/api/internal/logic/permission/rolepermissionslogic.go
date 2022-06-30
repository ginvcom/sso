package permission

import (
	"context"

	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RolePermissionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRolePermissionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RolePermissionsLogic {
	return &RolePermissionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RolePermissionsLogic) RolePermissions(req *types.RolePermissionsReq) (resp *types.RolePermissionsReply, err error) {
	// todo: add your logic here and delete this line

	return
}

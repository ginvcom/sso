package role

import (
	"context"

	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleDetailLogic {
	return &RoleDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleDetailLogic) RoleDetail(req *types.RoleDetailReq) (resp *types.RoleDetailReply, err error) {
	role, err := l.svcCtx.RoleModel.FindOneByRoleUuid(l.ctx, req.RoleUUID)
	if err != nil {
		return
	}
	resp = &types.RoleDetailReply{
		RoleUUID: role.RoleUuid,
		RoleName: role.RoleName,
		Summary:  role.Summary,
	}

	return
}

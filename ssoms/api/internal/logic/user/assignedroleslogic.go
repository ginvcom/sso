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
	// TODO 用户是否删除及状态判断和角色是否删除判断
	roleUUIDArray, err := l.svcCtx.UserToRoleModel.FindRoleUUIDArrByUserUuid(l.ctx, req.UUID)
	if err != nil {
		return
	}
	resp = &types.AssignedRolesReply{
		RoleUUIDArray: *roleUUIDArray,
	}

	return
}

package permission

import (
	"context"
	"errors"

	"sso/service/ssoms/api/internal/svc"
	"sso/service/ssoms/api/internal/types"

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
	role, err := l.svcCtx.RoleModel.FindOneByRoleUuid(l.ctx, req.RoleUUID)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return
	}

	if role.IsDelete == 1 {
		err = errors.New("role not exits")
		return
	}

	permissions, err := l.svcCtx.PermissionModel.PermissionByRoleUUID(l.ctx, req.RoleUUID, req.TopKey)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return
	}

	menuUUIDArray := make([]string, 0, 1)
	actionUUIDArray := make([]string, 0, 1)
	for _, p := range *permissions {
		if p.Type == 2 {
			menuUUIDArray = append(menuUUIDArray, p.ObjectUuid)
		} else if p.Type == 3 {
			actionUUIDArray = append(actionUUIDArray, p.ObjectUuid)
		}
	}
	resp = &types.RolePermissionsReply{
		RoleName:        role.RoleName,
		MenuUUIDArray:   menuUUIDArray,
		ActionUUIDArray: actionUUIDArray,
	}

	return
}

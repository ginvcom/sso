package role

import (
	"context"

	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"
	"sso/ssoms/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoleLogic) UpdateRole(req *types.UpdateRoleReq) (resp *types.UpdateRoleReply, err error) {
	// TODO 增加校验
	role := &model.Role{
		RoleUuid: req.RoleUUID,
		RoleName: req.RoleName,
		Summary:  req.Summary,
	}
	logx.Info(role)
	err = l.svcCtx.RoleModel.Update(l.ctx, role)
	if err != nil {
		return
	}
	resp = &types.UpdateRoleReply{
		Success: true,
	}

	return
}

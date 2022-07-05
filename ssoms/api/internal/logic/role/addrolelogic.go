package role

import (
	"context"

	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"
	"sso/ssoms/model"
	"sso/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRoleLogic {
	return &AddRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddRoleLogic) AddRole(req *types.RoleForm) (resp *types.AddRoleReply, err error) {
	uuid, err := util.UUID()
	if err != nil {
		return
	}
	// TODO 增加校验
	role := &model.Role{
		RoleUuid: uuid,
		RoleName: req.RoleName,
		Summary:  req.Summary,
	}
	logx.Info(role)
	_, err = l.svcCtx.RoleModel.Insert(l.ctx, role)
	if err != nil {
		return
	}
	resp = &types.AddRoleReply{
		RoleUUID: uuid,
	}

	return
}

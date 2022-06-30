package role

import (
	"context"

	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"
	"sso/ssoms/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleListLogic) RoleList(req *types.RoleListReq) (resp *types.RoleListReply, err error) {
	args := &model.RoleListArgs{
		RoleName: req.RoleName,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	total, err := l.svcCtx.RoleModel.ListCount(l.ctx, args)
	if err != nil {
		return
	}
	resp = &types.RoleListReply{
		Total: total,
		List:  make([]types.Role, 0, 1),
	}

	listData, err := l.svcCtx.RoleModel.ListData(l.ctx, args)
	if err != nil {
		return
	}
	for _, role := range *listData {
		item := types.Role{
			RoleUUID: role.RoleUuid,
			RoleName: role.RoleName,
			Summary:  role.Summary,
		}
		resp.List = append(resp.List, item)
	}

	return
}

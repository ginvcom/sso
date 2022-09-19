package user

import (
	"context"
	"errors"

	"sso/service/ssoms/api/internal/svc"
	"sso/service/ssoms/api/internal/types"

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
	user, err := l.svcCtx.UserModel.FindOneByUuid(l.ctx, req.UUID)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return
	}
	if user.IsDelete == 1 {
		err = errors.New("user has been deleted")
		return
	}

	roleUUIDArray, err := l.svcCtx.UserToRoleModel.FindRoleUUIDArrByUserUuid(l.ctx, req.UUID)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return
	}
	// resp = &types.AssignedRolesReply{
	// 	RoleUUIDArray: *roleUUIDArray,
	// }

	opts, err := l.svcCtx.RoleModel.Options(l.ctx)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return
	}
	assignedRoles := make([]types.Option, 0, 1)
	optionsRoles := make([]types.Option, 0, 1)

	for _, opt := range *opts {
		item := types.Option{
			Label: opt.RoleName,
			Value: opt.RoleUUID,
			Extra: opt.Summary,
		}
		assigned := false
		for _, roleUUID := range *roleUUIDArray {
			if opt.RoleUUID == roleUUID {
				assignedRoles = append(assignedRoles, item)
				assigned = true
				break
			}
		}
		if !assigned {
			optionsRoles = append(optionsRoles, item)
		}
	}

	resp = &types.AssignedRolesReply{
		UUID:     user.Uuid,
		Name:     user.Name,
		Assigned: assignedRoles,
		Options:  optionsRoles,
	}

	return
}

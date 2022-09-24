package role

import (
	"context"
	"errors"

	"sso/service/ssoms/api/internal/svc"
	"sso/service/ssoms/api/internal/types"
	"sso/service/ssoms/model"

	"github.com/ginvcom/util"
	"github.com/zeromicro/go-zero/core/logx"
)

type AssignedUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssignedUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssignedUsersLogic {
	return &AssignedUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssignedUsersLogic) AssignedUsers(req *types.AssignedUsersReq) (resp *types.AssignedUsersReply, err error) {
	role, err := l.svcCtx.RoleModel.FindOneByRoleUuid(l.ctx, req.RoleUUID)
	if err != nil {
		l.Logger.Error(err)
		return
	}

	if role.IsDelete == 1 {
		err = errors.New("role not exits")
		return
	}

	total, err := l.svcCtx.UserToRoleModel.CountUserUUIDArrByRoleUuid(l.ctx, req.RoleUUID)
	if err != nil {
		l.Logger.Error(err)
		return
	}
	args := &model.FindUserUUIDArrByRoleUuidArgs{
		RoleUUID: req.RoleUUID,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	userUUIDArray, err := l.svcCtx.UserToRoleModel.FindUserUUIDArrByRoleUuid(l.ctx, args)
	if err != nil {
		l.Logger.Error(err)
		return
	}

	userOptions, err := l.svcCtx.UserModel.UserOptionsInUUIDArray(l.ctx, userUUIDArray)
	if err != nil {
		l.Logger.Error(err)
		return
	}

	users := make([]types.UserOtion, 0, 1)

	for _, user := range *userOptions {
		item := types.UserOtion{
			Name:     user.Name,
			UUID:     user.Uuid,
			Mobile:   util.MobileEncode(user.Mobile),
			Gender:   user.Gender,
			Avatar:   user.Avatar,
			Status:   user.Status,
			IsDelete: user.IsDelete,
		}
		users = append(users, item)
	}

	resp = &types.AssignedUsersReply{
		List:     users,
		Total:    total,
		RoleName: role.RoleName,
	}

	return
}

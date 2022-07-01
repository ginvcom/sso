package role

import (
	"context"

	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"

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
	// TODO 用户是否删除及状态判断和角色是否删除判断
	userUUIDArray, err := l.svcCtx.UserToRoleModel.FindUserUUIDArrByRoleUuid(l.ctx, req.RoleUUID)
	if err != nil {
		return
	}

	userOptions, err := l.svcCtx.UserModel.UserOptionsInUUIDArray(l.ctx, userUUIDArray)
	if err != nil {
		return
	}

	users := make([]types.Option, 0, 1)

	for _, option := range *userOptions {
		item := types.Option{
			Label: option.Name,
			Value: option.UUID,
			Extra: option.Avatar,
		}
		users = append(users, item)
	}

	resp = &types.AssignedUsersReply{
		Users: users,
	}

	return
}

package role

import (
	"context"

	"sso/service/ssoms/api/internal/svc"
	"sso/service/ssoms/api/internal/types"
	"sso/service/ssoms/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssignUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssignUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssignUserLogic {
	return &AssignUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssignUserLogic) AssignUser(req *types.AssignUserReq) (resp *types.AssignUserReply, err error) {
	// TODO 静态职责分离约束判断
	// TODO 用户是否删除及状态判断和角色是否删除判断
	userToRole, err := l.svcCtx.UserToRoleModel.FindOne(l.ctx, req.UserUUID, req.RoleUUID)

	if err != nil && err != model.ErrNotFound {
		logx.WithContext(l.ctx).Error(err)
		return
	}

	err = nil
	if userToRole.UserUuid != "" && userToRole.RoleUuid != "" {
		if userToRole.IsDelete == 1 {
			args := &model.UserToRole{
				UserUuid: req.UserUUID,
				RoleUuid: req.RoleUUID,
				IsDelete: 0,
			}
			err = l.svcCtx.UserToRoleModel.Update(l.ctx, args)
			if err != nil {
				logx.WithContext(l.ctx).Error(err)
				return
			}
		}

		resp = &types.AssignUserReply{
			Success: true,
		}
		return
	}
	args := &model.UserToRole{
		UserUuid: req.UserUUID,
		RoleUuid: req.RoleUUID,
		IsDelete: 0,
	}
	_, err = l.svcCtx.UserToRoleModel.Insert(l.ctx, args)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return
	}

	resp = &types.AssignUserReply{
		Success: true,
	}

	return
}

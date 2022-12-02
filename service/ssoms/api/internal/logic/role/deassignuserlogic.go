package role

import (
	"context"

	"sso/service/ssoms/api/internal/svc"
	"sso/service/ssoms/api/internal/types"
	"sso/service/ssoms/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeassignUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeassignUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeassignUserLogic {
	return &DeassignUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeassignUserLogic) DeassignUser(req *types.DeassignUserReq) (resp *types.DeassignUserReply, err error) {
	// TODO 用户是否删除及状态判断和角色是否删除判断
	userToRole, err := l.svcCtx.UserToRoleModel.FindOne(l.ctx, req.UserUUID, req.RoleUUID)
	if err != nil {
		return
	}

	if userToRole.IsDelete == 0 {
		args := &model.UserToRole{
			UserUuid: req.UserUUID,
			RoleUuid: req.RoleUUID,
			IsDelete: 1,
		}
		err = l.svcCtx.UserToRoleModel.Update(l.ctx, args)
		if err != nil {
			return
		}
	}

	resp = &types.DeassignUserReply{
		Success: true,
	}

	return
}

package user

import (
	"context"
	"time"

	"sso/service/ssoms/api/internal/svc"
	"sso/service/ssoms/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type AssignRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssignRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssignRoleLogic {
	return &AssignRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssignRoleLogic) AssignRole(req *types.AssignRoleReq) (resp *types.AssignRoleReply, err error) {
	// TODO 静态职责分离约束判断
	// TODO 用户是否删除及状态判断和角色是否删除判断
	err = l.svcCtx.UserToRoleModel.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		now := time.Now().Local()
		// 1. 现有user_uuid = req.UUID的记录全部更新为is_delete=1
		err := l.svcCtx.UserToRoleModel.TransDeleteByUserUUID(ctx, session, req.UUID, now)
		if err != nil {
			return err
		}

		if len(req.RoleUUIDArray) == 0 {
			return nil
		}

		// 2. 插入或更新is_delete=0 insert into user_to_role values(xx,xx), (ff, ff) on dumplicate key update `is_delete`=0
		err = l.svcCtx.UserToRoleModel.TransAddRoleUUIDArrayByUserUUID(ctx, session, req.UUID, req.RoleUUIDArray, now)
		return err
	})

	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return
	}

	resp = &types.AssignRoleReply{
		Success: true,
	}
	return
}

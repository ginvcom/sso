package user

import (
	"context"
	"time"

	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"

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

	// 1. 现有user_uuid = req.UUID的记录全部更新为is_delete=1
	// 2. 尝试使用 insert into user_to_role values(xx,xx), (ff, ff) on dumplicate key update `is_delete`=0
	// ---- 第3步不行的话尝试以下：
	// 3. 查询当前用户具有哪些记录currentRoleUUIDArray(包括已删除的)
	// 4. currentRoleUUIDArray的在req.RoleUUIDArray 存在，更新为is_delete=0
	// 5. req.RoleUUIDArray的在currentRoleUUIDArray不存在，插入记录

	// roleUUIDArray, err := l.svcCtx.UserToRoleModel.FindAllRoleUUIDArrByUserUuid(l.ctx, req.UUID)
	// if err != nil {
	// 	return
	// }
	// reqMap := make(map[string]struct{})
	// for _, v := range req.RoleUUIDArray {
	// 	reqMap[v] = struct{}{}
	// }
	// for _, v := range *roleUUIDArray {
	// 	if _, ok := reqMap[v]; ok {

	// 	}
	// }
	err = l.svcCtx.UserToRoleModel.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		now := time.Now().Local()
		// 1. 现有user_uuid = req.UUID的记录全部更新为is_delete=1
		err := l.svcCtx.UserToRoleModel.TransDeleteByUserUUID(ctx, session, req.UUID, now)
		if err != nil {
			return err
		}

		// 2. 尝试使用 insert into user_to_role values(xx,xx), (ff, ff) on dumplicate key update `is_delete`=0

		return nil
	})
	if err != nil {
		return
	}

	resp = &types.AssignRoleReply{
		Success: true,
	}
	return
}

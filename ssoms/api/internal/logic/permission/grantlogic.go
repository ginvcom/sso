package permission

import (
	"context"
	"time"

	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"
	"sso/ssoms/model"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type GrantLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGrantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GrantLogic {
	return &GrantLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GrantLogic) Grant(req *types.GrantReq) (resp *types.GrantReply, err error) {

	err = l.svcCtx.PermissionModel.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		now := time.Now().Local()
		// 1. 现有user_uuid = req.UUID的记录全部更新为is_delete=1
		err := l.svcCtx.PermissionModel.TransDeletePermission(ctx, session, req.RoleUUID, req.TopKey, now)
		if err != nil {
			return err
		}

		dataArray := make([]model.Permission, 0, 1)
		for _, uuid := range req.MenuUUIDArray {
			item := model.Permission{
				ObjectUuid: uuid,
				Type:       2,
				RoleUuid:   req.RoleUUID,
				TopKey:     req.TopKey,
			}
			dataArray = append(dataArray, item)
		}
		for _, uuid := range req.ActionUUIDArray {
			item := model.Permission{
				ObjectUuid: uuid,
				Type:       3,
				RoleUuid:   req.RoleUUID,
				TopKey:     req.TopKey,
			}
			dataArray = append(dataArray, item)
		}

		if len(dataArray) == 0 {
			return nil
		}

		// 2. 插入或更新is_delete=0 insert into user_to_role values(xx,xx), (ff, ff) on dumplicate key update `is_delete`=0
		err = l.svcCtx.PermissionModel.TransAddPermission(ctx, session, dataArray, now)

		return err
	})
	if err != nil {
		return
	}

	resp = &types.GrantReply{
		Success: true,
	}
	return
}

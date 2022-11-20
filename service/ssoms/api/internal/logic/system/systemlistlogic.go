package system

import (
	"context"

	"sso/service/ssoms/api/internal/config"
	"sso/service/ssoms/api/internal/svc"
	"sso/service/ssoms/api/internal/types"
	"sso/service/ssoms/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type SystemListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SystemListLogic {
	return &SystemListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemListLogic) SystemList(req *types.SystemListReq) (resp *types.SystemListReply, err error) {
	l.Logger.Info(l.ctx.Value(config.UUID))
	l.Logger.Info(l.ctx.Value(config.Name))
	args := &model.ObjectListArgs{
		ObjectName: req.SystemName,
		Key:        req.SystemCode,
	}
	listData, err := l.svcCtx.ObjectModel.ListData(l.ctx, args)
	if err != nil {
		return
	}

	resp = &types.SystemListReply{
		List: make([]*types.System, 0, 1),
	}

	for _, obj := range *listData {
		item := types.System{
			UUID:       obj.Uuid,
			SystemCode: obj.Key,
			SystemName: obj.ObjectName,
			Domain:     obj.Identifier,
			Sort:       obj.Sort,
			SubType:    obj.SubType,
			Icon:       obj.Icon,
			Status:     obj.Status,
		}
		resp.List = append(resp.List, &item)
	}

	return
}

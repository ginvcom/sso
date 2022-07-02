package object

import (
	"context"

	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"
	"sso/ssoms/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type ObjectListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewObjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ObjectListLogic {
	return &ObjectListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ObjectListLogic) ObjectList(req *types.ObjectListReq) (resp *types.ObjectListReply, err error) {

	args := &model.ObjectListArgs{
		TopKey:     req.TopKey,
		ObjectName: req.ObjectName,
		Key:        req.Key,
	}
	listData, err := l.svcCtx.ObjectModel.ListData(l.ctx, args)
	if err != nil {
		return
	}

	for _, obj := range *listData {
		item := types.Object{
			UUID:       obj.Uuid,
			ObjectName: obj.ObjectName,
			Domain:     obj.Domain,
			Key:        obj.Key,
			Sort:       obj.Sort,
			Typ:        obj.Type,
			Icon:       obj.Icon,
			Status:     obj.Status,
			PUUID:      obj.Puuid,
		}
		resp.List = append(resp.List, item)
	}

	return
}

package object

import (
	"context"

	"sso/ssoms/api/internal/config"
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
	logx.WithContext(l.ctx).Info(l.ctx.Value(config.UUID))
	logx.WithContext(l.ctx).Info(l.ctx.Value(config.Name))
	args := &model.ObjectListArgs{
		TopKey:     req.TopKey,
		ObjectName: req.ObjectName,
		Key:        req.Key,
	}
	listData, err := l.svcCtx.ObjectModel.ListData(l.ctx, args)
	if err != nil {
		return
	}

	resp = &types.ObjectListReply{
		List: make([]*types.Object, 0, 1),
	}

	for _, obj := range *listData {
		item := types.Object{
			UUID:       obj.Uuid,
			ObjectName: obj.ObjectName,
			Identifier: obj.Identifier,
			Key:        obj.Key,
			Sort:       obj.Sort,
			Typ:        obj.Type,
			SubType:    obj.SubType,
			Icon:       obj.Icon,
			Status:     obj.Status,
			PUUID:      obj.Puuid,
		}
		resp.List = append(resp.List, &item)
	}

	if args.TopKey == "" {
		return
	}
	newList := makeTree(resp.List, "")
	resp.List = newList
	return
}

func makeTree(obs []*types.Object, puuid string) []*types.Object {
	var tree []*types.Object
	for _, child := range obs {
		if child.PUUID == puuid {
			item := child
			item.Children = makeTree(obs, item.UUID)
			tree = append(tree, item)
		}
	}
	return tree
}

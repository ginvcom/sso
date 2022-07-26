package object

import (
	"context"

	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"
	"sso/ssoms/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleOperationsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleOperationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleOperationsLogic {
	return &RoleOperationsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleOperationsLogic) RoleOperations(req *types.RoleOperationsReq) (resp *types.RoleOperationsReply, err error) {
	args := &model.ObjectListArgs{
		TopKey: req.TopKey,
		Status: 1,
	}
	listData, err := l.svcCtx.ObjectModel.ListData(l.ctx, args)
	if err != nil {
		return
	}

	resp = &types.RoleOperationsReply{
		List: make([]*types.ObjectOption, 0, 1),
	}

	for _, obj := range *listData {
		item := types.ObjectOption{
			Value:   obj.Uuid,
			Label:   obj.ObjectName,
			Typ:     obj.Type,
			SubType: obj.SubType,
			PUUID:   obj.Puuid,
		}
		resp.List = append(resp.List, &item)
	}

	newList := makeOptionTree(resp.List, "")
	resp.List = newList
	return
}

func makeOptionTree(obs []*types.ObjectOption, puuid string) []*types.ObjectOption {
	var tree []*types.ObjectOption
	for _, child := range obs {
		if child.PUUID == puuid {
			item := child
			item.Children = makeOptionTree(obs, item.Value)
			item.Apis = setOptionApi(obs, item.Value)
			if item.Typ == 2 {
				tree = append(tree, item)
			}
		}
	}

	return tree
}

func setOptionApi(obs []*types.ObjectOption, puuid string) []*types.ObjectOption {
	tree := make([]*types.ObjectOption, 0, 1)
	for _, child := range obs {
		if child.PUUID == puuid && child.Typ == 3 {
			tree = append(tree, child)
		}
	}
	return tree
}

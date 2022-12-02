package object

import (
	"context"

	"sso/service/ssoms/api/internal/svc"
	"sso/service/ssoms/api/internal/types"
	"sso/service/ssoms/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuOptionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuOptionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuOptionsLogic {
	return &MenuOptionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuOptionsLogic) MenuOptions(req *types.MenuOptionsReq) (resp *types.MenuOptionsReply, err error) {
	args := &model.ObjectListArgs{
		Typ:         2,
		Status:      1,
		ExcludeHide: req.ExcludeHide,
	}
	listData, err := l.svcCtx.ObjectModel.ListData(l.ctx, args)
	if err != nil {
		return
	}

	resp = &types.MenuOptionsReply{
		List: make([]*types.MenuOption, 0, 1),
	}

	for _, obj := range *listData {
		item := types.MenuOption{
			Value: obj.Uuid,
			Label: obj.ObjectName,
			PUUID: obj.Puuid,
		}
		resp.List = append(resp.List, &item)
	}

	newList := makeMenuOptionTree(resp.List, "")
	resp.List = newList

	return
}

func makeMenuOptionTree(obs []*types.MenuOption, puuid string) []*types.MenuOption {
	var tree []*types.MenuOption
	for _, child := range obs {
		if child.PUUID == puuid {
			item := child
			item.Children = makeMenuOptionTree(obs, item.Value)
			tree = append(tree, item)
		}
	}

	return tree
}

package home

import (
	"context"

	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"
	"sso/ssoms/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeStatisticLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeStatisticLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeStatisticLogic {
	return &HomeStatisticLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeStatisticLogic) HomeStatistic() (resp *types.StatisticReply, err error) {
	args := &model.RoleListArgs{}
	roleAmount, err := l.svcCtx.RoleModel.ListCount(l.ctx, args)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return
	}

	userAmountArgs := &model.UserListArgs{}
	userAmount, err := l.svcCtx.UserModel.ListCount(l.ctx, userAmountArgs)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return
	}

	permissionAmount, err := l.svcCtx.PermissionModel.CountPermisson(l.ctx)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return
	}

	objectAmounts, err := l.svcCtx.ObjectModel.CountByType(l.ctx)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return
	}

	resp = &types.StatisticReply{
		RoleAmount:       roleAmount,
		UserAmount:       userAmount,
		PermissionAmount: permissionAmount,
	}

	for _, obj := range *objectAmounts {
		if obj.Typ == 1 {
			resp.SystemAmount = obj.Count
		} else if obj.Typ == 2 {
			resp.MenuAmount = obj.Count
		} else {
			resp.ActionAmount = obj.Count
		}
	}

	return
}

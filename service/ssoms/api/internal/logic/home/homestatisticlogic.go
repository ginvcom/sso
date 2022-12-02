package home

import (
	"context"
	"time"

	"sso/service/ssoms/api/internal/svc"
	"sso/service/ssoms/api/internal/types"
	"sso/service/ssoms/model"

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
		l.Logger.Error(err)
		return
	}

	userAmountArgs := &model.UserListArgs{}
	userAmount, err := l.svcCtx.UserModel.ListCount(l.ctx, userAmountArgs)
	if err != nil {
		l.Logger.Error(err)
		return
	}

	permissionAmount, err := l.svcCtx.PermissionModel.CountPermisson(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return
	}

	objectAmounts, err := l.svcCtx.ObjectModel.CountByType(l.ctx)
	if err != nil {
		l.Logger.Error(err)
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

	month := time.Now().Format("2006-01")
	monthStatistic := &model.Statistic{
		Month:            month,
		UserAmount:       resp.UserAmount,
		RoleAmount:       resp.RoleAmount,
		SystemAmount:     resp.SystemAmount,
		MenuAmount:       resp.MenuAmount,
		ActionAmount:     resp.ActionAmount,
		PermissionAmount: resp.PermissionAmount,
	}
	statistic, err2 := l.svcCtx.StatisticModel.FindOneByMonth(l.ctx, month)
	if err2 != nil {
		if err2 == model.ErrNotFound {
			l.svcCtx.StatisticModel.Insert(l.ctx, monthStatistic)
		} else {
			l.Logger.Error(err2)
		}
	} else {
		if statistic.UserAmount != resp.UserAmount ||
			statistic.RoleAmount != resp.RoleAmount ||
			statistic.SystemAmount != resp.SystemAmount ||
			statistic.MenuAmount != resp.MenuAmount ||
			statistic.ActionAmount != resp.ActionAmount ||
			statistic.PermissionAmount != resp.PermissionAmount {
			monthStatistic.Id = statistic.Id
			err2 := l.svcCtx.StatisticModel.Update(l.ctx, monthStatistic)
			if err2 != nil {
				l.Logger.Error(err2)
			}
		}
	}

	yearsData, err := l.svcCtx.StatisticModel.FindYearsData(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return
	}

	resp.Statistics = make([]types.Statistic, 0, 1)

	for i := (len(yearsData) - 1); i >= 0; i-- {
		v := yearsData[i]
		menuData := types.Statistic{
			Month: v.Month,
			Type:  "菜单",
			Value: v.MenuAmount,
		}
		actionData := types.Statistic{
			Month: v.Month,
			Type:  "按钮",
			Value: v.ActionAmount,
		}
		resp.Statistics = append(resp.Statistics, menuData, actionData)
	}

	return
}

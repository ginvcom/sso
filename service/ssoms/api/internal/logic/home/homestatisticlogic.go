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
	// 1. 统计当前各种元素的数量
	// 统计元素目前共6种，包括 角色、用户、系统、菜单、操作、授权
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

	// 2. 获取当前月份的统计,如果没有记录, 则将第1步统计到的结果插入数据库
	// 如果有记录，但是和当前查出来的结果不一致，则更新当前月份的数据
	// 插入和更新操作出错只记录错误日志，不中断程序继续进行
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
	// 3. 获取12个月内的统计数据
	yearsData, err := l.svcCtx.StatisticModel.FindYearsData(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return
	}

	resp.Statistics = make([]types.Statistic, 0, 1)
	// 数据库倒序拿数据，所以这里从后往前循环拿数组的数据
	// 统计元素目前共6种，包括 角色、用户、系统、菜单、操作、授权
	// 只拿取了菜单和按钮,实际可以按需要拿取更多的元素
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

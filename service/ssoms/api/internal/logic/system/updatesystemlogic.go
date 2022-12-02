package system

import (
	"context"
	"errors"

	"sso/service/ssoms/api/internal/config"
	"sso/service/ssoms/api/internal/svc"
	"sso/service/ssoms/api/internal/types"
	"sso/service/ssoms/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSystemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSystemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSystemLogic {
	return &UpdateSystemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSystemLogic) UpdateSystem(req *types.UpdateSystemReq) (resp *types.UpdateSystemReply, err error) {
	// TODO 增加校验
	userUUID := l.ctx.Value(config.UUID).(string)
	if userUUID == "" {
		l.Logger.Info("missing user info")
		err = errors.New("missing user info")
		return
	}

	userName := l.ctx.Value(config.Name).(string)
	if userName == "" {
		l.Logger.Info("missing user info")
		err = errors.New("missing user info")
		return
	}

	// 防止重复添加的校验, 状态忽略，只看没有删除的
	isExistArgs := &model.ObjectIsExistArgs{
		Key:         req.SystemCode,
		Typ:         1,
		SubType:     req.SubType,
		TopKey:      req.SystemCode,
		ExcludeUUID: req.UUID,
	}

	exist, err := l.svcCtx.ObjectModel.IsExist(l.ctx, isExistArgs)
	if err != nil {
		l.Logger.Error(err)
		return
	}

	if exist {
		err = errors.New("object already exists")
		return
	}

	obj := &model.Object{
		Uuid:       req.UUID,
		ObjectName: req.SystemName,
		Identifier: req.Domain,
		Key:        req.SystemCode,
		Sort:       req.Sort,
		Type:       1,
		SubType:    req.SubType,
		Icon:       req.Icon,
		Status:     req.Status,
	}
	logx.Info(obj)
	err = l.svcCtx.ObjectModel.Update(l.ctx, obj)
	if err != nil {
		l.Logger.Error("更新系统出错", err)
		return
	}

	resp = &types.UpdateSystemReply{
		Success: true,
	}

	return
}

package system

import (
	"context"
	"errors"

	"sso/service/ssoms/api/internal/config"
	"sso/service/ssoms/api/internal/svc"
	"sso/service/ssoms/api/internal/types"
	"sso/service/ssoms/model"

	"github.com/ginvcom/util"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddSystemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSystemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSystemLogic {
	return &AddSystemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSystemLogic) AddSystem(req *types.SystemForm) (resp *types.AddSystemReply, err error) {
	uuid, err := util.UUID()
	if err != nil {
		return
	}

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

	logx.Info(req)

	var typ int64 = 1 // 系统的type值是1

	// 防止重复添加的校验, 状态忽略，只看没有删除的
	isExistArgs := &model.ObjectIsExistArgs{
		Key:     req.SystemCode,
		Typ:     typ, // 系统的type值是1
		SubType: req.SubType,
		TopKey:  req.SystemCode,
	}

	exist, err := l.svcCtx.ObjectModel.IsExist(l.ctx, isExistArgs)
	if err != nil {
		return
	}

	if exist {
		err = errors.New("object already exists")
		return
	}

	object := &model.Object{
		Uuid:       uuid,
		ObjectName: req.SystemName,
		Identifier: req.Domain,
		Key:        req.SystemCode,
		Sort:       req.Sort,
		Type:       typ,
		SubType:    req.SubType,
		Icon:       req.Icon,
		Status:     req.Status,
		TopKey:     req.SystemCode,
	}

	_, err = l.svcCtx.ObjectModel.Insert(l.ctx, object)
	if err != nil {
		l.Logger.Error("新增系统出错", err)
		return
	}

	resp = &types.AddSystemReply{
		UUID: uuid,
	}

	return
}

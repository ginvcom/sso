package object

import (
	"context"
	"errors"

	"sso/service/ssoms/api/internal/svc"
	"sso/service/ssoms/api/internal/types"
	"sso/service/ssoms/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateObjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateObjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateObjectLogic {
	return &UpdateObjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateObjectLogic) UpdateObject(req *types.ObjectForm) (resp *types.UpdateObjectReply, err error) {
	// TODO 增加校验

	// 防止重复添加的校验, 状态忽略，只看没有删除的
	isExistArgs := &model.ObjectIsExistArgs{
		Key:         req.Key,
		Typ:         req.Typ,
		SubType:     req.SubType,
		TopKey:      req.TopKey,
		ExcludeUUID: req.UUID,
	}

	exist, err := l.svcCtx.ObjectModel.IsExist(l.ctx, isExistArgs)
	if err != nil {
		return
	}

	if exist {
		err = errors.New("object already exists")
		return
	}

	obj := &model.Object{
		Uuid:       req.UUID,
		ObjectName: req.ObjectName,
		Identifier: req.Identifier,
		Key:        req.Key,
		Sort:       req.Sort,
		Type:       req.Typ,
		SubType:    req.SubType,
		Extra:      req.Extra,
		Icon:       req.Icon,
		Status:     req.Status,
		Puuid:      req.PUUID,
	}
	logx.Info(obj)
	err = l.svcCtx.ObjectModel.Update(l.ctx, obj)
	if err != nil {
		return
	}

	resp = &types.UpdateObjectReply{
		Success: true,
	}

	return
}

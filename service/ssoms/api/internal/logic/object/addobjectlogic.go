package object

import (
	"context"
	"errors"

	"sso/service/ssoms/api/internal/svc"
	"sso/service/ssoms/api/internal/types"
	"sso/service/ssoms/model"

	"github.com/ginvcom/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddObjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddObjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddObjectLogic {
	return &AddObjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddObjectLogic) AddObject(req *types.ObjectForm) (resp *types.AddObjectReply, err error) {
	uuid, err := util.UUID()
	if err != nil {
		return
	}

	logx.Info(req)
	if req.TopKey == "" {

		err = errors.New("topKey is required oh")
		return
	}

	if req.Typ == 1 && req.TopKey != req.Key {
		err = errors.New("topKey is must equal to key")
		return
	}

	// TODO 增加校验

	// 防止重复添加的校验, 状态忽略，只看没有删除的
	isExistArgs := &model.ObjectIsExistArgs{
		Key:     req.Key,
		Typ:     req.Typ,
		SubType: req.SubType,
		TopKey:  req.TopKey,
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
		TopKey:     req.TopKey,
	}

	logx.Info(object)
	_, err = l.svcCtx.ObjectModel.Insert(l.ctx, object)
	if err != nil {
		return
	}
	resp = &types.AddObjectReply{
		UUID: uuid,
	}

	return
}

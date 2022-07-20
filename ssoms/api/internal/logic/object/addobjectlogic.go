package object

import (
	"context"
	"errors"

	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"
	"sso/ssoms/model"

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
	object := &model.Object{
		Uuid:       uuid,
		ObjectName: req.ObjectName,
		Domain:     req.Domain,
		Key:        req.Key,
		Sort:       req.Sort,
		Type:       req.Typ,
		Icon:       req.Icon,
		Status:     req.Status,
		Puuid:      req.PUUID,
	}

	logx.Info(object)
	_, err = l.svcCtx.ObjectModel.Insert(l.ctx, object, req.TopKey)
	if err != nil {
		return
	}
	resp = &types.AddObjectReply{
		UUID: uuid,
	}

	return
}

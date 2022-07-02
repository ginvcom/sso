package object

import (
	"context"

	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"
	"sso/ssoms/model"

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

func (l *UpdateObjectLogic) UpdateObject(req *types.UpdateObjectReq) (resp *types.UpdateObjectReply, err error) {
	// TODO 增加校验
	obj := &model.Object{
		Uuid:       req.UUID,
		ObjectName: req.ObjectName,
		Domain:     req.Domain,
		Key:        req.Key,
		Sort:       req.Sort,
		Type:       req.Typ,
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

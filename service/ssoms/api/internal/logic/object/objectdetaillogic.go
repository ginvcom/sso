package object

import (
	"context"

	"sso/service/ssoms/api/internal/svc"
	"sso/service/ssoms/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ObjectDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewObjectDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ObjectDetailLogic {
	return &ObjectDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ObjectDetailLogic) ObjectDetail(req *types.ObjectDetailReq) (resp *types.ObjectForm, err error) {
	obj, err := l.svcCtx.ObjectModel.FindOneByUuid(l.ctx, req.UUID)
	if err != nil {
		return
	}
	resp = &types.ObjectForm{
		UUID:       obj.Uuid,
		ObjectName: obj.ObjectName,
		Identifier: obj.Identifier,
		Key:        obj.Key,
		Sort:       obj.Sort,
		Typ:        obj.Type,
		SubType:    obj.SubType,
		Extra:      obj.Extra,
		Icon:       obj.Icon,
		Status:     obj.Status,
		PUUID:      obj.Puuid,
		TopKey:     obj.TopKey,
	}

	return
}

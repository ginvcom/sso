package system

import (
	"context"

	"sso/service/ssoms/api/internal/svc"
	"sso/service/ssoms/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SystemDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SystemDetailLogic {
	return &SystemDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemDetailLogic) SystemDetail(req *types.SystemDetailReq) (resp *types.SystemForm, err error) {
	obj, err := l.svcCtx.ObjectModel.FindOneByUuid(l.ctx, req.UUID)
	if err != nil {
		return
	}
	resp = &types.SystemForm{
		UUID:       obj.Uuid,
		SystemCode: obj.Key,
		SystemName: obj.ObjectName,
		Domain:     obj.Identifier,
		Sort:       obj.Sort,
		SubType:    obj.SubType,
		Icon:       obj.Icon,
		Status:     obj.Status,
	}

	return
}

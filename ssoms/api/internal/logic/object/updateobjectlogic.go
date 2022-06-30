package object

import (
	"context"

	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}

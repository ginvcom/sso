package object

import (
	"context"

	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"

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

func (l *AddObjectLogic) AddObject(req *types.AddObjectReq) (resp *types.AddObjectReply, err error) {
	// todo: add your logic here and delete this line

	return
}

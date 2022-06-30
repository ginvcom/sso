package object

import (
	"context"

	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteObjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteObjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteObjectLogic {
	return &DeleteObjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteObjectLogic) DeleteObject(req *types.DeleteObjectReq) (resp *types.DeleteObjectReply, err error) {
	// todo: add your logic here and delete this line

	return
}

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
	// TODO 存在子对象, 不可以被删除
	// TODO key=ssoms系统的不可以被删
	err = l.svcCtx.ObjectModel.Delete(l.ctx, req.UUID)
	if err != nil {
		return
	}
	resp = &types.DeleteObjectReply{
		Success: true,
	}

	return
}

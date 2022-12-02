package system

import (
	"context"
	"errors"

	"sso/service/ssoms/api/internal/config"
	"sso/service/ssoms/api/internal/svc"
	"sso/service/ssoms/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSystemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSystemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSystemLogic {
	return &DeleteSystemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSystemLogic) DeleteSystem(req *types.DeleteSystemReq) (resp *types.DeleteSystemReply, err error) {
	// TODO 存在子对象, 不可以被删除
	userUUID := l.ctx.Value(config.UUID).(string)
	if userUUID == "" {
		l.Logger.Info("missing user info")
		err = errors.New("missing user info")
		return
	}

	userName := l.ctx.Value(config.Name).(string)
	if userName == "" {
		l.Logger.Info("missing user info")
		err = errors.New("missing user info")
		return
	}
	_, err = l.svcCtx.ObjectModel.FindOneByUuid(l.ctx, req.UUID)
	if err != nil {
		l.Logger.Error(err)
		return
	}

	// TODO key=ssoms系统的不可以被删
	err = l.svcCtx.ObjectModel.Delete(l.ctx, req.UUID)
	if err != nil {
		return
	}
	resp = &types.DeleteSystemReply{
		Success: true,
	}

	return
}

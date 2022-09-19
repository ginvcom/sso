package user

import (
	"context"
	"errors"

	"sso/service/ssoms/api/internal/config"
	"sso/service/ssoms/api/internal/svc"
	"sso/service/ssoms/api/internal/types"
	"sso/service/ssoms/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoEditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoEditLogic {
	return &InfoEditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoEditLogic) InfoEdit(req *types.InfoEditReq) (resp *types.InfoEditReply, err error) {
	uuid := l.ctx.Value(config.UUID).(string)
	if uuid == "" {
		logx.WithContext(l.ctx).Info("missing user info")
		err = errors.New("missing user info")
		return
	}

	if req.Introduction == "" {
		err = errors.New("introduction is required")
		return
	}

	user := &model.User{
		Uuid:         uuid,
		Introduction: req.Introduction,
	}
	logx.Info(user)
	err = l.svcCtx.UserModel.BasicUpdate(l.ctx, user)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return
	}

	resp = &types.InfoEditReply{
		Success: true,
	}

	return
}

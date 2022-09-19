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

type AvatarUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAvatarUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AvatarUploadLogic {
	return &AvatarUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AvatarUploadLogic) AvatarUpload(req *types.AvatarUploadReq) (resp *types.AvatarUploadReply, err error) {
	uuid := l.ctx.Value(config.UUID).(string)
	if uuid == "" {
		logx.WithContext(l.ctx).Info("missing user info")
		err = errors.New("missing user info")
		return
	}

	if req.Avatar == "" {
		err = errors.New("avatar is required")
		return
	}

	user := &model.User{
		Uuid:   uuid,
		Avatar: req.Avatar,
	}
	err = l.svcCtx.UserModel.BasicUpdate(l.ctx, user)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return
	}

	resp = &types.AvatarUploadReply{
		Success: true,
	}

	return
}

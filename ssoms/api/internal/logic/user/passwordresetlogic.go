package user

import (
	"context"

	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PasswordResetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPasswordResetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PasswordResetLogic {
	return &PasswordResetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PasswordResetLogic) PasswordReset() (resp *types.PasswordResetReply, err error) {
	// todo: add your logic here and delete this line

	return
}

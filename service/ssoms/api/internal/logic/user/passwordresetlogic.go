package user

import (
	"context"
	"errors"

	"sso/service/ssoms/api/internal/config"
	"sso/service/ssoms/api/internal/svc"
	"sso/service/ssoms/api/internal/types"

	"github.com/ginvcom/util"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
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

func (l *PasswordResetLogic) PasswordReset(req *types.PasswordResetReq) (resp *types.PasswordResetReply, err error) {
	// 拿到jwt换取的uuid
	uuid := l.ctx.Value(config.UUID).(string)
	if uuid == "" {
		logx.WithContext(l.ctx).Info("missing user info")
		err = errors.New("missing user info")
		return
	}

	if req.Password == "" {
		err = errors.New("password required")
		return
	}

	// 判断新密码和确认密码是否一致
	if req.Password != req.ConfirmPassword {
		err = errors.New("new password and confirmation password are not equal")
		return
	}

	// 判断新密码和确认密码是否一致
	if req.Password == req.OldPassword {
		err = errors.New("new password and old password are not equal")
		return
	}

	// 判断旧密码是否正确
	user, err := l.svcCtx.UserModel.FindOneByUuid(l.ctx, uuid)
	if err != nil {
		if err == sqlc.ErrNotFound {
			err = errors.New("account error")
		}

		return
	}

	currentPassword := util.MD5(req.OldPassword + user.Salt)

	if currentPassword != user.Password {
		err = errors.New("old password error")
		return
	}

	salt := util.RandomStr(6)
	password := util.MD5(req.Password + salt)
	err = l.svcCtx.UserModel.PasswordReset(l.ctx, uuid, password, salt)
	if err != nil {
		return
	}

	resp = &types.PasswordResetReply{
		Success: true,
	}

	return
}

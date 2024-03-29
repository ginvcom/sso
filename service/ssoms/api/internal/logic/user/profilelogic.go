package user

import (
	"context"
	"errors"

	"sso/service/ssoms/api/internal/config"
	"sso/service/ssoms/api/internal/svc"
	"sso/service/ssoms/api/internal/types"

	"github.com/ginvcom/util"
	"github.com/zeromicro/go-zero/core/logx"
)

type ProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProfileLogic {
	return &ProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProfileLogic) Profile() (resp *types.UserForm, err error) {
	uuid := l.ctx.Value(config.UUID).(string)
	if uuid == "" {
		l.Logger.Info("missing user info")
		err = errors.New("missing user info")
		return
	}

	user, err := l.svcCtx.UserModel.FindOneByUuid(l.ctx, uuid)
	if err != nil {
		l.Logger.Error(err)
		return
	}

	resp = &types.UserForm{
		UUID:         user.Uuid,
		Name:         user.Name,
		Avatar:       user.Avatar,
		Mobile:       user.Mobile,
		Status:       user.Status,
		Gender:       user.Gender,
		Birth:        util.TimeToString(user.Birth, "YYYY-MM-DD"),
		Introduction: user.Introduction,
	}

	return
}

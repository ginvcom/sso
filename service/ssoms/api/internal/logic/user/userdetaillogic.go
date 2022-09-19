package user

import (
	"context"

	"sso/service/ssoms/api/internal/svc"
	"sso/service/ssoms/api/internal/types"

	"github.com/ginvcom/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.UserDetailReq) (resp *types.UserForm, err error) {
	user, err := l.svcCtx.UserModel.FindOneByUuid(l.ctx, req.UUID)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
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

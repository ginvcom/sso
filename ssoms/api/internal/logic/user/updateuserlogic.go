package user

import (
	"context"

	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"
	"sso/ssoms/model"
	"sso/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserReq) (resp *types.UpdateUserReply, err error) {
	// TODO 增加校验
	birth, err := util.StringToTime(req.Birth, "YYYY-MM-DD")
	if err != nil {
		return
	}
	user := &model.User{
		Name:     req.Name,
		Mobile:   req.Mobile,
		Avatar:   req.Avatar,
		Birth:    birth,
		Gender:   req.Gender,
		Password: req.Password,
		Uuid:     req.UUID,
	}
	logx.Info(user)
	l.svcCtx.UserModel.Insert(l.ctx, user)
	resp = &types.UpdateUserReply{
		Success: true,
	}
	return
}

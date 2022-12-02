package user

import (
	"context"

	"sso/service/ssoms/api/internal/svc"
	"sso/service/ssoms/api/internal/types"
	"sso/service/ssoms/model"

	"github.com/ginvcom/util"

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
		l.Logger.Info(err)
		return
	}

	user := &model.User{
		Uuid:         req.UUID,
		Name:         req.Name,
		Mobile:       req.Mobile,
		Avatar:       req.Avatar,
		Birth:        birth,
		Introduction: req.Introduction,
		Status:       req.Status,
		Gender:       req.Gender,
		Password:     req.Password,
	}
	if req.Password != "" {
		salt := util.RandomStr(6)
		user.Salt = salt
		user.Password = util.MD5(req.Password + salt)
	}
	logx.Info(user)
	err = l.svcCtx.UserModel.Update(l.ctx, user)
	if err != nil {
		l.Logger.Error(err)
		return
	}

	resp = &types.UpdateUserReply{
		Success: true,
	}

	return
}

package user

import (
	"context"

	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"
	"sso/ssoms/model"
	"sso/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req *types.AddUserReq) (resp *types.AddUserReply, err error) {
	uuid, err := util.UUID()
	if err != nil {
		return
	}
	birth, err := util.StringToTime(req.Birth, "YYYY-MM-DD")
	if err != nil {
		return
	}
	// TODO 增加校验
	user := &model.User{
		Name:     req.Name,
		Mobile:   req.Mobile,
		Avatar:   req.Avatar,
		Birth:    birth,
		Gender:   req.Gender,
		Password: req.Password,
		Uuid:     uuid,
		Status:   req.Status,
	}
	logx.Info(user)
	_, err = l.svcCtx.UserModel.Insert(l.ctx, user)
	if err != nil {
		return
	}
	resp = &types.AddUserReply{
		UUID: uuid,
	}

	return
}

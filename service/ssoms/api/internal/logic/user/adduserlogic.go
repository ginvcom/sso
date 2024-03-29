package user

import (
	"context"

	"sso/service/ssoms/api/internal/svc"
	"sso/service/ssoms/api/internal/types"
	"sso/service/ssoms/model"

	"github.com/ginvcom/util"

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

func (l *AddUserLogic) AddUser(req *types.UserForm) (resp *types.AddUserReply, err error) {
	uuid, err := util.UUID()
	if err != nil {
		return
	}
	birth, err := util.StringToTime(req.Birth, "YYYY-MM-DD")
	if err != nil {
		l.Logger.Error("err")
		return
	}
	// TODO 增加校验

	salt := util.RandomStr(6)
	password := util.MD5(req.Password + salt)
	user := &model.User{
		Name:         req.Name,
		Mobile:       req.Mobile,
		Avatar:       req.Avatar,
		Birth:        birth,
		Gender:       req.Gender,
		Password:     password,
		Salt:         salt,
		Uuid:         uuid,
		Introduction: req.Introduction,
		Status:       req.Status,
	}
	l.Logger.Info(user)
	_, err = l.svcCtx.UserModel.Insert(l.ctx, user)
	if err != nil {
		l.Logger.Error(err)
		return
	}

	resp = &types.AddUserReply{
		UUID: uuid,
	}

	return
}

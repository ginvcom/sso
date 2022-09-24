package user

import (
	"context"

	"sso/service/ssoms/api/internal/svc"
	"sso/service/ssoms/api/internal/types"
	"sso/service/ssoms/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFilterOptionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFilterOptionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFilterOptionsLogic {
	return &UserFilterOptionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFilterOptionsLogic) UserFilterOptions(req *types.UserFilterOptionsReq) (resp *types.UserFilterOptionsReply, err error) {

	l.Logger.Info(req.Name)

	args := &model.UserFilterOptionsArgs{
		Name:  req.Name,
		Limit: 20,
	}

	options, err := l.svcCtx.UserModel.FilterOptions(l.ctx, args)
	if err != nil {
		l.Logger.Error(err)
		return
	}

	resp = &types.UserFilterOptionsReply{
		Options: make([]types.Option, 0, 1),
	}

	for _, option := range *options {
		item := types.Option{
			Label: option.Name,
			Value: option.UUID,
			Extra: option.Avatar,
		}
		resp.Options = append(resp.Options, item)
	}

	return
}

package role

import (
	"context"

	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleOptionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleOptionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleOptionsLogic {
	return &RoleOptionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleOptionsLogic) RoleOptions() (resp *types.OptionsReply, err error) {
	options, err := l.svcCtx.RoleModel.Options(l.ctx)
	if err != nil {
		return
	}

	resp = &types.OptionsReply{
		Options: make([]types.Option, 0, 1),
	}
	for _, option := range *options {
		item := types.Option{
			Label: option.RoleName,
			Value: option.RoleUUID,
			Extra: option.Summary,
		}
		resp.Options = append(resp.Options, item)
	}

	return
}

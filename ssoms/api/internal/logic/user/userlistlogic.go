package user

import (
	"context"

	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"
	"sso/ssoms/model"
	"sso/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserList(req *types.UserListReq) (resp *types.UserListReply, err error) {

	args := &model.UserListArgs{
		Name:     req.Name,
		Mobile:   req.Mobile,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	total, err := l.svcCtx.UserModel.ListCount(l.ctx, args)
	if err != nil {
		return
	}
	resp = &types.UserListReply{
		Total: total,
		List:  make([]types.User, 0, 1),
	}

	listData, err := l.svcCtx.UserModel.ListData(l.ctx, args)
	if err != nil {
		return
	}
	for _, user := range *listData {
		item := types.User{
			UUID:   user.Uuid,
			Name:   user.Name,
			Avatar: user.Avatar,
			Mobile: util.MobileEncode(user.Mobile),
			Status: user.Status,
			Gender: user.Gender,
		}
		resp.List = append(resp.List, item)
	}
	return
}

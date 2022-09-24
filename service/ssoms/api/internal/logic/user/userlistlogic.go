package user

import (
	"context"

	"sso/service/ssoms/api/internal/svc"
	"sso/service/ssoms/api/internal/types"
	"sso/service/ssoms/model"

	"github.com/ginvcom/util"

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
	l.Logger.Info("args", args)
	total, err := l.svcCtx.UserModel.ListCount(l.ctx, args)
	if err != nil {
		l.Logger.Error(err)
		return
	}

	resp = &types.UserListReply{
		Total: total,
		List:  make([]types.User, 0, 1),
	}

	listData, err := l.svcCtx.UserModel.ListData(l.ctx, args)
	if err != nil {
		l.Logger.Error(err)
		return
	}

	var userUUIDArray []string
	for _, role := range *listData {
		userUUIDArray = append(userUUIDArray, role.Uuid)
	}
	rolesResp, err := l.svcCtx.UserToRoleModel.ListRoleByUserUUidArray(l.ctx, &userUUIDArray)
	if err != nil {
		l.Logger.Error(err)
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
		item.Roles = make([]types.Option, 0, 1)
		for _, roleItem := range *rolesResp {
			if roleItem.UserUuid == item.UUID {
				role := types.Option{
					Label: roleItem.RoleName,
					Value: roleItem.RoleUUID,
				}
				item.Roles = append(item.Roles, role)
			}
		}
		resp.List = append(resp.List, item)
	}
	return
}

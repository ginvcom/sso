package role

import (
	"context"

	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssignUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssignUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssignUserLogic {
	return &AssignUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssignUserLogic) AssignUser(req *types.AssignUserReq) (resp *types.AssignUserReply, err error) {
	// todo: add your logic here and delete this line

	return
}

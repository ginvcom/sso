package object

import (
	"context"

	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleOperationsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleOperationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleOperationsLogic {
	return &RoleOperationsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleOperationsLogic) RoleOperations(req *types.RoleOperationsReq) (resp *types.RoleOperationsReply, err error) {
	// todo: add your logic here and delete this line

	return
}

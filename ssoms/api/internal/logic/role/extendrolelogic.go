package role

import (
	"context"

	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExtendRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExtendRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExtendRoleLogic {
	return &ExtendRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExtendRoleLogic) ExtendRole(req *types.AddInheritanceReq) (resp *types.AddInheritanceReply, err error) {
	// todo: add your logic here and delete this line

	return
}

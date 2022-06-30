package permission

import (
	"context"

	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GrantLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGrantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GrantLogic {
	return &GrantLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GrantLogic) Grant(req *types.GrantReq) (resp *types.GrantReply, err error) {
	// todo: add your logic here and delete this line

	return
}

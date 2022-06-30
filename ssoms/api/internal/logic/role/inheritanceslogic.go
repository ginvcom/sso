package role

import (
	"context"

	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InheritancesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInheritancesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InheritancesLogic {
	return &InheritancesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InheritancesLogic) Inheritances(req *types.InheritancesReq) (resp *types.InheritancesReply, err error) {
	// todo: add your logic here and delete this line

	return
}

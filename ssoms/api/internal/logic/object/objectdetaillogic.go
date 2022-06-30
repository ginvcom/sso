package object

import (
	"context"

	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ObjectDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewObjectDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ObjectDetailLogic {
	return &ObjectDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ObjectDetailLogic) ObjectDetail(req *types.ObjectDetailReq) (resp *types.ObjectDetailReply, err error) {
	// todo: add your logic here and delete this line

	return
}

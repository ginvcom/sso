package object

import (
	"context"

	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ObjectListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewObjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ObjectListLogic {
	return &ObjectListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ObjectListLogic) ObjectList(req *types.ObjectListReq) (resp *types.ObjectListReply, err error) {
	// todo: add your logic here and delete this line

	return
}

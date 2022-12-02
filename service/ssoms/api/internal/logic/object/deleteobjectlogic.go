package object

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"sso/service/ssoms/api/internal/config"
	"sso/service/ssoms/api/internal/svc"
	"sso/service/ssoms/api/internal/types"
	"sso/service/ssoms/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteObjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteObjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteObjectLogic {
	return &DeleteObjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteObjectLogic) DeleteObject(req *types.DeleteObjectReq) (resp *types.DeleteObjectReply, err error) {
	// TODO 存在子对象, 不可以被删除
	userUUID := l.ctx.Value(config.UUID).(string)
	if userUUID == "" {
		l.Logger.Info("missing user info")
		err = errors.New("missing user info")
		return
	}

	userName := l.ctx.Value(config.Name).(string)
	if userName == "" {
		l.Logger.Info("missing user info")
		err = errors.New("missing user info")
		return
	}
	object, err := l.svcCtx.ObjectModel.FindOneByUuid(l.ctx, req.UUID)
	if err != nil {
		l.Logger.Error(err)
		return
	}
	// TODO key=ssoms系统的不可以被删
	err = l.svcCtx.ObjectModel.Delete(l.ctx, req.UUID)
	if err != nil {
		return
	}
	resp = &types.DeleteObjectReply{
		Success: true,
	}
	if object.Type != 1 {
		summary := "删除菜单"
		if object.Type == 2 {
			if object.SubType == 2 {
				summary = "删除菜单组"
			} else if object.SubType == 3 {
				summary = "删除隐藏菜单"
			}
		} else {
			if object.SubType == 1 {
				summary = "删除[GET]操作"
			} else if object.SubType == 2 {
				summary = "删除[POST]操作"
			} else if object.SubType == 3 {
				summary = "删除[PUT]操作"
			} else if object.SubType == 4 {
				summary = "删除[PAT]操作"
			} else if object.SubType == 5 {
				summary = "删除[DEL]操作"
			}
		}
		logSummary := fmt.Sprintf("%s: %s", summary, object.ObjectName)
		logData, err2 := json.Marshal(*object)
		if err2 != nil {
			err = err2
			l.Logger.Error(err)
			return
		}
		l.Logger.Info(logData)
		objectLogData := &model.ObjectLog{
			ObjectUuid: req.UUID,
			UserUuid:   userUUID,
			UserName:   userName,
			Type:       object.Type,
			LogType:    3,
			LogSummary: logSummary,
			LogData:    "",
		}

		_, err3 := l.svcCtx.ObjectLogModel.Insert(l.ctx, objectLogData)
		if err3 != nil {
			l.Logger.Error(err3)
		}
	}
	return
}

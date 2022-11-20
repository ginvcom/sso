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

type UpdateObjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateObjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateObjectLogic {
	return &UpdateObjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateObjectLogic) UpdateObject(req *types.UpdateObjectReq) (resp *types.UpdateObjectReply, err error) {
	// TODO 增加校验
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

	// 防止重复添加的校验, 状态忽略，只看没有删除的
	isExistArgs := &model.ObjectIsExistArgs{
		Key:         req.Key,
		Typ:         req.Typ,
		SubType:     req.SubType,
		TopKey:      req.TopKey,
		ExcludeUUID: req.UUID,
	}

	exist, err := l.svcCtx.ObjectModel.IsExist(l.ctx, isExistArgs)
	if err != nil {
		return
	}

	if exist {
		err = errors.New("object already exists")
		return
	}

	obj := &model.Object{
		Uuid:       req.UUID,
		ObjectName: req.ObjectName,
		Identifier: req.Identifier,
		Key:        req.Key,
		Sort:       req.Sort,
		Type:       req.Typ,
		SubType:    req.SubType,
		Extra:      req.Extra,
		Icon:       req.Icon,
		Status:     req.Status,
		Puuid:      req.PUUID,
	}
	logx.Info(obj)
	err = l.svcCtx.ObjectModel.Update(l.ctx, obj)
	if err != nil {
		return
	}

	if req.Typ != 1 {
		summary := "更新菜单"
		if req.Typ == 2 {
			if req.SubType == 2 {
				summary = "更新菜单组"
			} else if req.SubType == 3 {
				summary = "更新隐藏菜单"
			}
		} else {
			if req.SubType == 1 {
				summary = "更新[GET]操作"
			} else if req.SubType == 2 {
				summary = "更新[POST]操作"
			} else if req.SubType == 3 {
				summary = "更新[PUT]操作"
			} else if req.SubType == 4 {
				summary = "更新[PAT]操作"
			} else if req.SubType == 5 {
				summary = "更新[DEL]操作"
			}
		}
		logSummary := fmt.Sprintf("%s: %s", summary, req.ObjectName)
		logData, err2 := json.Marshal(*req)
		if err2 != nil {
			err = err2
			l.Logger.Error(err)
			return
		}

		objectLogData := &model.ObjectLog{
			ObjectUuid: req.UUID,
			UserUuid:   userUUID,
			UserName:   userName,
			Type:       req.Typ,
			LogType:    2,
			LogSummary: logSummary,
			LogData:    string(logData),
		}

		_, err3 := l.svcCtx.ObjectLogModel.Insert(l.ctx, objectLogData)
		if err3 != nil {
			l.Logger.Error(err3)
		}
	}

	resp = &types.UpdateObjectReply{
		Success: true,
	}

	return
}

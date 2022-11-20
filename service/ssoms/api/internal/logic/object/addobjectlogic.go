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

	"github.com/ginvcom/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddObjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddObjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddObjectLogic {
	return &AddObjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddObjectLogic) AddObject(req *types.ObjectForm) (resp *types.AddObjectReply, err error) {
	uuid, err := util.UUID()
	if err != nil {
		return
	}

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

	logx.Info(req)
	if req.TopKey == "" {

		err = errors.New("topKey is required oh")
		return
	}

	if req.Typ == 1 && req.TopKey != req.Key {
		err = errors.New("topKey is must equal to key")
		return
	}

	// TODO 增加校验

	// 防止重复添加的校验, 状态忽略，只看没有删除的
	isExistArgs := &model.ObjectIsExistArgs{
		Key:     req.Key,
		Typ:     req.Typ,
		SubType: req.SubType,
		TopKey:  req.TopKey,
	}

	exist, err := l.svcCtx.ObjectModel.IsExist(l.ctx, isExistArgs)
	if err != nil {
		return
	}

	if exist {
		err = errors.New("object already exists")
		return
	}

	object := &model.Object{
		Uuid:       uuid,
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
		TopKey:     req.TopKey,
	}

	logx.Info(object)
	_, err = l.svcCtx.ObjectModel.Insert(l.ctx, object)
	if err != nil {
		return
	}

	if req.Typ != 1 {
		summary := "添加菜单"
		if req.Typ == 2 {
			if req.SubType == 2 {
				summary = "添加菜单组"
			} else if req.SubType == 3 {
				summary = "添加隐藏菜单"
			}
		} else {
			if req.SubType == 1 {
				summary = "添加[GET]操作"
			} else if req.SubType == 2 {
				summary = "添加[POST]操作"
			} else if req.SubType == 3 {
				summary = "添加[PUT]操作"
			} else if req.SubType == 4 {
				summary = "添加[PAT]操作"
			} else if req.SubType == 5 {
				summary = "添加[DEL]操作"
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
			ObjectUuid: uuid,
			UserUuid:   userUUID,
			UserName:   userName,
			Type:       req.Typ,
			LogType:    1,
			LogSummary: logSummary,
			LogData:    string(logData),
		}

		_, err3 := l.svcCtx.ObjectLogModel.Insert(l.ctx, objectLogData)
		if err3 != nil {
			l.Logger.Error(err3)
		}
	}

	resp = &types.AddObjectReply{
		UUID: uuid,
	}

	return
}

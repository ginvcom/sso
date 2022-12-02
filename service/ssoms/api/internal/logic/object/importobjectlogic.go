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

type ImportObjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImportObjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImportObjectLogic {
	return &ImportObjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImportObjectLogic) ImportObject(req *types.ImportObjectReq) (resp *types.ImportObjectReply, err error) {
	userUUID := l.ctx.Value(config.UUID).(string)
	if userUUID == "" {
		err = errors.New("missing user info")
		resp = &types.ImportObjectReply{
			Status: "failed",
			Msg:    err.Error(),
		}
		err = nil
		return
	}

	userName := l.ctx.Value(config.Name).(string)
	if userName == "" {
		err = errors.New("missing user info")
		resp = &types.ImportObjectReply{
			Status: "failed",
			Msg:    err.Error(),
		}
		err = nil
		return
	}

	importObject := &model.Object{
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
		TopKey:     req.TopKey,
	}
	msgType := "菜单"
	object, err := l.svcCtx.ObjectModel.FindOneByUuid(l.ctx, req.UUID)
	if err != nil {
		if err == model.ErrNotFound {
			_, err = l.svcCtx.ObjectModel.Insert(l.ctx, object)
			if err != nil {
				resp = &types.ImportObjectReply{
					Status: "failed",
					Msg:    err.Error(),
				}
				err = nil
				return
			}

			if req.Typ != 1 {
				summary := "导入添加菜单"
				if req.Typ == 2 {
					if req.SubType == 2 {
						msgType = "菜单组"
						summary = "导入添加菜单组"
					} else if req.SubType == 3 {
						msgType = "隐藏菜单"
						summary = "导入添加隐藏菜单"
					}
				} else {
					if req.SubType == 1 {
						summary = "导入添加[GET]操作"
						msgType = "[GET]操作"
					} else if req.SubType == 2 {
						msgType = "[POST]操作"
						summary = "导入添加[POST]操作"
					} else if req.SubType == 3 {
						msgType = "[PUT]操作"
						summary = "导入添加[PUT]操作"
					} else if req.SubType == 4 {
						msgType = "[PAT]操作"
						summary = "导入添加[PAT]操作"
					} else if req.SubType == 5 {
						msgType = "[DEL]操作"
						summary = "导入添加[DEL]操作"
					}
				}
				logSummary := fmt.Sprintf("%s: %s", summary, req.ObjectName)
				logData, err2 := json.Marshal(*req)
				if err2 != nil {
					err = err2
					l.Logger.Error(err)
					resp = &types.ImportObjectReply{
						Status: "failed",
						Msg:    err.Error(),
					}
					err = nil
					return
				}

				objectLogData := &model.ObjectLog{
					ObjectUuid: req.UUID,
					UserUuid:   userUUID,
					UserName:   userName,
					Type:       req.Typ,
					LogType:    4,
					LogSummary: logSummary,
					LogData:    string(logData),
				}

				_, err3 := l.svcCtx.ObjectLogModel.Insert(l.ctx, objectLogData)
				if err3 != nil {
					l.Logger.Error(err3)
				}
			}

		}

		resp = &types.ImportObjectReply{
			Status: "success",
			Msg:    "新增" + msgType,
		}
		return
	}

	if object.Key == req.Key &&
		object.Type == req.Typ &&
		object.SubType == req.SubType &&
		object.Status == req.Status &&
		object.Icon == req.Icon &&
		object.Sort == req.Sort &&
		object.Identifier == req.Identifier &&
		object.Puuid == req.PUUID &&
		object.ObjectName == req.ObjectName {

		resp = &types.ImportObjectReply{
			Status: "ignore",
			Msg:    "存在相同记录",
		}
		return
	}
	// 执行更新操作
	err = l.svcCtx.ObjectModel.Update(l.ctx, importObject)
	if err != nil {
		return
	}

	if req.Typ != 1 {
		msgType = "菜单"
		summary := "导入更新菜单"
		if req.Typ == 2 {
			if req.SubType == 2 {
				msgType = "菜单组"
				summary = "导入更新菜单组"
			} else if req.SubType == 3 {
				msgType = "隐藏菜单"
				summary = "导入更新隐藏菜单"
			}
		} else {
			if req.SubType == 1 {
				msgType = "[GET]操作"
				summary = "导入更新[GET]操作"
			} else if req.SubType == 2 {
				msgType = "[POST]操作"
				summary = "导入更新[POST]操作"
			} else if req.SubType == 3 {
				msgType = "[PUT]操作"
				summary = "导入更新[PUT]操作"
			} else if req.SubType == 4 {
				msgType = "[PAT]操作"
				summary = "导入更新[PAT]操作"
			} else if req.SubType == 5 {
				msgType = "[DEL]操作"
				summary = "导入更新[DEL]操作"
			}
		}
		logSummary := fmt.Sprintf("%s: %s", summary, req.ObjectName)
		logData, err2 := json.Marshal(*req)
		if err2 != nil {
			err = err2
			l.Logger.Error(err)
			resp = &types.ImportObjectReply{
				Status: "failed",
				Msg:    err.Error(),
			}
			err = nil
			return
		}

		objectLogData := &model.ObjectLog{
			ObjectUuid: req.UUID,
			UserUuid:   userUUID,
			UserName:   userName,
			Type:       req.Typ,
			LogType:    5,
			LogSummary: logSummary,
			LogData:    string(logData),
		}

		_, err3 := l.svcCtx.ObjectLogModel.Insert(l.ctx, objectLogData)
		if err3 != nil {
			l.Logger.Error(err3)
		}
	}

	resp = &types.ImportObjectReply{
		Status: "success",
		Msg:    "更新" + msgType,
	}
	return
}

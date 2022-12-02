package logic

import (
	"context"
	"errors"
	"fmt"

	"sso/service/auth/api/internal/svc"
	"sso/service/auth/api/internal/types"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type MenusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenusLogic {
	return &MenusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenusLogic) Menus(req *types.SessionMenusReq) (resp *types.SessionMenusReply, err error) {
	if req.SystemCode == "" {
		err = errors.New("params SystemCode is required")
		return
	}

	l.Logger.Info(req.SystemCode)
	uuid, err := l.isPass(req)
	if err != nil {
		l.Logger.Errorf("authorization:%s, realRequestPath:%s", uuid)
		return
	}

	l.Logger.Info(uuid)

	roleUUIDArray, err := l.svcCtx.UserToRoleModel.FindRoleUUIDArrByUserUuid(l.ctx, uuid)
	if err != nil {
		return
	}

	l.Logger.Info(roleUUIDArray)

	listData, err := l.svcCtx.ObjectModel.FindMenusInRoleUUIDArray(l.ctx, req.SystemCode, roleUUIDArray)

	if err != nil {
		l.Logger.Error(err)
		return
	}

	resp = &types.SessionMenusReply{
		List: make([]*types.Menu, 0, 1),
	}
	logx.Info(listData)

	for _, obj := range listData {
		item := types.Menu{
			UUID: obj.Uuid,
			Meta: types.Meta{
				Name:       obj.ObjectName,
				Icon:       obj.Icon,
				Identifier: obj.Identifier,
				SubType:    obj.SubType,
				Parents:    make([]string, 0, 1),
			},
			Key:   obj.Key,
			PUUID: obj.Puuid,
		}
		resp.List = append(resp.List, &item)
	}
	newList := makeTree(resp.List, "")
	resp.List = newList
	logx.Info(newList)
	return
}

func makeTree(obs []*types.Menu, puuid string) []*types.Menu {
	tree := make([]*types.Menu, 0, 1)
	for _, child := range obs {
		if child.PUUID == puuid {
			item := child
			item.Children = makeTree(obs, item.UUID)
			if puuid != "" {
				item.Meta.Parents = append(item.Meta.Parents, puuid)
			}
			tree = append(tree, item)
		}
	}
	return tree
}

func (l *MenusLogic) isPass(req *types.SessionMenusReq) (uuid string, err error) {
	tok, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(l.svcCtx.Config.Auth.AccessSecret), nil
	})
	claims, ok := tok.Claims.(jwt.MapClaims) // 解析token中对内容
	if !ok {
		logx.Error("tokenExipreError")
		err = fmt.Errorf("tokenExipreError: tok.Claims is not ok ,tok.Claims: %+v , claims : %+v , ok:%v ", tok.Claims, claims, ok)
		return
	}
	uuid, _ = claims["uuid"].(string)
	if uuid == "" {
		logx.Error("tokenExipreError")
		err = fmt.Errorf("tokenExipreError: JwtAuthLogic.isPass invalid userId  tokRaw:%s , tokValid :%v", tok.Raw, tok.Valid)
	}

	return
}

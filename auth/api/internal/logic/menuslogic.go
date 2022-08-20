package logic

import (
	"context"
	"errors"
	"fmt"

	"sso/auth/api/internal/svc"
	"sso/auth/api/internal/types"

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
	logx.WithContext(l.ctx).Info(req.SystemCode)
	uuid, err := l.isPass(req)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("authorization:%s, realRequestPath:%s", uuid)
		return
	}

	logx.WithContext(l.ctx).Info(uuid)

	roleUUIDArray, err := l.svcCtx.UserToRoleModel.FindRoleUUIDArrByUserUuid(l.ctx, uuid)
	if err != nil {
		return
	}

	logx.WithContext(l.ctx).Info(roleUUIDArray)

	listData, err := l.svcCtx.ObjectModel.FindMenusInRoleUUIDArray(l.ctx, req.SystemCode, roleUUIDArray)

	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return
	}

	resp = &types.SessionMenusReply{
		List: make([]*types.Menu, 0, 1),
	}

	for _, obj := range *listData {
		item := types.Menu{
			UUID: obj.Uuid,
			Meta: types.Meta{
				Name:       obj.ObjectName,
				Icon:       obj.Icon,
				Identifier: obj.Identifier,
				SubType:    obj.SubType,
			},
			Key:   obj.Key,
			PUUID: obj.Puuid,
		}
		resp.List = append(resp.List, &item)
	}
	newList := makeTree(resp.List, "")
	resp.List = newList

	return
}

func makeTree(obs []*types.Menu, puuid string) []*types.Menu {
	var tree []*types.Menu
	for _, child := range obs {
		if child.PUUID == puuid {
			item := child
			item.Children = makeTree(obs, item.UUID)
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

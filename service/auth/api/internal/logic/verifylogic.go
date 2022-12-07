package logic

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"sso/service/auth/api/internal/svc"
	"sso/service/auth/api/internal/types"
	"sso/service/auth/model"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyLogic {
	return &VerifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyLogic) Verify(req *types.VerifyRequestReq) (resp *types.VerifyRequestReply, err error) {
	logx.Info("验证请求的账号参数", *req)

	// 获取uuid并且到redis校验是否过期(退出登录需要主动失效)
	uuid, name, err := l.isPass(req)
	if err != nil {
		// Authentication 身份验证错误
		err = fmt.Errorf("authentication error: %w", err)
		l.Logger.Error(err)
		return nil, err
	}

	// 注意身份验证报错是authentication error
	// 权限校验报错是authorization error
	// 明确错误结果，如果是授权对象不存在，告知对象不存在
	// 如果授权对象存在但是未授权，告知需要授权的对象名称
	if !l.urlNoAuth(req) {
		// 需要权限校验的页面.
		roleUUIDArr, err := l.svcCtx.UserToRoleModel.FindRoleUUIDArrByUserUuid(l.ctx, uuid)
		if err != nil {
			// authorization 授权查询错误，未授权
			err = fmt.Errorf("authorization error: %w", err)
			l.Logger.Error(err)
			return nil, err
		}
		var subType int64
		if req.Method == "GET" {
			subType = 1
		} else if req.Method == "POST" {
			subType = 2
		} else if req.Method == "PUT" {
			subType = 3
		} else if req.Method == "PATCH" {
			subType = 4
		} else if req.Method == "DELETE" {
			subType = 5
		}

		objectArgs := &model.ObjectFindOneArgs{
			Key:     req.URI,
			TopKey:  req.SystemCode,
			Typ:     3,
			SubType: subType,
		}
		objectRowBuilder := l.svcCtx.ObjectModel.RowBuilder()
		objects, err := l.svcCtx.ObjectModel.CustomFind(l.ctx, objectRowBuilder, objectArgs)
		if err != nil {
			if err == model.ErrNotFound {
				err = fmt.Errorf("no access path matches: \"%s %s\"", req.Method, req.URI)
				return nil, err
			}
			l.Logger.Error(err)
			return nil, err
		}
		object := &model.Object{}
		for _, o := range objects {
			ok, err2 := KeyMatch(req.URI, o.Key)
			if err2 != nil {
				l.Logger.Error(err2)
			} else {
				if ok {
					object = o
				}
			}
		}

		if object.Id == 0 {
			err = fmt.Errorf("no access path matches: \"%s %s\"", req.Method, req.URI)
			return nil, err
		}

		exits, err := l.svcCtx.ObjectModel.FindObjectInRoleUUIDArray(l.ctx, req.SystemCode, object.Key, subType, roleUUIDArr)
		if err != nil || !exits {
			err2 := fmt.Errorf("authorization error: %s: %w", object.ObjectName, err)
			l.Logger.Error(err2)
			err = fmt.Errorf("authorization error: %s", object.ObjectName)
			return nil, err
		}
	}

	resp = &types.VerifyRequestReply{
		UUID: uuid,
		Name: name,
	}

	return
}

// 当前url是否需要授权验证
func (l *VerifyLogic) urlNoAuth(req *types.VerifyRequestReq) bool {
	l.Logger.Info(l.svcCtx.Config.NoAuthUrls, req.MenuURI)
	menuNoAuthPaths, ok := l.svcCtx.Config.NoAuthUrls[req.MenuURI]
	if !ok {
		return false
	}

	for _, p := range menuNoAuthPaths {
		if req.Method == p.Method {
			ok, err := KeyMatch(req.URI, p.Path)
			if err != nil {
				l.Logger.Error(err)
				continue
			}
			if ok {
				return true
			}
		}
	}

	return false
}

func (l *VerifyLogic) isPass(req *types.VerifyRequestReq) (uuid, name string, err error) {
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
	name, _ = claims["name"].(string)
	if uuid == "" || name == "" {
		logx.Error("tokenExipreError")
		err = fmt.Errorf("tokenExipreError: JwtAuthLogic.isPass invalid userId  tokRaw:%s , tokValid :%v", tok.Raw, tok.Valid)
	}

	return
}

// KeyMatch2 determines whether key1 matches the pattern of key2 (similar to RESTful path), key2 can contain a *.
// For example, "/foo/bar" matches "/foo/*", "/resource1" matches "/:resource"
func KeyMatch(key1 string, key2 string) (bool, error) {
	key2 = strings.Replace(key2, "/*", "/.*", -1)

	re := regexp.MustCompile(`:[^/]+`)
	key2 = re.ReplaceAllString(key2, "$1[^/]+$2")

	res, err := regexp.MatchString(key2, key1)
	if err != nil {
		return false, err
	}
	return res, nil
}

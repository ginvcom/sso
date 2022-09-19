package logic

import (
	"context"
	"fmt"

	"sso/service/auth/api/internal/svc"
	"sso/service/auth/api/internal/types"

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
	logx.Info(req.Token, req.SystemCode, req.ServiceCode, req.URI)

	// 获取uuid并且到redis校验是否过期(退出登录需要主动失效)
	uuid, name, err := l.isPass(req)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("authorization:%s, realRequestPath:%s", uuid, name)
		return nil, err
	}

	if l.urlNoAuth(req.URI) == false {
		// 需要权限校验的页面.
	}

	resp = &types.VerifyRequestReply{
		UUID: uuid,
		Name: name,
	}

	return
}

// 当前url是否需要授权验证
func (l *VerifyLogic) urlNoAuth(path string) bool {
	_, ok := l.svcCtx.Config.NoAuthUrls[path]
	return ok
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

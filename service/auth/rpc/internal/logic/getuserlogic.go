package logic

import (
	"context"
	"fmt"

	"sso/service/auth/rpc/internal/svc"
	"sso/service/auth/rpc/types/auth"

	"github.com/ginvcom/util"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *auth.TokenReq) (*auth.UserInfoReply, error) {
	// 获取uuid并且到redis校验是否过期(退出登录需要主动失效)
	uuid, name, err := l.isPass(in)
	if err != nil {
		// Authentication 身份验证错误
		err = fmt.Errorf("authentication error: %w", err)
		l.Logger.Error(err)
		return nil, err
	}
	user, err := l.svcCtx.UserModel.FindOneByUuid(l.ctx, uuid)
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}

	resp := &auth.UserInfoReply{
		Uuid:   uuid,
		Name:   name,
		Mobile: user.Mobile,
		Gender: user.Gender,
		Avatar: user.Avatar,
		Birth:  util.TimeToString(user.Birth, "YYYY-MM-DD"),
	}

	return resp, nil
}

func (l *GetUserLogic) isPass(in *auth.TokenReq) (uuid, name string, err error) {
	tok, err := jwt.Parse(in.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(l.svcCtx.Config.JwtAuth.AccessSecret), nil
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

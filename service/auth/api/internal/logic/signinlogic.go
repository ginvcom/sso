package logic

import (
	"context"
	"errors"
	"fmt"
	"time"

	"sso/service/auth/api/internal/svc"
	"sso/service/auth/api/internal/types"
	"sso/service/auth/model"

	"github.com/ginvcom/util"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

type SignInLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignInLogic {
	return &SignInLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignInLogic) SignIn(req *types.SignInReq) (resp *types.SignInReply, err error) {
	userRowBuilder := l.svcCtx.UserModel.RowBuilder()
	user, err := l.svcCtx.UserModel.CustomFindOne(l.ctx, userRowBuilder, req.Mobile)
	if err != nil {
		if err == sqlc.ErrNotFound {
			err = errors.New("mobile not exits")
		}

		return
	}

	password := util.MD5(req.Password + user.Salt)

	if password != user.Password {
		err = errors.New("password error")
		return
	}

	objectArgs := &model.ObjectFindOneArgs{
		TopKey: req.SystemCode,
	}

	objectRowBuilder := l.svcCtx.ObjectModel.RowBuilder()
	object, err := l.svcCtx.ObjectModel.CustomFindOne(l.ctx, objectRowBuilder, objectArgs)
	if err != nil {
		l.Logger.Error(err)
		err = fmt.Errorf("system \"%s\" not exits", req.SystemCode)
		return
	}

	now := time.Now().Unix()
	seconds := l.svcCtx.Config.Auth.AccessExpire

	if req.Remember == "on" {
		seconds = l.svcCtx.Config.Auth.RememberAccessExpire
	}
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, seconds, user)
	if err != nil {
		return
	}

	resp = &types.SignInReply{
		Redirect:    object.Identifier,
		AccessToken: jwtToken,
		Name:        user.Name,
		Mobile:      user.Mobile,
		Avatar:      user.Avatar,
		Gender:      user.Gender,
		Expire:      seconds,
	}

	return
}

func (l *SignInLogic) getJwtToken(secretKey string, iat, seconds int64, user *model.User) (string, error) {
	// claims := make(jwt.MapClaims)
	// claims["exp"] = iat + seconds // jwt的过期时间，等于签发时间加上配置的过期时间
	// claims["iat"] = iat           // jwt的签发时间
	// claims["uuid"] = user.Uuid
	// claims["name"] = user.Name
	// token := jwt.New(jwt.SigningMethodHS256)
	// token.Claims = claims
	// return token.SignedString([]byte(secretKey))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":  iat + seconds,
		"iat":  iat,
		"uuid": user.Uuid,
		"name": user.Name,
	})

	return token.SignedString([]byte(secretKey))
}

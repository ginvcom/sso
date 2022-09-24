package handler

import (
	"errors"
	"net/http"
	"sso/service/gateway/api/internal/svc"
	"sso/service/gateway/api/internal/types"
	"text/template"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

func SignOutHandler(svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) (err error) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl, err := template.ParseFiles("./template/signOut.html")
	if err != nil {
		logx.Error(err)
		err = errors.New(svc.ErrParsingHtml)
		return
	}

	loginInfo := &types.LoginInfo{
		Redirect: "/?system=" + svcCtx.Meta.SystemCode,
	}

	// 销毁jwt
	// res, err := c.Pass(url)
	// if err != nil {
	// 	log.Println("pass error")
	// 	return
	// }

	// if err != nil {
	// 	loginInfo.Msg = "Internal Server Error"
	// 	tmpl.Execute(c.W, loginInfo)
	// 	return
	// }

	// 删除cookie
	expires := time.Unix(0, 0)
	tokenCookie := &http.Cookie{
		Name:    "ssoToken",
		Expires: expires,
	}
	userCookie := &http.Cookie{
		Name:    "ssoUser",
		Expires: expires,
	}
	http.SetCookie(w, tokenCookie)
	http.SetCookie(w, userCookie)

	tmpl.Execute(w, loginInfo)
	return
}

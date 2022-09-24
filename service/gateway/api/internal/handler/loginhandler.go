package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"sso/service/gateway/api/internal/svc"
	"sso/service/gateway/api/internal/types"
	"strings"
	"text/template"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

func LoginHandler(svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) (err error) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl, err := template.ParseFiles("./template/sign.html")
	if err != nil {
		logx.Error(err)
		err = errors.New(svc.ErrParsingHtml)
		return
	}

	loginInfo := &types.LoginInfo{
		SystemCode: svcCtx.Meta.SystemCode,
	}

	if r.Method == http.MethodGet {
		tmpl.Execute(w, loginInfo)
		return
	}

	body, err := svc.GetBody(svcCtx, r)
	if err != nil {
		loginInfo.Msg = "Internal Server Error"
		tmpl.Execute(w, loginInfo)
		return
	}

	err = json.Unmarshal(body, loginInfo)
	if err != nil {
		loginInfo.Msg = "Internal Server Error"
		tmpl.Execute(w, loginInfo)
		return
	}

	if loginInfo.Mobile == "" {
		loginInfo.MobileMsg = "mobile is reqiured"
	}

	if loginInfo.Password == "" {
		loginInfo.PasswordMsg = "passord is reqiured"
	}

	if loginInfo.MobileMsg != "" || loginInfo.PasswordMsg != "" {
		tmpl.Execute(w, loginInfo)
		return
	}

	status, res, err := svcCtx.Request(body)

	if err != nil {
		loginInfo.Msg = "Internal Server Error"
		tmpl.Execute(w, loginInfo)
		return
	}
	if status != http.StatusOK {
		loginInfo.Msg = http.StatusText(status)
		tmpl.Execute(w, loginInfo)
		return
	}

	var resp types.LoginResponse
	err = json.Unmarshal(res, &resp)

	if err != nil {
		loginInfo.Msg = "Response Data Error"
		tmpl.Execute(w, loginInfo)
		return
	}

	if resp.Code == 0 {
		u := types.UserBasic{
			Name:   resp.Data.Name,
			Avatar: resp.Data.Avatar,
			Gender: resp.Data.Gender,
		}
		user, err2 := json.Marshal(u)
		if err2 != nil {
			loginInfo.Msg = "Response Data Format Error"
			tmpl.Execute(w, loginInfo)
			return
		}

		count := len(loginInfo.Password)
		loginInfo.Password = strings.Repeat("*", count)
		loginInfo.Redirect = resp.Data.Redirect
		expires := time.Now().Add(6 * time.Hour).Local()
		if resp.Data.Expire > 0 {
			expires = time.Now().Add(time.Duration(resp.Data.Expire) * time.Second).Local()
		}
		tokenCookie := &http.Cookie{
			Name:     "ssoToken",
			Value:    resp.Data.AccessToken,
			Expires:  expires,
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
			// Secure:   true,
		}
		http.SetCookie(w, tokenCookie)

		userCookie := &http.Cookie{
			Name:   "ssoUser",
			Domain: svcCtx.Config.UserCookieDomin,
			// 需要使用url.QueryEscape转义一下，不然中文等特殊字符部分不能存入cokie
			Value:   url.QueryEscape(string(user)),
			Expires: expires,
			// Secure:   true,
			SameSite: http.SameSiteStrictMode,
		}
		http.SetCookie(w, userCookie)
	} else {
		loginInfo.Msg = resp.Msg
	}

	tmpl.Execute(w, loginInfo)
	return
}

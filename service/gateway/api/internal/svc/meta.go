package svc

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	ActionProxy              int = iota
	ActionSignIn                 // 登录
	ActionSignOut                // 退出登录
	ActionVerify                 // 验证
	ActionSessionMenus           // 获取菜单
	ActionSessionMenuActions     // 获取菜单下的操作
	ActionImage                  // 获取图片
)

type Meta struct {
	Destination string
	Action      int
	SystemCode  string
	ServiceCode string
	URI         string // 请求地址
	Token       string
	UUID        string // 登录后通过jwt拿到的uuid
	Name        string // 登录后通过jwt拿到的用户名
	Method      string
}

func GetMeta(r *http.Request) (c Meta, err error) {
	c = Meta{}
	destination := r.Header.Get("Accept")
	c.SystemCode = r.Header.Get("x-client-system")
	c.ServiceCode = r.Header.Get("x-client-service")
	c.URI = r.URL.RequestURI()
	fmt.Println("c.URI", c.URI)
	if strings.Contains(destination, "text/html") {
		c.SystemCode = r.URL.Query().Get("system")
		if c.SystemCode == "" {
			c.SystemCode = "ssoms"
		}
		if r.URL.Path == "/sign-out" {
			c.ServiceCode = "auth"
			c.URI = "/sign-out"
			c.Action = ActionSignOut
		} else {
			c.ServiceCode = "auth"
			c.URI = "/sign-in"
			c.Action = ActionSignIn
		}
	} else if strings.Contains(destination, "image/*") {
		c.Action = ActionImage
	} else {
		authorizationCookie, err2 := r.Cookie("ssoToken")
		if err2 != nil {
			logx.Error(err2)
			err = err2
			return
		}
		c.Token = authorizationCookie.Value
		if c.ServiceCode == "auth" {
			if c.URI == "/session-menus" {
				c.Action = ActionSessionMenus
			} else if c.URI == "/session-menu-actions" {
				c.Action = ActionSessionMenuActions
			}
		} else {
			c.Action = ActionProxy
		}
	}
	c.Method = r.Method
	return
}

package svc

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

func (svcCtx *ServiceContext) Request(body []byte) (status int, resp []byte, err error) {
	resp = []byte("")
	newBody := bytes.NewBuffer(body)
	// requestURL := svcCtx.Meta.URI
	serviceAddress, ok := svcCtx.Config.Upstreams[svcCtx.Meta.ServiceCode]
	if !ok {
		err = errors.New(ErrServiceNotExits + ", serviceCode: " + svcCtx.Meta.ServiceCode)
		return
	}

	requestURL := serviceAddress + svcCtx.Meta.URI

	// 获取菜单和获取页面操作对象的接口为了适配GET请求类型，采用地址传参
	if svcCtx.Meta.Action == ActionSessionMenus || svcCtx.Meta.Action == ActionSessionMenuActions {
		requestURL = fmt.Sprintf("%s%s?token=%s&systemCode=%s", serviceAddress, svcCtx.Meta.URI, svcCtx.Meta.Token, svcCtx.Meta.SystemCode)
	}

	// 记录转发日志，对登录密码进行模糊处理
	bodyStr := string(body)
	if svcCtx.Meta.Action == ActionSignIn {
		re, err2 := regexp.Compile(`"password":"([^"]*)"`)
		if err2 == nil {
			subArr := re.FindAllStringSubmatch(bodyStr, -1)
			for _, sub := range subArr {
				starStr := strings.Repeat("*", len(sub[1]))
				newStr := fmt.Sprintf(`"password":"%s"`, starStr)
				bodyStr = strings.Replace(bodyStr, sub[0], newStr, 1)
			}
		}
	}
	logx.Infof("proxy request: serviceCode: %s,\nurl: [%s]%s,\nbody: %s", svcCtx.Meta.ServiceCode, svcCtx.Meta.Method, requestURL, bodyStr)

	req, err := http.NewRequest(svcCtx.Meta.Method, requestURL, newBody)
	if err != nil {
		logx.Info(requestURL, err)
		return
	}

	if svcCtx.Meta.Method != "GET" {
		req.Header.Set("Content-Type", "application/json;charset=utf-8")
	}

	req.Header.Set("x-origin-uri", svcCtx.Meta.URI)
	req.Header.Set("x-origin-uuid", svcCtx.Meta.UUID)
	req.Header.Set("x-origin-name", svcCtx.Meta.Name)
	req.Header.Set("Authorization", svcCtx.Meta.Token)
	// req = req.WithContext(ctx)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logx.Info(requestURL, err)
		err = fmt.Errorf(ErrServiceNotResponding)
		return
	}

	status = res.StatusCode

	if status != http.StatusOK {
		return
	}

	resp, err = io.ReadAll(res.Body)
	return
}

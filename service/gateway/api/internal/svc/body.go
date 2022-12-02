package svc

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"sso/service/gateway/api/internal/types"
)

func GetBody(c *ServiceContext, r *http.Request) (body []byte, err error) {
	if c.Meta.Action == ActionSignIn {
		req := types.LoginReq{
			SystemCode: c.Meta.SystemCode,
			Mobile:     r.PostFormValue("mobile"),
			Password:   r.PostFormValue("password"),
			Remember:   r.PostFormValue("remember"),
		}
		body, err = json.Marshal(req)
		if err != nil {
			err = errors.New("request params error")
			return
		}
	} else if c.Meta.Action == ActionVerify {
		referer := r.Header.Get("Referer")
		refererURL, err2 := url.Parse(referer)
		if err2 != nil {
			err = err2
			return
		}
		req := types.VerifyReq{
			SystemCode:  c.Meta.SystemCode,
			ServiceCode: c.Meta.ServiceCode,
			Token:       c.Meta.Token,
			URI:         c.Meta.URI,
			Method:      r.Method,
			MenuURI:     refererURL.Path,
		}
		body, err = json.Marshal(req)
		if err != nil {
			err = errors.New("request params error")
			return
		}
	} else {
		body, err = io.ReadAll(r.Body)
		if err != nil {
			err = errors.New("read body error")
			return
		}
	}

	return
}

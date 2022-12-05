package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sso/service/gateway/api/internal/svc"
	"sso/service/gateway/api/internal/types"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

func VerifyHandler(svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) (err error) {
	meta := svcCtx.Meta
	meta.Action = svc.ActionVerify
	meta.Method = http.MethodPost
	meta.ServiceCode = "auth"
	meta.URI = r.URL.Path
	authCtx := &svc.ServiceContext{
		Config: svcCtx.Config,
		Meta:   meta,
	}
	body, err := svc.GetBody(authCtx, r)
	if err != nil {
		logx.Info(err)
		res := types.VerifyResponse{
			Code: -9999,
			Msg:  err.Error(),
		}
		resStr, _ := json.Marshal(res)
		w.Write(resStr)
		return
	}
	authCtx.Meta.URI = "/verify-request"
	status, resp, err := authCtx.Request(body)
	if err != nil {
		logx.Info(err)
		res := types.VerifyResponse{
			Code: -9999,
			Msg:  err.Error(),
		}
		resStr, _ := json.Marshal(res)
		w.Write(resStr)
		return
	}

	if status != http.StatusOK {
		w.WriteHeader(status)
		w.Write([]byte(http.StatusText(status)))
		err = fmt.Errorf("verify request failed: %s(%d)", http.StatusText(status), status)
		return
	}

	var realResp types.VerifyResponse
	err = json.Unmarshal(resp, &realResp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if realResp.Code != 0 {
		err = fmt.Errorf("verify request error, code: %d, msg: %v", realResp.Code, realResp.Msg)
		if strings.Contains(realResp.Msg, "authentication error") {
			w.WriteHeader(http.StatusUnauthorized)
		} else if realResp.Msg == "object not exits" {
			err = fmt.Errorf("object not exits: %s %s", meta.Method, meta.URI)
		}
		w.Write(resp)
		return
	}

	// 鉴权通过, 附传用户uuid和用户名请求接口
	svcCtx.Meta.UUID = realResp.Data.UUID
	svcCtx.Meta.Name = realResp.Data.Name
	return nil
}

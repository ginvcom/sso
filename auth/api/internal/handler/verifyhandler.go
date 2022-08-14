package handler

import (
	"net/http"

	"sso/auth/api/internal/logic"
	"sso/auth/api/internal/svc"
	"sso/auth/api/internal/types"

	"github.com/ginvcom/util"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func verifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VerifyRequestReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewVerifyLogic(r.Context(), svcCtx)
		resp, err := l.Verify(&req)
		util.Response(w, resp, err)
	}
}

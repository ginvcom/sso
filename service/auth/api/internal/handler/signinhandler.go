package handler

import (
	"net/http"

	"github.com/ginvcom/util"
	"github.com/zeromicro/go-zero/rest/httpx"
	"sso/service/auth/api/internal/logic"
	"sso/service/auth/api/internal/svc"
	"sso/service/auth/api/internal/types"
)

func signInHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SignInReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSignInLogic(r.Context(), svcCtx)
		resp, err := l.SignIn(&req)

		util.Response(w, resp, err)
	}
}

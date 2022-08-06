package auth

import (
	"net/http"

	"sso/auth/api/internal/logic/auth"
	"sso/auth/api/internal/svc"
	"sso/auth/api/internal/types"

	"github.com/ginvcom/util"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SignInHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SignInReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := auth.NewSignInLogic(r.Context(), svcCtx)
		resp, err := l.SignIn(&req)
		util.Response(w, resp, err)
	}
}

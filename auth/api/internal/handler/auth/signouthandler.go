package auth

import (
	"net/http"

	"sso/auth/api/internal/logic/auth"
	"sso/auth/api/internal/svc"

	"github.com/ginvcom/util"
)

func SignOutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := auth.NewSignOutLogic(r.Context(), svcCtx)
		resp, err := l.SignOut(r)
		util.Response(w, resp, err)
	}
}

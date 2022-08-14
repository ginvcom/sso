package handler

import (
	"net/http"

	"sso/auth/api/internal/logic"
	"sso/auth/api/internal/svc"

	"github.com/ginvcom/util"
)

func signOutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewSignOutLogic(r.Context(), svcCtx)
		resp, err := l.SignOut()
		util.Response(w, resp, err)
	}
}

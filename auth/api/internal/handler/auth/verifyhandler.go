package auth

import (
	"net/http"

	"sso/auth/api/internal/logic/auth"
	"sso/auth/api/internal/svc"

	"github.com/ginvcom/util"
)

func VerifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := auth.NewVerifyLogic(r.Context(), svcCtx)
		resp, err := l.Verify(r)
		util.Response(w, resp, err)
	}
}

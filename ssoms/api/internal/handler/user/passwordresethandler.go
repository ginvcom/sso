package user

import (
	"net/http"

	"sso/ssoms/api/internal/logic/user"
	"sso/ssoms/api/internal/svc"

	"github.com/ginvcom/util"
)

func PasswordResetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewPasswordResetLogic(r.Context(), svcCtx)
		resp, err := l.PasswordReset()
		util.Response(w, resp, err)
	}
}

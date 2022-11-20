package user

import (
	"net/http"

	"sso/service/ssoms/api/internal/logic/user"
	"sso/service/ssoms/api/internal/svc"

	"github.com/ginvcom/util"
)

func ProfileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewProfileLogic(r.Context(), svcCtx)
		resp, err := l.Profile()

		util.Response(w, resp, err)
	}
}

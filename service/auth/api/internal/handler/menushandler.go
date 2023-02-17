package handler

import (
	"net/http"

	"sso/service/auth/api/internal/logic"
	"sso/service/auth/api/internal/svc"

	"github.com/ginvcom/util"
)

func menusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewMenusLogic(r.Context(), svcCtx)
		resp, err := l.Menus()

		util.Response(w, resp, err)
	}
}

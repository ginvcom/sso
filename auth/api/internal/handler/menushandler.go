package handler

import (
	"net/http"

	"github.com/ginvcom/util"
	"github.com/zeromicro/go-zero/rest/httpx"
	"sso/auth/api/internal/logic"
	"sso/auth/api/internal/svc"
	"sso/auth/api/internal/types"
)

func menusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SessionMenusReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewMenusLogic(r.Context(), svcCtx)
		resp, err := l.Menus(&req)
		util.Response(w, resp, err)
	}
}

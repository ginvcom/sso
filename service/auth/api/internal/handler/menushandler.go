package handler

import (
	"net/http"

	"sso/service/auth/api/internal/logic"
	"sso/service/auth/api/internal/svc"
	"sso/service/auth/api/internal/types"

	"github.com/ginvcom/util"
	"github.com/zeromicro/go-zero/rest/httpx"
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
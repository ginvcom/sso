package role

import (
	"net/http"
	"sso/util"

	"github.com/zeromicro/go-zero/rest/httpx"
	"sso/ssoms/api/internal/logic/role"
	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"
)

func InheritancesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InheritancesReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := role.NewInheritancesLogic(r.Context(), svcCtx)
		resp, err := l.Inheritances(&req)
		util.Response(w, resp, err)
	}
}

package role

import (
	"net/http"
	"sso/util"

	"sso/ssoms/api/internal/logic/role"
	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddRoleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := role.NewAddRoleLogic(r.Context(), svcCtx)
		resp, err := l.AddRole(&req)
		util.Response(w, resp, err)
	}
}

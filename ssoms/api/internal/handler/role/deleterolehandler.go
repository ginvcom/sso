package role

import (
	"net/http"
	"sso/util"

	"github.com/zeromicro/go-zero/rest/httpx"
	"sso/ssoms/api/internal/logic/role"
	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"
)

func DeleteRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteRoleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := role.NewDeleteRoleLogic(r.Context(), svcCtx)
		resp, err := l.DeleteRole(&req)
		util.Response(w, resp, err)
	}
}

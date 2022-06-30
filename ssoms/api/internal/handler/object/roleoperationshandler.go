package object

import (
	"net/http"
	"sso/util"

	"github.com/zeromicro/go-zero/rest/httpx"
	"sso/ssoms/api/internal/logic/object"
	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"
)

func RoleOperationsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleOperationsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := object.NewRoleOperationsLogic(r.Context(), svcCtx)
		resp, err := l.RoleOperations(&req)
		util.Response(w, resp, err)
	}
}

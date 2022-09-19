package object

import (
	"net/http"

	"github.com/ginvcom/util"

	"sso/service/ssoms/api/internal/logic/object"
	"sso/service/ssoms/api/internal/svc"
	"sso/service/ssoms/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
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

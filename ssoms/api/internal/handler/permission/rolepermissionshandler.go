package permission

import (
	"net/http"

	"github.com/ginvcom/util"

	"sso/ssoms/api/internal/logic/permission"
	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RolePermissionsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RolePermissionsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := permission.NewRolePermissionsLogic(r.Context(), svcCtx)
		resp, err := l.RolePermissions(&req)
		util.Response(w, resp, err)
	}
}

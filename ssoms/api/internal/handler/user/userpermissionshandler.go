package user

import (
	"net/http"
	"sso/util"

	"github.com/zeromicro/go-zero/rest/httpx"
	"sso/ssoms/api/internal/logic/user"
	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"
)

func UserPermissionsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserPermissionsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUserPermissionsLogic(r.Context(), svcCtx)
		resp, err := l.UserPermissions(&req)
		util.Response(w, resp, err)
	}
}
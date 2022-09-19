package role

import (
	"net/http"

	"github.com/ginvcom/util"

	"sso/service/ssoms/api/internal/logic/role"
	"sso/service/ssoms/api/internal/svc"
	"sso/service/ssoms/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AssignedUsersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AssignedUsersReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := role.NewAssignedUsersLogic(r.Context(), svcCtx)
		resp, err := l.AssignedUsers(&req)
		util.Response(w, resp, err)
	}
}

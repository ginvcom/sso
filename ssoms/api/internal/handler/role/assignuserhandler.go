package role

import (
	"net/http"

	"github.com/ginvcom/util"

	"sso/ssoms/api/internal/logic/role"
	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AssignUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AssignUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := role.NewAssignUserLogic(r.Context(), svcCtx)
		resp, err := l.AssignUser(&req)
		util.Response(w, resp, err)
	}
}

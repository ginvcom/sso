package role

import (
	"net/http"
	"sso/util"

	"github.com/zeromicro/go-zero/rest/httpx"
	"sso/ssoms/api/internal/logic/role"
	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"
)

func ExtendRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddInheritanceReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := role.NewExtendRoleLogic(r.Context(), svcCtx)
		resp, err := l.ExtendRole(&req)
		util.Response(w, resp, err)
	}
}

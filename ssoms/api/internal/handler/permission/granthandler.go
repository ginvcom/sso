package permission

import (
	"net/http"
	"sso/util"

	"github.com/zeromicro/go-zero/rest/httpx"
	"sso/ssoms/api/internal/logic/permission"
	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"
)

func GrantHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GrantReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := permission.NewGrantLogic(r.Context(), svcCtx)
		resp, err := l.Grant(&req)
		util.Response(w, resp, err)
	}
}

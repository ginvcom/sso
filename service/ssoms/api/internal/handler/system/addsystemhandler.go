package system

import (
	"net/http"

	"github.com/ginvcom/util"
	"github.com/zeromicro/go-zero/rest/httpx"
	"sso/service/ssoms/api/internal/logic/system"
	"sso/service/ssoms/api/internal/svc"
	"sso/service/ssoms/api/internal/types"
)

func AddSystemHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SystemForm
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := system.NewAddSystemLogic(r.Context(), svcCtx)
		resp, err := l.AddSystem(&req)

		util.Response(w, resp, err)
	}
}

package object

import (
	"net/http"
	"sso/util"

	"github.com/zeromicro/go-zero/rest/httpx"
	"sso/ssoms/api/internal/logic/object"
	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"
)

func AddObjectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddObjectReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := object.NewAddObjectLogic(r.Context(), svcCtx)
		resp, err := l.AddObject(&req)
		util.Response(w, resp, err)
	}
}

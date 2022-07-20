package object

import (
	"net/http"

	"github.com/ginvcom/util"

	"sso/ssoms/api/internal/logic/object"
	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddObjectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ObjectForm
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := object.NewAddObjectLogic(r.Context(), svcCtx)
		resp, err := l.AddObject(&req)
		util.Response(w, resp, err)
	}
}

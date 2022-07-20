package object

import (
	"net/http"

	"github.com/ginvcom/util"

	"sso/ssoms/api/internal/logic/object"
	"sso/ssoms/api/internal/svc"
	"sso/ssoms/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteObjectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteObjectReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := object.NewDeleteObjectLogic(r.Context(), svcCtx)
		resp, err := l.DeleteObject(&req)
		util.Response(w, resp, err)
	}
}

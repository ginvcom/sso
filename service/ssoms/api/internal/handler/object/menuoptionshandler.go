package object

import (
	"net/http"

	"sso/service/ssoms/api/internal/logic/object"
	"sso/service/ssoms/api/internal/svc"
	"sso/service/ssoms/api/internal/types"

	"github.com/ginvcom/util"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func MenuOptionsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MenuOptionsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := object.NewMenuOptionsLogic(r.Context(), svcCtx)
		resp, err := l.MenuOptions(&req)
		util.Response(w, resp, err)
	}
}

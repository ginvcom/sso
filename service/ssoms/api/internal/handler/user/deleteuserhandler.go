package user

import (
	"net/http"

	"github.com/ginvcom/util"
	"github.com/zeromicro/go-zero/rest/httpx"
	"sso/service/ssoms/api/internal/logic/user"
	"sso/service/ssoms/api/internal/svc"
	"sso/service/ssoms/api/internal/types"
)

func DeleteUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewDeleteUserLogic(r.Context(), svcCtx)
		resp, err := l.DeleteUser(&req)

		util.Response(w, resp, err)
	}
}

package user

import (
	"net/http"

	"sso/service/ssoms/api/internal/logic/user"
	"sso/service/ssoms/api/internal/svc"
	"sso/service/ssoms/api/internal/types"

	"github.com/ginvcom/util"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PasswordResetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PasswordResetReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewPasswordResetLogic(r.Context(), svcCtx)
		resp, err := l.PasswordReset(&req)
		util.Response(w, resp, err)
	}
}

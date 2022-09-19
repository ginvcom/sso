package user

import (
	"net/http"

	"sso/service/ssoms/api/internal/logic/user"
	"sso/service/ssoms/api/internal/svc"
	"sso/service/ssoms/api/internal/types"

	"github.com/ginvcom/util"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AvatarUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AvatarUploadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewAvatarUploadLogic(r.Context(), svcCtx)
		resp, err := l.AvatarUpload(&req)
		util.Response(w, resp, err)
	}
}

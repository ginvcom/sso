package role

import (
	"net/http"

	"sso/service/ssoms/api/internal/logic/role"
	"sso/service/ssoms/api/internal/svc"

	"github.com/ginvcom/util"
)

func RoleOptionsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := role.NewRoleOptionsLogic(r.Context(), svcCtx)
		resp, err := l.RoleOptions()

		util.Response(w, resp, err)
	}
}

package role

import (
	"net/http"

	"github.com/ginvcom/util"

	"sso/ssoms/api/internal/logic/role"
	"sso/ssoms/api/internal/svc"
)

func RoleOptionsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := role.NewRoleOptionsLogic(r.Context(), svcCtx)
		resp, err := l.RoleOptions()
		util.Response(w, resp, err)
	}
}

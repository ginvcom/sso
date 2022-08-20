package home

import (
	"net/http"

	"sso/ssoms/api/internal/logic/home"
	"sso/ssoms/api/internal/svc"

	"github.com/ginvcom/util"
)

func HomeStatisticHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := home.NewHomeStatisticLogic(r.Context(), svcCtx)
		resp, err := l.HomeStatistic()
		util.Response(w, resp, err)
	}
}

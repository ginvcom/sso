package handler

import (
	"io"
	"net/http"
	"os"
	"sso/service/gateway/api/internal/svc"
)

func ImageHandler(svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) (err error) {
	w.Header().Set("Content-Type", "image/png")
	file, err := os.Open("./image/" + r.URL.Path)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	defer file.Close()

	buffer, err := io.ReadAll(file)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Write(buffer)

	return
}

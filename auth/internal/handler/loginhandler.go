package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hpf/auth/internal/logic"
	"hpf/auth/internal/svc"
	"hpf/auth/internal/types"
	"hpf/common/response"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			response.HttpResult(r, w, resp, err)
		} else {
			response.HttpResult(r, w, resp, nil)
		}
	}
}

package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hpf/common/response"
	"hpf/service/auth/internal/logic"
	"hpf/service/auth/internal/svc"
	"hpf/service/auth/internal/types"
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

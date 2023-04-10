package handler

import (
	"hpf/common/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hpf/service/card/api/internal/logic"
	"hpf/service/card/api/internal/svc"
	"hpf/service/card/api/internal/types"
)

func listHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewListLogic(r.Context(), svcCtx)
		resp, err := l.List(&req, r)
		if err != nil {
			response.HttpResult(r, w, resp, err)
		} else {
			response.HttpResult(r, w, resp, nil)
		}
	}
}

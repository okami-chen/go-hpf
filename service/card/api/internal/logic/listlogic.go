package logic

import (
	"context"
	"hpf/service/card/api/internal/svc"
	"hpf/service/card/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.SearchRequest, request *http.Request) (resp *types.SearchResponse, err error) {
	l.Logger.Debug(request.Header.Get("Authorization"))
	return &types.SearchResponse{UserId: "2"}, nil
}

package logic

import (
	"context"
	"hpf/service/card/api/internal/entity"
	"hpf/service/card/api/internal/services"
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

func (l *ListLogic) List(req *types.SearchRequest, request *http.Request) (resp *entity.Card, err error) {
	service := services.NewCardServiceImpl(l.svcCtx.DbEngin)
	ret := service.FindOne(int64(1))
	return ret, nil
}

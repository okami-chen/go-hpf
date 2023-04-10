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
	l.Logger.Debug(request.Header.Get("Authorization"))
	dao := &services.CardModelImpl{
		Table: "sec_card",
		Db:    l.svcCtx.DbEngin,
	}
	ret := dao.FindOne(int64(1))
	return ret, nil
}

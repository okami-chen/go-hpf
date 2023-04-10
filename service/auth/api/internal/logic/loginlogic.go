package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-module/carbon/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"hpf/service/auth/api/internal/svc"
	"hpf/service/auth/api/internal/types"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	time := carbon.Now()
	t, _ := l.getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, time.Timestamp(), 86400, 1)
	response := &types.LoginResponse{Token: t, UserId: 1, Page: "/pages/card/index/index"}
	return response, nil
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["expire_at"] = iat + seconds
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

package middleware

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type ExampleMiddleware struct {
}

func NewExampleMiddleware() *ExampleMiddleware {
	return &ExampleMiddleware{}
}

func (m *ExampleMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("example middle")
		//w.WriteHeader(http.StatusUnauthorized)
		//response.JwtUnauthorizedResult(w, r, nil)
		next(w, r)
	}
}

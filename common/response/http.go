package response

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/status"
	"hpf/common/errorx"
	"net/http"
)

type TResponse struct {
	TraceId string `json:"trace_id"`
	SpanId  string `json:"span_id"`
}

type HttpCustomResponse struct {
	Success bool        `json:"success"`
	Code    uint64      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Trace   interface{} `json:"trace,omitempty"`
}

// http返回
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {

	if err == nil {
		//成功返回
		resp := &HttpCustomResponse{
			Success: true,
			Code:    0,
			Message: "",
			Data:    resp,
		}
		logx.WithContext(r.Context()).Infow("http Response", logx.Field("response", resp))
		resp.Trace = TResponse{
			TraceId: traceIdFromContext(r.Context()),
			SpanId:  spanIdFromContext(r.Context()),
		}
		httpx.WriteJson(w, http.StatusOK, resp)
	} else {
		//错误返回
		errcode := uint64(500)
		errmsg := "服务器错误"

		causeErr := errors.Cause(err)                  // err类型
		if e, ok := causeErr.(*errorx.CodeError); ok { //自定义错误类型
			//自定义CodeError
			errcode = e.GetErrCode()
			errmsg = e.GetErrMsg()
		} else {
			if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
				grpcCode := uint64(gstatus.Code())
				errcode = grpcCode
				errmsg = gstatus.Message()
			}
		}

		resp := &HttpCustomResponse{
			Success: false,
			Code:    errcode,
			Message: errmsg,
			Data:    nil,
		}

		logx.WithContext(r.Context()).Errorw("system error",
			logx.Field("error_code", errcode),
			logx.Field("error_message", errmsg),
			logx.Field("response", resp),
		)

		resp.Trace = TResponse{
			TraceId: traceIdFromContext(r.Context()),
			SpanId:  spanIdFromContext(r.Context()),
		}

		httpx.WriteJson(w, http.StatusOK, resp)
	}
}

func JwtUnauthorizedResult(w http.ResponseWriter, r *http.Request, err error) {
	httpx.WriteJson(w, http.StatusUnauthorized, &HttpCustomResponse{
		Success: false,
		Code:    401,
		Message: "鉴权失败",
		Data:    nil,
		Trace: TResponse{
			TraceId: traceIdFromContext(r.Context()),
			SpanId:  spanIdFromContext(r.Context()),
		}})
}

func spanIdFromContext(ctx context.Context) string {
	spanCtx := trace.SpanContextFromContext(ctx)
	if spanCtx.HasSpanID() {
		return spanCtx.SpanID().String()
	}

	return ""
}

func traceIdFromContext(ctx context.Context) string {
	spanCtx := trace.SpanContextFromContext(ctx)
	if spanCtx.HasTraceID() {
		return spanCtx.TraceID().String()
	}
	return ""
}

package errorx

import (
	"context"
	"github.com/golang-module/carbon/v2"
	"go.opentelemetry.io/otel/trace"
)

const defaultCode = 1001

type CustomCodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CustomCodeErrorResponse struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Trace   interface{} `json:"trace,omitempty"`
}

type CustomTraceResponse struct {
	Time    string `json:"time"`
	TraceId string `json:"trace"`
	SpanId  string `json:"span"`
}

func NewCodeError(code int, msg string) error {
	return &CustomCodeError{Code: code, Msg: msg}
}

func NewCustomError(msg string) error {
	return NewCodeError(defaultCode, msg)
}

func (e *CustomCodeError) Error() string {
	return e.Msg
}

func (e *CustomCodeError) Data(r context.Context) *CustomCodeErrorResponse {
	return &CustomCodeErrorResponse{
		Success: false,
		Code:    e.Code,
		Message: e.Msg,
		Data:    nil,
		Trace: &CustomTraceResponse{
			Time:    carbon.Now().ToDateTimeString(),
			TraceId: traceIdFromContext(r),
			SpanId:  spanIdFromContext(r),
		},
	}
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

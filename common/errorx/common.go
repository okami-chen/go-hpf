package errorx

import (
	"fmt"
)

type CodeError struct {
	errCode uint64
	errMsg  string
}

func (e *CodeError) GetErrCode() uint64 {
	return e.errCode
}

func (e *CodeError) GetErrMsg() string {
	return e.errMsg
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("{\"ErrCode\":%d,\"ErrMsg\":\"%s\"}", e.errCode, e.errMsg)
}

func NewErrCodeMsg(errCode uint64, errMsg string) *CodeError {
	return &CodeError{errCode: errCode, errMsg: errMsg}
}

package errorx

const defaultCode = 1001

type CustomCodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CustomCodeErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
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

func (e *CustomCodeError) Data() *CustomCodeErrorResponse {
	return &CustomCodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}

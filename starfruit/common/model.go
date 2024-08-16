package common

type BusinessCode int

// Result 成功返回
type Result struct {
	Code BusinessCode `json:"code"`
	Msg  string       `json:"msg"`
	Data any          `json:"data"`
}

// Success 成功
func (r *Result) Success(data any) *Result {
	r.Code = 200
	r.Msg = "success"
	r.Data = data
	return r
}

// Fail 失败返回
func (r *Result) Fail(code BusinessCode, msg string) *Result {
	r.Code = code
	r.Msg = msg
	return r
}

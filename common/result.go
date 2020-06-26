package common

// 返回结果结构体
type Result struct {
	OK   bool        `json:"ok"`
	MSG  string      `json:"msg"`
	DATA interface{} `json:"data,omitempty"`
}

func NewResult(OK bool, MSG string, DATA interface{}) *Result {
	return &Result{OK: OK, MSG: MSG, DATA: DATA}
}

func Ok() *Result {
	return &Result{OK: true}
}

func Success(DATA interface{}) *Result {
	return &Result{OK: true, MSG: "Success", DATA: DATA}
}

func Fail(MSG string) *Result {
	return &Result{OK: false, MSG: MSG}
}

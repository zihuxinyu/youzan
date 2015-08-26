package youzan

import "fmt"

const (
	ErrCodeOK = 0
	ErrCodeAccessTokenExpired = 42001 // access_token 过期(无效)返回这个错误
	ErrCodeSuiteAccessTokenExpired = 42009 // suite_access_token 过期(无效)返回这个错误
)

type Error struct {
 	ErrorResponse struct {
		              Code int `json:"code"`
		              Msg  string `json:"msg"`
	              } `json:"error_response"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("errcode: %d, errmsg: %s", e.ErrorResponse.Code, e.ErrorResponse.Msg)
}

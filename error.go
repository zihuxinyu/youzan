package youzan

import "fmt"

const (
	ErrCodeOK = 0
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

package shop
import (
	"net/http"


	"github.com/zihuxinyu/youzan"
	"github.com/zihuxinyu/youzan/shop/request"
	"github.com/zihuxinyu/youzan/shop/response"
)


type Client youzan.Client

func NewClient(appId, appSecret string, clt *http.Client) *Client {
	return (*Client)(youzan.NewClient(appId, appSecret, clt))
}

func (clt *Client)  GetBasic() (basic response.Basic, err error) {

	req := new(request.Basic)
	req.Method = request.MethodBasicGet

	type result struct {
		response.Basic   `json:"response"`
		youzan.Error
	}


	res := new(result)

	err = ((*youzan.Client)(clt)).Post(req, &res)
	if err != nil {
		return
	}
	if res.ErrorResponse.Code != youzan.ErrCodeOK {
		err = &res.Error
	}

	basic = res.Basic

	return
}

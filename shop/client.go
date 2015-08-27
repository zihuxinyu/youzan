package shop
import (
	"net/http"

	"github.com/zihuxinyu/youzan"
	"github.com/zihuxinyu/youzan/shop/request"
	"github.com/zihuxinyu/youzan/shop/response"
)

const (
	MethodGet string = "kdt.shop.basic.get" //获取店铺基本信息
)
type Client youzan.Client

func NewClient(appId, appSecret string, clt *http.Client) *Client {
	return (*Client)(youzan.NewClient(appId, appSecret, clt))
}

//获取店铺基本信息
func (clt *Client)  Get() (resp response.Basic, err error) {

	type basic struct {
		youzan.CommonHeader
	}
	req := new(basic)
	req.Method = MethodGet

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

	resp = res.Basic

	return
}

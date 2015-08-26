package trade
import (
	"net/http"

	"github.com/zihuxinyu/youzan"
	"github.com/zihuxinyu/youzan/trade/request"
	"github.com/zihuxinyu/youzan/trade/response"
)

const (
	MethodTradeGet   string = "kdt.trade.get" //获取单笔交易的信息
	MethodTradeGetSold string = "kdt.trades.sold.get"//查询卖家已卖出的交易列表
)
type Client youzan.Client

func NewClient(appId, appSecret string, clt *http.Client) *Client {
	return (*Client)(youzan.NewClient(appId, appSecret, clt))
}

// 获取单笔交易的信息
func (clt *Client)  Get(req *request.Single) (resp response.TradeDetail, err error) {

	if req.Method == "" {
		req.Method = MethodTradeGet
	}

	type result struct {
		response.TradeDetail   `json:"response"`
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

	resp = res.TradeDetail

	return
}


func (clt *Client)  GetSold(req *request.Sold) (resp response.Sold, err error) {

	if req.Method == "" {
		req.Method = MethodTradeGetSold
	}

	type result struct {
		response.Sold   `json:"response"`
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

	resp = res.Sold

	return
}


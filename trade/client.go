package trade
import (
	"net/http"

	"github.com/zihuxinyu/youzan"
	"github.com/zihuxinyu/youzan/trade/request"
	"github.com/zihuxinyu/youzan/trade/response"
)

const (
	MethodGet   string = "kdt.trade.get" //获取单笔交易的信息
	MethodGetSold string = "kdt.trades.sold.get"//查询卖家已卖出的交易列表
	MethodClose  string = "kdt.trade.close"// 卖家关闭一笔交易
	MethodMemoUpdate string = "kdt.trade.memo.update"// 修改一笔交易备注
)
type Client youzan.Client

func NewClient(appId, appSecret string, clt *http.Client) *Client {
	return (*Client)(youzan.NewClient(appId, appSecret, clt))
}

// 获取单笔交易的信息
func (clt *Client)  Get(req *request.Single) (resp response.Trade, err error) {

	if req.Method == "" {
		req.Method = MethodGet
	}

	type result struct {
		Response struct {
			         Trade response.Trade   `json:"trade"`
		         } `json:"response"`
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

	resp = res.Response.Trade

	return
}

//卖家关闭一笔交易
func (clt *Client)  Close(req *request.Close) (resp response.Trade, err error) {

	if req.Method == "" {
		req.Method = MethodClose
	}

	type result struct {
		Response struct {
			         Trade response.Trade   `json:"trade"`
		         } `json:"response"`
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

	resp = res.Response.Trade

	return
}


//修改一笔交易备注
func (clt *Client)  MemoUpdate(req *request.MemoUpdate) (resp response.Trade, err error) {

	if req.Method == "" {
		req.Method = MethodMemoUpdate
	}

	type result struct {
		Response struct {
			         Trade response.Trade   `json:"trade"`
		         } `json:"response"`
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

	resp = res.Response.Trade

	return
}

//查询卖家已卖出的交易列表
func (clt *Client)  GetSold(req *request.Sold) (resp response.Sold, err error) {

	if req.Method == "" {
		req.Method = MethodGetSold
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


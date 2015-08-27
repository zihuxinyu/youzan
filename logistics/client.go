package logistics
import (
	"net/http"

	"github.com/zihuxinyu/youzan"

	"github.com/zihuxinyu/youzan/logistics/request"
	"github.com/zihuxinyu/youzan/logistics/response"
	"errors"
)

const (
	MethodTraceSearch string = "kdt.logistics.trace.search"// 物流流转信息查询
	MethodOnlineMarkSign string = "kdt.logistics.online.marksign"// 卖家标记签收
	MethodOnlineConfirm string = "kdt.logistics.online.confirm" // 卖家确认发货
)
type Client youzan.Client

func NewClient(appId, appSecret string, clt *http.Client) *Client {
	return (*Client)(youzan.NewClient(appId, appSecret, clt))
}

//物流流转信息查询
// 用户根据交易编号查询物流流转信息，如：2010-8-10 15:23:00 到达杭州集散地。
func (clt *Client)  TraceSearch(req *request.TraceSearch) (resp response.TraceSearchResult, err error) {

	if req.Method == "" {
		req.Method = MethodTraceSearch
	}


	type result struct {
		response.TraceSearchResult   `json:"response"`
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

	resp = res.TraceSearchResult

	return
}

//卖家标记签收
// 标记签收的目的是让交易流程继续走下去，标记签收后交易状态会由【卖家已发货】变为【买家已签收】，通常到店自提的订单需要卖家做标记签收操作
func (clt *Client)  OnlineMarkSign(tid string) (isSuccess bool, err error) {

	if len(tid) == 0 {
		err = errors.New("请输入订单号")
		return
	}
	type OnlineMarkSign struct {
		youzan.CommonHeader
		Tid string `json:"tid"`
	}
	req := new(OnlineMarkSign)

	req.Method = MethodOnlineMarkSign
	req.Tid = tid


	type result struct {
		Response struct {
			         IsSuccess bool `json:"is_success"`
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

	isSuccess = res.Response.IsSuccess

	return
}


//卖家确认发货
//确认发货的目的是让交易流程继续走下去，确认发货后交易状态会由【买家已付款】变为【卖家已发货】
func (clt *Client)  OnlineConfirm(req *request.OnlineConfirm) (isSuccess bool, err error) {

	if req.Method == "" {
		req.Method = MethodOnlineConfirm
	}



	type result struct {
		Response struct {
			         Shipping struct {
				                  IsSuccess bool `json:"is_success"`
			                  } `json:"shipping"`
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
	isSuccess = res.Response.Shipping.IsSuccess

	return
}

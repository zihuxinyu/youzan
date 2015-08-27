package request
import "github.com/zihuxinyu/youzan"
const (
	Status_TRADE_NO_CREATE_PAY string = "TRADE_NO_CREATE_PAY"               //没有创建支付交易
	Status_WAIT_BUYER_PAY  string = "WAIT_BUYER_PAY"                       //等待买家付款
	Status_WAIT_SELLER_SEND_GOODS  string = "WAIT_SELLER_SEND_GOODS"        //等待卖家发货，即：买家已付款
	Status_WAIT_BUYER_CONFIRM_GOODS  string = "WAIT_BUYER_CONFIRM_GOODS"    //等待买家确认收货，即：卖家已发货
	Status_TRADE_BUYER_SIGNED  string = "TRADE_BUYER_SIGNED"                //买家已签收
	Status_TRADE_CLOSED  string = "TRADE_CLOSED"                            //付款以后用户退款成功，交易自动关闭
	Status_ALL_WAIT_PAY  string = "ALL_WAIT_PAY"                            //包含：WAIT_BUYER_PAY、TRADE_NO_CREATE_PAY
	Status_ALL_CLOSED  string = "ALL_CLOSED"                                //所有关闭订单
)
type Sold struct {
	youzan.CommonHeader
	Fields       string   `json:"fields,omitempty"`        //需要返回的交易对象字段，如tid,title,receiver_city等。可选值：TradeDetail交易结构体中所有字段均可返回；多个字段用“,”分隔。如果为空则返回所有
	Status       string   `json:"status,omitempty"`        //交易状态，默认查询所有交易状态的数据，除了默认值外每次只能查询一种状态。
                                                           //	可选值：
                                                           //TRADE_NO_CREATE_PAY（没有创建支付交易）
                                                           //WAIT_BUYER_PAY（等待买家付款）
                                                           //WAIT_SELLER_SEND_GOODS（等待卖家发货，即：买家已付款）
                                                           //WAIT_BUYER_CONFIRM_GOODS（等待买家确认收货，即：卖家已发货）
                                                           //TRADE_BUYER_SIGNED（买家已签收）
                                                           //TRADE_CLOSED（付款以后用户退款成功，交易自动关闭）
                                                           //ALL_WAIT_PAY（包含：WAIT_BUYER_PAY、TRADE_NO_CREATE_PAY）
                                                           //ALL_CLOSED（所有关闭订单）
	StartCreated string   `json:"start_created,omitempty"` //交易创建开始时间。查询在该时间之后（包含该时间）创建的交易，为空则不限制
	EndCreated   string   `json:"end_created,omitempty"`   //交易创建结束时间。查询在该时间之前创建的交易，为空则不限制
	StartUpdate  string   `json:"start_update,omitempty"`  //交易状态更新的开始时间。查询在该时间之后（包含该时间）交易状态更新过的交易，为空则不限制
	EndUpdate    string   `json:"end_update,omitempty"`    //交易状态更新的结束时间。查询在该时间之前交易状态更新过的交易，为空则不限制
	WeixinUserId int64   `json:"weixin_user_id,omitempty"` //微信粉丝ID
	BuyerNick    string   `json:"buyer_nick,omitempty"`    //买家昵称
	PageNo       int64   `json:"page_no,omitempty"`        //页码。取值范围：大于零的整数；默认值：1
	PageSize     int64   `json:"page_size,omitempty"`      //每页条数。取值范围：大于零的整数；默认值：40；最大值：100
	UseHasNext   bool   `json:"use_has_next,omitempty"`    //是否启用has_next的分页方式，是的话返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段

}

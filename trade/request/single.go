package request
import "github.com/zihuxinyu/youzan"
type Single struct {
	youzan.CommonHeader

	Fields           string    `json:"fields,omitempty"`              //需要返回的交易对象字段
                                                                      // 如tid,title,receiver_city等。
                                                                      // 可选值：TradeDetail交易结构体中所有字段均可返回；
                                                                      // 多个字段用“,”分隔。
                                                                      // 如果为空则返回所有
	Tid              string    `json:"tid"`                           //交易编号
	SubTradePageNo   int64     `json:"sub_trade_page_no,omitempty"`   //指定获取子交易的第几页，不传则获取全部
	subTradePageSize int64     `json:"sub_trade_page_size,omitempty"` //指定获取子交易每页条数，不传则获取全部，上限500
}

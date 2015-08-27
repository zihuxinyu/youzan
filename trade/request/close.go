package request
import "github.com/zihuxinyu/youzan"

//关闭一笔交易
type Close struct {
	youzan.CommonHeader
	Tid         string `json:"tid"`                    //交易编号
	CloseReason string `json:"close_reason,omitempty"` //关闭交易的原因
	Fields      string `json:"fields,omitempty"`       //需要返回的交易对象字段，
                                                       // 如tid,title,receiver_city等。可选值：TradeDetail交易结构体中所有字段均可返回；
                                                       // 多个字段用“,”分隔。如果为空则返回所有
}

package request
import "github.com/zihuxinyu/youzan"

//修改一笔交易备注
type MemoUpdate struct {
	youzan.CommonHeader
	Tid    string `json:"tid"`              //交易编号
	Memo   string `json:"memo,omitempty"`   //交易备注
	Flag   int `json:"flag,omitempty"`   //交易备注加星标，取值为1-5
	Fields string `json:"fields,omitempty"` //需要返回的交易对象字段，如tid,title,receiver_city等。
                                            // 可选值：TradeDetail交易结构体中所有字段均可返回；
                                            // 多个字段用“,”分隔。如果为空则返回所有
}

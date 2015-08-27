package request
import "github.com/zihuxinyu/youzan"
type TraceSearch struct {
	youzan.CommonHeader
	Tid string `json:"tid"` //交易编号
}

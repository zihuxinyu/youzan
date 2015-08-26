package response
import "encoding/json"
type Sold struct {
	TotalResults json.Number `json:"total_results"`    //搜索到的交易总数
	Trades       []Trade `json:"trades"` //搜索到的交易列表
	HasNext      bool `json:"has_next"`        //是否存在下一页
}

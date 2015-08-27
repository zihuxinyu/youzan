package request
import "github.com/zihuxinyu/youzan"

//分页获取商品自定义标签列表
type TagsGetPage struct {
	youzan.CommonHeader
	PageNo   int64  `json:"page_no,omitempty"`   //页码
	PageSize int64  `json:"page_size,omitempty"` //每页条数
	OrderBy  int64  `json:"order_by,omitempty"`  //排序方式。
                                                 // 格式为column:asc/desc，
                                                 // column可选值：created 创建时间
}
// 获取商品自定义标签列表
type TagsGet struct {
	youzan.CommonHeader
	IsSort bool `json:"is_sort"` //是否排序（按序号）
}

package request
import "github.com/zihuxinyu/youzan"

type Region struct {
	youzan.CommonHeader
	Level    int64 `json:"level"`                 // 必须 是 要获取的区域等级，
                                                  //          0 获取所有区域列表，
                                                  //          1 获取省份列表，
                                                  //          2 获取城市列表，
                                                  //          3 获取县区列表，
                                                  //          4 获取指定ID区域及其上级区域信息

	ParentId int64   `json:"parent_id,omitempty"` // 必须 否 区域父级ID
                                                  ///	       当level为0或1时，parent_id可以为空；
                                                  //         当level为2时，parent_id须为某省份ID；
                                                  //         当level为3时，parent_id须为某城市ID

	Id       int64     `json:"id,omitempty"`      // 必须 否  区域ID
                                                  //          当level为4时，id参数生效且不能为空

	Fields   string     `json:"fields,omitempty"` // 必须 否  需要返回的区域地名对象字段，如id,name等。
                                                  //          可选值：区域数据结构体中所有字段均可返回；
                                                  //          多个字段用“,”分隔。
                                                  //          如果为空则返回所有
}

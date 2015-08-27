package response
import "encoding/json"
//商品分类数据结构
type GoodsCategory struct {
	Cid           int64 `json:"cid"`                      //商品分类的ID
	ParentCid     int64 `json:"parent_cid"`               //父分类ID，等于0时，代表的是一级的分类
	Name          string `json:"name"`                    //商品分类的名称
	IsParent      bool `json:"is_parent"`                 //该分类是否为父分类(即：该分类是否还有子分类)
	SubCategories []GoodsCategory `json:"sub_categories"` //子分类列表。
                                                          // 如果is_parent为false，则该字段为空
}


//商品推广栏目数据结构
type GoodsPromotionCategory struct {
	Id   int64 `json:"id"`   //商品推广栏目的ID
	Name string `json:"name"` //商品推广栏目的名称
}


//商品标签数据结构
type GoodsTag struct {
	Id       int64  `json:"id"`         //商品标签的ID
	Name     string  `json:"name"`      //商品标签的名称
	TagType  json.Number  `json:"type"` //todo 官方标记number 输出为string
										// 商品标签类型，
                                        // 0 自定义，
                                        // 1 未分类，
                                        // 2 最新，
                                        // 3 最热，
                                        // 4 隐藏
	Created  string  `json:"created"`   //商品标签创建时间
	ItemNum  int64  `json:"item_num"`   //商品标签内的商品数量
	TagUrl   string  `json:"tag_url"`   //商品标签的展示的URL地址
	ShareUrl string  `json:"share_url"` //分享出去的商品标签展示地址
	Desc     string  `json:"desc"`      //商品标签描述
}

//商品标签列表
type GoodsTagList struct {
	Tags         []GoodsTag `json:"tags"`      //商品自定义标签列表
	TotalResults string `json:"total_results"` //结果总数
}

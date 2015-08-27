package request
import "github.com/zihuxinyu/youzan"

const (
	InventoryBannerForShelved string = "for_shelved"//已下架的
	InventoryBannerSoldOut string = "sold_out"//已售罄的

)
// 根据外部编号取商品Sku
type SkuCustomGet struct {
	youzan.CommonHeader
	Fields  string `json:"fields,omitempty"` //需要返回的Sku对象字段，如sku_id,num_iid,quantity等。可选值：Sku结构体中所有字段均可返回；多个字段用“,”分隔。如果为空则返回所有
	OuterId string `json:"outer_id"`         //商家编码（商家为Sku设置的外部编号）
	NumIid  int64 `json:"num_iid"`           //商品数字编号
}

//获取仓库中的商品列表
type InventoryGet struct {
	youzan.CommonHeader
	Fields   string `json:"fields,omitempty"`   //需要返回的商品对象字段，如title,price,desc等。可选值：Item商品结构体中所有字段均可返回；多个字段用“,”分隔。如果为空则返回所有
	Q        string `json:"q,omitempty"`        //搜索字段。搜索商品的title
	Banner   string `json:"banner,omitempty"`   //分类字段。可选值：for_shelved（已下架的）/ sold_out（已售罄的）
	TagId    int64 `json:"tag_id,omitempty"`    //商品标签的ID
	PageNo   int64 `json:"page_no,omitempty"`   //页码
	PageSize int64 `json:"page_size,omitempty"` //每页条数
	OrderBy  string `json:"order_by,omitempty"` //排序方式。格式为column:asc/desc，column可选值：created 创建时间 / modified 修改时间
}
// 根据商品货号获取商品
type CustomGet struct {
	youzan.CommonHeader
	Fields  string `json:"fields,omitempty"` //需要返回的Sku对象字段，如sku_id,num_iid,quantity等。可选值：Sku结构体中所有字段均可返回；多个字段用“,”分隔。如果为空则返回所有
	OuterId string `json:"outer_id"`         //商家编码（商家为Sku设置的外部编号）
}

//商品上架 下架
type UpdateListing struct {
	youzan.CommonHeader
	NumIid string `json:"num_iid"`          //商品数字编号
	Fields string `json:"fields,omitempty"` //需要返回的商品对象字段，如title,price,desc等。可选值：Item商品结构体中所有字段均可返回；多个字段用“,”分隔。如果为空则返回所有
}

//得到单个商品信息
type ItemGet struct {
	youzan.CommonHeader
	NumIid string `json:"num_iid,omitempty"` //商品数字编号
	Fields string `json:"fields,omitempty"`  //需要返回的商品对象字段，如title,price,desc等。可选值：Item商品结构体中所有字段均可返回；多个字段用“,”分隔。如果为空则返回所有
	Alias  string `json:"alias,omitempty"`   //商品别名
                                             //调用时，参数 num_iid 和 alias 必须选其一
}


//更新SKU信息
type SkuUpdate struct {
	NumIid   int64  `json:"num_iid"`   //商品数字编号
	SkuId    int64  `json:"sku_id"`    //Sku数字ID
	Quantity int64  `json:"quantity"`  //Sku的库存数量
	Price    float64  `json:"price"`   //Sku的销售价格。精确到2位小数；单位：元
	OuterId  string  `json:"outer_id"` //商家编码（商家为Sku设置的外部编号）
}



//更新单个商品信息
type ItemUpdate struct {
	NumIid         string `json:"num_iid"`           //商品数字编号
	Cid            string `json:"cid"`               //商品分类的叶子类目id，可参考API：kdt.itemcategories.get
	PromotionCid   string `json:"promotion_cid"`     //商品推广栏目id，可参考API：kdt.itemcategories.promotions.get
	TagIds         string `json:"tag_ids"`           //商品标签id串，结构如：1234,1342,...，可参考API：kdt.itemcategories.tags.get
	Price          string `json:"price"`             //商品价格。取值范围：0.01-100000000；精确到2位小数；单位：元
	Title          string `json:"title"`             //商品标题。不能超过100字，受违禁词控制
	Desc           string `json:"desc"`              //商品描述。字数要大于5个字符，小于25000个字符 ，受违禁词控制
	Images         string `json:"images[]"`          //商品图片文件列表，可一次上传多张。最大支持 1M，支持的文件类型：gif,jpg,jpeg,png
                                                     //	注：图片参数不参与通讯协议签名，参数名中的中括号"[]"必须有，否则不能正常工作
	KeepItemImgIds string `json:"keep_item_img_ids"` //编辑商品时保留商品已有图片，值为图片的id，多个图片id用逗号分隔。如果不传递该参数，则保留所有图片；如果为空，则删除所有图片
	PostFee        string `json:"post_fee"`          //运费。取值范围：0.00-999.00；精确到2位小数；单位：元
	SkuProperties  string `json:"sku_properties"`    //Sku的属性串。格式：pText:vText;pText:vText，多个sku之间用逗号分隔，如：颜色:黄色;尺寸:M,颜色:黄色;尺寸:S。pText和vText文本中不可以存在冒号和分号以及逗号

                                                     //	为了兼顾移动端商品界面展示的美观，目前有赞仅支持Sku的属性个数小于等于三个（比如：颜色、尺寸、重量 这三个属性）
	SkuQuantities  string `json:"sku_quantities"`    // Sku的数量串。结构如：num1,num2,num3 如：2,3
	SkuPrices      string `json:"sku_prices"`        // Sku的价格串。结构如：10.00,5.00,... 精确到2位小数。单位:元
	SkuOuterIds    string `json:"sku_outer_ids"`     //Sku的商家编码（商家为Sku设置的外部编号）串。结构如：1234,1342,... 。sku_properties, sku_quantities, sku_prices, sku_outer_ids在输入数据时要一一对应，即使商家编码为空，也要用逗号相连
	SkusWithJson   string `json:"skus_with_json"`    //商品Sku信息的Json字符串， 该字段支持商品的所有Sku信息（sku_properties、sku_quantities、sku_prices、sku_outer_ids中内容）一起写入。
                                                     //	注意：仅仅当传入该字段且不传入sku_properties、sku_quantities、sku_prices、sku_outer_ids字段时，系统才会对该字段进行解析， 否则默认只解析sku_properties、sku_quantities、sku_prices、sku_outer_ids。
                                                     //json格式参考http://open.koudaitong.com/doc/api?method=kdt.item.update

	OriginPrice    string `json:"origin_price"`      //显示在“原价”一栏中的信息
	BuyUrl         string `json:"buy_url"`           //该商品的外部购买地址。当用户购买环境不支持微信或微博支付时会跳转到此地址
	OuterId        string `json:"outer_id"`          //商品货号（商家为商品设置的外部编号）
	BuyQuota       string `json:"buy_quota"`         //每人限购多少件。0代表无限购，默认为0
	Quantity       string `json:"quantity"`          //商品总库存。当商品没有Sku的时候有效，商品有Sku时，总库存会自动按所有Sku库存之和计算
	HideQuantity   string `json:"hide_quantity"`     //是否隐藏商品库存。在商品展示时不显示商品的库存，默认0：显示库存，设置为1：不显示库存
	Fields         string `json:"fields"`            //需要返回的商品对象字段，如title,price,desc等。可选值：Item商品结构体中所有字段均可返回；多个字段用“,”分隔。如果为空则返回所有
}



type  ItemAdd struct {
	Cid           string `json:"cid"`            //商品分类的叶子类目id，可参考API：kdt.itemcategories.get商品标题。不能超过100字，受违禁词控制
	PromotionCid  string `json:"promotion_cid"`  //商品推广栏目id，可参考API：kdt.itemcategories.promotions.get商品描述。字数要大于5个字符，小于25000个字符 ，受违禁词控制
	TagIds        string `json:"tag_ids"`        //商品标签id串，结构如：1234,1342,...，可参考API：kdt.itemcategories.tags.get是否是虚拟商品。0为否，1为是。目前不支持虚拟商品
	Price         string `json:"price"`          //商品价格。取值范围：0.01-100000000；精确到2位小数；单位：元。需要在Sku价格所决定的的区间内
	Title         string `json:"title"`          //商品标题。不能超过100字，受违禁词控制
	Desc          string `json:"desc"`           //商品描述。字数要大于5个字符，小于25000个字符 ，受违禁词控制
	IsVirtual     string `json:"is_virtual"`     //是否是虚拟商品。0为否，1为是。目前不支持虚拟商品
	Images        string `json:"images[]"`       //商品图片文件列表，可一次上传多张。最大支持 1M，支持的文件类型：gif,jpg,jpeg,png
	PostFee       string `json:"post_fee"`       //运费。取值范围：0.00-999.00；精确到2位小数；单位：元
	SkuProperties string `json:"sku_properties"` //Sku的属性串。格式：pText:vText;pText:vText，多个sku之间用逗号分隔，如：颜色:黄色;尺寸:M,颜色:黄色;尺寸:S。
                                                 //pText和vText文本中不可以存在冒号和分号以及逗号
                                                 //为了兼顾移动端商品界面展示的美观，目前有赞仅支持Sku的属性个数小于等于三个（比如：颜色、尺寸、重量 这三个属性）。无Sku则为空
	SkuQuantities string `json:"sku_quantities"` //Sku的数量串。结构如：num1,num2,num3 如：2,3。无Sku则为空
	SkuPrices     string `json:"sku_prices"`     //Sku的价格串。结构如：10.00,5.00,... 精确到2位小数。单位:元。无Sku则为空
	SkuOuter_ids  string `json:"sku_outer_ids"`  //Sku的商家编码（商家为Sku设置的外部编号）串。结构如：1234,1342,... 。sku_properties, sku_quantities, sku_prices, sku_outer_ids在输入数据时要一一对应，即使商家编码为空，也要用逗号相连。无Sku则为空
	SkusWith_json string `json:"skus_with_json"` //商品Sku信息的Json字符串， 调用时，参数 sku_properties、sku_quantities、sku_prices、sku_outer_ids四个字段组合方式 和 skus_with_json 单字段输入方式 选其一个方式即可，无Sku则为空。
	OriginPrice   string `json:"origin_price"`   //具体参见kdt.item.update文档描述。
	BuyUrl        string `json:"buy_url"`        //显示在“原价”一栏中的信息
	OuterId       string `json:"outer_id"`       //该商品的外部购买地址。当用户购买环境不支持微信或微博支付时会跳转到此地址
	BuyQuota      string `json:"buy_quota"`      //商品货号（商家为商品设置的外部编号）
	Quantity      string `json:"quantity"`       //每人限购多少件。0代表无限购，默认为0
	HideQuantity  string `json:"hide_quantity"`  //商品总库存。当商品没有Sku的时候有效，商品有Sku时，总库存会自动按所有Sku库存之和计算
	Fields        string `json:"fields"`         //是否隐藏商品库存。在商品展示时不显示商品的库存，默认0：显示库存，设置为1：不显示库存
	IsDisplay     string `json:"is_display"`     //需要返回的商品对象字段，如title,price,desc等。可选值：Item商品结构体中所有字段均可返回；多个字段用“,”分隔。如果为空则返回所有

}

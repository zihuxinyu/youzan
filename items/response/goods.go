package response
import "encoding/json"

//Sku数据结构
type GoodsSku struct {
	OuterId            string `json:"outer_id"`                      //商家编码（商家为Sku设置的外部编号）
	SkuId              string `json:"sku_id"`                        //Sku的ID，sku_id 在系统里并不是唯一的，结合商品ID一起使用才是唯一的。
	SkuUniqueCode      string `json:"sku_unique_code"`               //Sku在系统中的唯一编号，可以在开发者的系统中用作 Sku 的唯一ID，但不能用于调用接口
	NumIid             string `json:"num_iid"`                       //Sku所属商品的数字编号
	Quantity           string `json:"quantity"`                      //属于这个Sku的商品的数量
	PropertiesName     string `json:"properties_name"`               //Sku所对应的销售属性的中文名字串，格式如：pid1:vid1:pid_name1:vid_name1;pid2:vid2:pid_name2:vid_name2……
	PropertiesNameJson json.RawMessage `json:"properties_name_json"` //Sku所对应的销售属性的Json字符串（需另行解析）， 该字段内容与properties_name字段除了格式不一样，内容完全一致。 由于产品规格信息难以避免涉及到‘:’、‘,’、‘;’这些与解析规则冲突的字符，所以增加该字段。
	WithHoldQuantity   string `json:"with_hold_quantity"`            //商品在付款减库存的状态下，该Sku上未付款的订单数量
	Price              string `json:"price"`                         //商品的这个Sku的价格；精确到2位小数；单位：元
	Created            string `json:"created"`                       //Sku创建日期，时间格式：yyyy-MM-dd HH:mm:ss
	Modified           string `json:"modified"`                      //Sku最后修改日期，时间格式：yyyy-MM-dd HH:mm:ss
}

//商品图片数据结构
type GoodsImage   struct {
	ID        int `json:"id"`           //商品图片的ID
	Created   string `json:"created"`   //图片创建时间，时间格式：yyyy-MM-dd HH:mm:ss
	URL       string `json:"url"`       //图片链接地址
	Thumbnail string `json:"thumbnail"` //图片缩略图链接地址
	Medium    string `json:"medium"`    //中号大小图片链接地址
	Combine   string `json:"combine"`   //组合图片链接地址
}



//商品标签数据结构
type GoodsTag   struct {
	ID       int `json:"id"`           //商品标签的ID
	Name     string `json:"name"`      //商品标签的名称
	Type     string `json:"type"`      //商品标签类型，0 自定义，1 未分类，2 最新，3 最热，4 隐藏
	Created  string `json:"created"`   //商品标签创建时间
	ItemNum  int `json:"item_num"`     //商品标签内的商品数量
	TagURL   string `json:"tag_url"`   //商品标签的展示的URL地址
	ShareURL string `json:"share_url"` //分享出去的商品标签展示地址
	Desc     string `json:"desc"`      //商品标签描述
}

//商品二维码数据结构
type  GoodsQrcode  struct {
	ID              int `json:"id"`                   //商品二维码的ID
	Name            string `json:"name"`              //二维码的名称
	Desc            string `json:"desc"`              //二维码的描述
	Created         string `json:"created"`           //商品二维码创建时间，时间格式：yyyy-MM-dd HH:mm:ss
	Type            string `json:"type"`              //商品二维码类型。可选值：
                                                      //	GOODS_SCAN_FOLLOW(扫码关注后购买商品)
                                                      //	GOODS_SCAN_FOLLOW_DECREASE(扫码关注后减优惠额)
                                                      //	GOODS_SCAN_FOLLOW_DISCOUNT(扫码关注后折扣)
                                                      //	GOODS_SCAN_DECREASE(扫码直接减优惠额)
                                                      //	GOODS_SCAN_DISCOUNT(扫码直接折扣)
	Discount        string `json:"discount"`          //折扣，格式：9.0；单位：折，精确到小数点后 1 位。如果没有折扣，则为 0
	Decrease        string `json:"decrease"`          //减金额优惠，格式：5.00；单位：元；精确到：分。如果没有减额优惠，则为 0
	LinkURL         string `json:"link_url"`          //扫码直接购买的二维码基于这个url生成。如果不是扫码直接购买的类型，则为空
	WeixinQrcodeURL string `json:"weixin_qrcode_url"` //扫码关注购买的二维码图片地址。如果不是扫码关注购买的类型，则为空
}


//商品数据结构
type GoodsDetail struct {
	NumIid              int `json:"num_iid"`                  //商品数字编号
	Alias               string `json:"alias"`                 //商品别称
	Title               string `json:"title"`                 //商品标题
	Cid                 int `json:"cid"`                      //商品分类的叶子类目id，可参考API：kdt.itemcategories.get
	PromotionCid        int `json:"promotion_cid"`            //商品推广栏目id，可参考API：kdt.itemcategories.promotions.get
	TagIds              string `json:"tag_ids"`               //商品标签id串，结构如：1234,1342,...，可参考API：kdt.itemcategories.tags.get
	Desc                string `json:"desc"`                  //商品描述
	OriginPrice         string `json:"origin_price"`          //显示在“原价”一栏中的信息
	OuterID             string `json:"outer_id"`              //商品货号（商家为商品设置的外部编号，可与商家外部系统对接）
	OuterBuyURL         string `json:"outer_buy_url"`         //商品外部购买链接
	BuyQuota            int `json:"buy_quota"`                //每人限购多少件。0代表无限购，默认为0
	Created             string `json:"created"`               //商品的发布时间
	IsVirtual           bool `json:"is_virtual"`              //是否为虚拟商品
	IsListing           bool `json:"is_listing"`              //商品上架状态。true 为已上架，false 为已下架
	IsLock              bool `json:"is_lock"`                 //商品是否锁定。true 为已锁定，false 为未锁定
	IsUsed              bool `json:"is_used"`                 //是否为二手商品
	ProductType         string `json:"product_type"`          //商品类型
	AutoListingTime     string `json:"auto_listing_time"`     //商品定时上架（定时开售）的时间。没设置则为空
	DetailURL           string `json:"detail_url"`            //适合wap应用的商品详情url
	ShareURL            string `json:"share_url"`             //分享出去的商品详情url
	PicURL              string `json:"pic_url"`               //商品主图片地址
	PicThumbURL         string `json:"pic_thumb_url"`         //商品主图片缩略图地址
	Num                 int `json:"num"`                      //商品数量
	SoldNum             int `json:"sold_num"`                 //商品销量
	Price               int `json:"price"`                    //商品价格，格式：5.00；单位：元；精确到：分
	PostType            int `json:"post_type"`                //运费类型。 1：统一邮费 2：运费模版;
	PostFee             string `json:"post_fee"`              //运费（针对“统一运费”），格式：5.00；单位：元；精确到：分
	DeliveryTemplateFee string `json:"delivery_template_fee"` //运费（针对“运费模版”），格式：5.00；单位：元；精确到：分 若存在运费区间，中间用逗号隔开，如 “5.00,9.00”
	Skus                []GoodsSku `json:"skus"`              //Sku列表
	ItemImgs            []GoodsImage `json:"item_imgs"`       //商品图片列表
	ItemQrcodes         []GoodsQrcode `json:"item_qrcodes"`   //商品二维码列表。
	ItemTags            []GoodsTag`json:"item_tags"`          //商品标签数据结构
	ItemType            int `json:"item_type"`                //商品类型。
                                                              //0：普通商品；
                                                              //10：分销商品;
	IsSupplierItem      bool `json:"is_supplier_item"`        //是否是供货商商品
	LikeCount           int `json:"like_count"`               //喜欢数目

}


//商品列表
type GoodsList struct {
	Items        []GoodsDetail`json:"items"`
	TotalResults int `json:"total_results"`
}

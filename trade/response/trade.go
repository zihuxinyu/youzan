package response
import (
	"encoding/json"
)

type TradeDetail struct {
	Trade Trade `json:"trade"`
}

type Trade struct {
	Tid              string `json:"tid"`                    //交易编号
	Num              int64   `json:"num"`                   //商品购买数量。当一个trade对应多个order的时候，值为所有商品购买数量之和
	NumIid           string  `json:"num_iid"`               //商品数字编号。当一个trade对应多个order的时候，值为第一个交易明细中的商品的编号
	Price            string  `json:"price"`                 //商品价格。精确到2位小数；单位：元。当一个trade对应多个order的时候，值为第一个交易明细中的商品的价格
	PicPath          string `json:"pic_path"`               //商品主图片地址。当一个trade对应多个order的时候，值为第一个交易明细中的商品的图片地址
	PicThumbPath     string `json:"pic_thumb_path"`         //商品主图片缩略图地址
	Title            string `json:"title"`                  //交易标题，以首个商品标题作为此标题的值
	Type             string `json:"type"`                   //交易类型。取值范围：
                                                            //FIXED （一口价）
                                                            //GIFT （送礼）
                                                            //BULK_PURCHASE（来自分销商的采购）
                                                            //PRESENT （赠品领取）
                                                            //COD （货到付款）
                                                            //QRCODE（扫码商家二维码直接支付的交易）
	WeixinUserID     json.Number  `json:"weixin_user_id"`   //微信粉丝ID
	BuyerType        json.Number `json:"buyer_type"`        //买家类型，取值范围：0 为未知，1 为微信粉丝，2 为微博粉丝
	BuyerID          json.Number `json:"buyer_id"`          //买家ID，当 buyer_type 为 1 时，buyer_id 的值等于 weixin_user_id 的值
	BuyerNick        string `json:"buyer_nick"`             //买家昵称
	BuyerMessage     string `json:"buyer_message"`          //买家购买附言

	SellerFlag       int64  `json:"seller_flag"`            //卖家备注星标，取值范围 1、2、3、4、5；
                                                            //	如果为0，表示没有备注星标
	TradeMemo        string `json:"trade_memo"`             //卖家对该交易的备注
	ReceiverCity     string `json:"receiver_city"`          //收货人的所在城市。
                                                            //	PS：如果订单类型是送礼订单，收货地址在sub_trades字段中；如果物流方式是到店自提，收货地址在fetch_detail字段中

	ReceiverDistrict string `json:"receiver_district"`      //收货人的所在地区
	ReceiverName     string `json:"receiver_name"`          //收货人的姓名
	ReceiverState    string `json:"receiver_state"`         //收货人的所在省份
	ReceiverAddress  string `json:"receiver_address"`       //收货人的详细地址
	ReceiverZip      string `json:"receiver_zip"`           //收货人的邮编
	ReceiverMobile   string `json:"receiver_mobile"`        //收货人的手机号码
	Feedback         int64 `json:"feedback"`                //交易维权状态。
                                                            //0 无维权，1 顾客发起维权，2 顾客拒绝商家的处理结果，3 顾客接受商家的处理结果，9 商家正在处理， 101 等待卖家同意退款申请，102 等待卖家同意退款申请（维权失败过），103 卖家不同意退款申请，104 已经申请有赞客满介入， 105 卖家已同意退款，106 已退货，等待卖家确认收货，107 维权已经关闭，110 退款成功。
                                                            //备注：1到10的状态码是微信那边的维权状态码，100以上的状态码是有赞这边的维权状态码
	RefundState      string `json:"refund_state"`           //退款状态。取值范围：
                                                            //NO_REFUND（无退款）
                                                            //PARTIAL_REFUNDING（部分退款中）
                                                            //PARTIAL_REFUNDED（已部分退款）
                                                            //PARTIAL_REFUND_FAILED（部分退款失败）
                                                            //FULL_REFUNDING（全额退款中）
                                                            //FULL_REFUNDED（已全额退款）
                                                            //FULL_REFUND_FAILED（全额退款失败）
	OuterTid         string `json:"outer_tid"`              //外部交易编号。比如，如果支付方式是微信支付，就是财付通的交易单号
	Status           string `json:"status"`                 //交易状态。取值范围：
                                                            //TRADE_NO_CREATE_PAY (没有创建支付交易)
                                                            //WAIT_BUYER_PAY (等待买家付款)
                                                            //WAIT_SELLER_SEND_GOODS (等待卖家发货，即：买家已付款)
                                                            //WAIT_BUYER_CONFIRM_GOODS (等待买家确认收货，即：卖家已发货)
                                                            //TRADE_BUYER_SIGNED (买家已签收)
                                                            //TRADE_CLOSED (付款以后用户退款成功，交易自动关闭)
                                                            //TRADE_CLOSED_BY_USER (付款以前，卖家或买家主动关闭交易)
	ShippingType     string `json:"shipping_type"`          //创建交易时的物流方式。取值范围：express（快递），fetch（到店自提）

	PostFee          json.Number `json:"post_fee"`          //运费。单位：元，精确到分
	TotalFee         json.Number `json:"total_fee"`         //商品总价（商品价格乘以数量的总金额）。单位：元，精确到分
	RefundedFee      json.Number `json:"refunded_fee"`      //交易完成后退款的金额。单位：元，精确到分
	DiscountFee      json.Number `json:"discount_fee"`      //交易优惠金额（不包含交易明细中的优惠金额）。单位：元，精确到分

	Payment          string `json:"payment"`                //实付金额。单位：元，精确到分
	Created          string  `json:"created"`               //交易创建时间
	UpdateTime       string `json:"update_time"`            //交易更新时间。当交易的：状态改变、备注更改、星标更改 等情况下都会刷新更新时间
	PayTime          string `json:"pay_time"`               //买家付款时间
	PayType          string `json:"pay_type"`               //支付类型。取值范围：
                                                            //WEIXIN (微信支付)
                                                            //ALIPAY (支付宝支付)
                                                            //BANKCARDPAY (银行卡支付)
                                                            //PEERPAY (代付)
                                                            //CODPAY (货到付款)
                                                            //BAIDUPAY (百度钱包支付)
                                                            //PRESENTTAKE (直接领取赠品)
                                                            //COUPONPAY（优惠券/码全额抵扣）
                                                            //BULKPURCHASE（来自分销商的采购）
	ConsignTime      string `json:"consign_time"`           //卖家发货时间
	SignTime         string `json:"sign_time"`              //买家签收时间
	BuyerArea        string `json:"buyer_area"`             //买家下单的地区
	AdjustFee        string `json:"adjust_fee"`             //卖家手工调整订单金额。精确到2位小数；单位：元。若卖家减少订单金额10元2分，则这里为10.02；若卖家增加订单金额10元2分，则这里为-10.02
	Profit           string `json:"profit"`
	Handled          string `json:"handled"`
	SubTrades        []Trade `json:"sub_trades"`            //交易中包含的子交易，目前：仅在送礼订单中会有子交易
	Orders           []Order `json:"orders"`                //交易明细列表
	FetchDetail      Fetch `json:"fetch_detail"`            //如果是到店自提交易，返回自提详情，否则返回空
	CouponDetails    []Coupon `json:"coupon_details"`       //在交易中使用到的卡券的详情，包括：优惠券、优惠码
	PromotionDetails []Promotion `json:"promotion_details"` //在交易中使用到优惠活动详情，包括：满减满送

}


//交易明细数据结构
type Order struct {
	Oid                   int64 `json:"oid"`                                //交易明细编号。该编号并不唯一，只用于区分交易内的多条明细记录
	NumIid                json.Number `json:"num_iid"`                      //商品数字编号
	SkuID                 json.Number `json:"sku_id"`                       //Sku的ID，sku_id 在系统里并不是唯一的，结合商品ID一起使用才是唯一的。
	SkuUniqueCode         string `json:"sku_unique_code"`                   //Sku在系统中的唯一编号，可以在开发者的系统中用作 Sku 的唯一ID，但不能用于调用接口
	Num                   json.Number `json:"num"`                          //商品购买数量
	OuterSkuID            string `json:"outer_sku_id"`                      //商家编码（商家为Sku设置的外部编号）
	OuterItemID           string `json:"outer_item_id"`                     //商品货号（商家为商品设置的外部编号）
	Title                 string `json:"title"`                             //商品标题
	SellerNick            string `json:"seller_nick"`                       //卖家昵称
	FenxiaoPrice          json.Number `json:"fenxiao_price"`                //商品在分销商那边的出售价格。精确到2位小数；单位：元。如果是采购单才有值，否则值为 0
	FenxiaoPayment        json.Number `json:"fenxiao_payment"`              //商品在分销商那边的实付金额。精确到2位小数；单位：元。如果是采购单才有值，否则值为 0
	Price                 json.Number `json:"price"`                        //商品价格。精确到2位小数；单位：元
	TotalFee              json.Number `json:"total_fee"`                    //应付金额（商品价格乘以数量的总金额）
	DiscountFee           json.Number `json:"discount_fee"`                 //交易明细内的优惠金额。精确到2位小数，单位：元
	Payment               json.Number `json:"payment"`                      //实付金额。精确到2位小数，单位：元
	SkuPropertiesName     string `json:"sku_properties_name"`               //SKU的值，即：商品的规格。如：机身颜色:黑色;手机套餐:官方标配
	PicPath               string `json:"pic_path"`                          //商品主图片地址
	PicThumbPath          string `json:"pic_thumb_path"`                    //商品主图片缩略图地址
	ItemType              int `json:"item_type"`                            //商品类型。0：普通商品；10：分销商品
	BuyerMessages         []BuyerMessage  `json:"buyer_messages"`           //交易明细中的买家留言列表
	OrderPromotionDetails []OrderPromotion `json:"order_promotion_details"` //交易明细中的优惠信息列表
}

//交易明细中买家留言的数据结构
type BuyerMessage struct {
	Title   string `json:"title"`   //留言的标题
	Content string `json:"content"` //留言内容
}
type  OrderPromotion struct {
	PromotionName string `json:"promotion_name"`    //优惠的名称
	PromotionType string `json:"promotion_type"`    //优惠的类型。可选值：
                                                    //	MEMBER_CARD_DISCOUNT（会员卡折扣）
                                                    //SCAN_DISCOUNT（扫码折扣）
                                                    //SCAN_DECREASE（扫码减额优惠）
	ApplyAt       string `json:"apply_at"`          //应用优惠的时间
	DiscountFee   json.Number `json:"discount_fee"` //优惠的金额，单位：元，精确到小数点后两位

}

//到店自提详情
type Fetch struct {
	FetcherName   string `json:"fetcher_name"`   //领取人（即：预约人）的姓名
	FetcherMobile string `json:"fetcher_mobile"` //领取人的手机号
	FetchTime     string `json:"fetch_time"`     //预约的领取时间
	ShopName      string `json:"shop_name"`      //自提点名称
	ShopMobile    string `json:"shop_mobile"`    //自提点联系方式
	ShopState     string `json:"shop_state"`     //自提点所在省份
	ShopCity      string `json:"shop_city"`      //自提点所在城市
	ShopDistrict  string `json:"shop_district"`  //自提点所在地区
	ShopAddress   string `json:"shop_address"`   //自提点详细地址
}

//订单中使用到的卡券的数据结构
type  Coupon struct {
	CouponID          int64 `json:"coupon_id"`           //该组卡券的ID
	CouponName        string `json:"coupon_name"`        //该组卡券的名称
	CouponType        string `json:"coupon_type"`        //卡券的类型。可选值：PROMOCARD（优惠券）、PROMOCODE（优惠码）
	CouponContent     string `json:"coupon_content"`     //卡券内容。当卡券类型为优惠码时，值为优惠码字符串
	CouponDescription string `json:"coupon_description"` //卡券的说明
	CouponCondition   string `json:"coupon_condition"`   //卡券使用条件说明
	UsedAt            string `json:"used_at"`            //使用时间
	DiscountFee       json.Number `json:"discount_fee"`  //优惠的金额，单位：元，精确到小数点后两位

}

//订单中使用到的优惠活动的数据结构
type Promotion struct {
	PromotionID        int64 `json:"promotion_id"`         //该优惠活动的ID
	PromotionName      string `json:"promotion_name"`      //该优惠活动的名称
	PromotionType      string `json:"promotion_type"`      //优惠的类型。可选值：FULLREDUCE（满减满送）
	PromotionCondition string `json:"promotion_condition"` //优惠使用条件说明
	UsedAt             string `json:"used_at"`             //使用时间
	DiscountFee        json.Number `json:"discount_fee"`   //优惠的金额，单位：元，精确到小数点后两位

}

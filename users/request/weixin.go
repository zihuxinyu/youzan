package request
import "github.com/zihuxinyu/youzan"

//根据微信粉丝用户的 openid 或 user_id 获取用户信息
type WeixinFollowerGet struct {
	youzan.CommonHeader
	Fields       string `json:"fields"`        //需要返回的微信粉丝对象字段，如nick,avatar等。可选值：CrmWeixinFans微信粉丝结构体中所有字段均可返回；多个字段用“,”分隔。如果为空则返回所有
	WeixinOpenid string `json:"weixin_openid"` //微信粉丝用户的openid
	UserId       string `json:"user_id"`       //微信粉丝用户ID。
                                               //	调用时，参数 weixin_openid 和 user_id 选其一即可
}

//根据多个微信粉丝用户的 openid 或 user_id 获取用户信息
type WeixinFollowerGets struct {
	youzan.CommonHeader

	Fields        string `json:"fields"`         //需要返回的微信粉丝对象字段，如nick,avatar等。
                                                 // 可选值：CrmWeixinFans微信粉丝结构体中所有字段均可返回；
                                                 // 多个字段用“,”分隔。如果为空则返回所有
	WeixinOpenids string `json:"weixin_openids"` //微信粉丝用户的openid，多个用逗号分隔
	UserIds       string `json:"user_ids"`       //微信粉丝用户ID，多个用逗号分隔。
                                                 //	调用时，参数 weixin_openids 和 user_ids 选其一即可
}

//查询微信粉丝用户信息
type WeixinFollowersGet struct {
	youzan.CommonHeader

	Fields      string  `json:"fields,omitempty"`       //需要返回的微信粉丝对象字段，如nick,avatar等。
                                                        // 可选值：CrmWeixinFans微信粉丝结构体中所有字段均可返回；
                                                        // 多个字段用“,”分隔。如果为空则返回所有
	StartFollow string  `json:"start_follow,omitempty"` //关注的起始时间。查询在该时间之后（包含该时间）关注的粉丝，为空则不限制
	EndFollow   string  `json:"end_follow,omitempty"`   //关注的结束时间。查询在该时间之前关注的粉丝，为空则不限制
	PageNo      int64  `json:"page_no,omitempty"`       //页码
	PageSize    int64  `json:"page_size,omitempty"`     //每页条数，最大值：5000
}

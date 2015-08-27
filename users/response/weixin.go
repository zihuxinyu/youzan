package response
import "encoding/json"

//微信粉丝用户数据结构
type CrmWeixinFans struct {
	UserId       json.Number   `json:"user_id"`        //微信粉丝用户ID
	WeixinOpenid string `json:"weixin_openid"` //微信粉丝用户的openid，可用于微信Api
	Nick         string `json:"nick"`          //微信粉丝用户的昵称
	Avatar       string `json:"avatar"`        //微信粉丝用户的头像Url
	FollowTime   string `json:"follow_time"`   //关注时间
	Sex          string `json:"sex"`           //性别。可选值：m(男)，f(女)。未知则为空
	Province     string `json:"province"`      //所在身份
	City         string `json:"city"`          //所在城市
	Tags         []CrmUserTag `json:"tags"`    //粉丝标签列表
	UnionId      string `json:"union_id"`      //公众号绑定到微信开放平台帐号后，粉丝的唯一ID
}


//用户标签
type CrmUserTag struct {
	id   string `json:"id"`   //标签ID
	name string `json:"name"` //标签名
}


type CrmWeixinFansList struct {

	TotalResults json.Number `json:"total_results"`
	Users        []CrmWeixinFans `json:"users"`

}

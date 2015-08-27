package Items
import (
	"net/http"

	"github.com/zihuxinyu/youzan"
	"github.com/zihuxinyu/youzan/users/request"
	"github.com/zihuxinyu/youzan/users/response"
)
const (
	MethodWeixinFollowerGet  string = "kdt.users.weixin.follower.get"// 根据微信粉丝用户的 openid 或 user_id 获取用户信息
	MethodWeixinFollowerGets  string = "kdt.users.weixin.follower.gets"// 根据多个微信粉丝用户的 openid 或 user_id 获取用户信息
	MethodWeixinFollowersGet  string = "kdt.users.weixin.followers.get"//  查询微信粉丝用户信息
)

type Client youzan.Client

func NewClient(appId, appSecret string, clt *http.Client) *Client {
	return (*Client)(youzan.NewClient(appId, appSecret, clt))
}

//根据微信粉丝用户的 openid 或 user_id 获取用户信息
func (clt *Client)  WxFollowerGet(req *request.WeixinFollowerGet) (resp response.CrmWeixinFans, err error) {

	if req.Method == "" {
		req.Method = MethodWeixinFollowerGet
	}

	type result struct {
		Response struct {
			         User response.CrmWeixinFans `json:"user"`
		         } `json:"response"`
		youzan.Error
	}


	res := new(result)

	err = ((*youzan.Client)(clt)).Post(req, &res)
	if err != nil {
		return
	}
	if res.ErrorResponse.Code != youzan.ErrCodeOK {
		err = &res.Error
	}

	resp = res.Response.User

	return
}

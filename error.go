package youzan

import "fmt"

const (
	ErrCodeOK = 0  //无返回code时，匹配为0值
)


type Error struct {
	ErrorResponse struct {
		              Code int `json:"code"`
		              Msg  string `json:"msg"`
	              } `json:"error_response"`
}



func (e *Error) Error() string {
	var extra string
	switch e.ErrorResponse.Code {
	case -1:extra = "系统错误。系统内部错误，请直接联系技术支持，或邮件给 openapi@youzan.com。"
	case 40001:extra = "未指定 AppId。请求时传入 AppId。"
	case 40002:extra = "无效的App。申请有效的 AppId。"
	case 40003:extra = "无效的时间参数。以当前时间重新发起请求；如果系统时间和服务器时间误差超过10分钟，请调整系统时间。"
	case 40004:extra = "请求没有签名。请使用协议规范对请求中的参数进行签名。"
	case 40005:extra = "签名校验失败。检查 AppId 和 AppSecret 是否正确；如果是自行开发的协议分装，请检查代码。"
	case 40006:extra = "未指定请求的 Api 方法。指定 Api 方法。"
	case 40007:extra = "请求非法的方法。检查请求的方法的值。"
	case 40008:extra = "校验团队信息失败。检查团队是否有效、是否绑定微信。"
	case 40009:extra = "未指定 AccessToken。请求时传入 AccessToken。"
	case 40010:extra = "AccessToken不存在或已过期。传入有效的 AccessToken。"
	case 40011:extra = "无效的 AccessToken。传入有效的 AccessToken。"
	case 40012:extra = "请求的 Api 方法不在访问范围内。检查请求的方法。"
	case 41000:extra = "请求方法的应用级输入参数错误。阅读接口文档，检查调用接口时是否缺少必传参数。"
	case 50000:extra = "请求方法时业务逻辑发生错误。阅读返回的 error_response 里的 msg 字段，做相应的逻辑调整。"
	default:
		extra="未知错误。"
	}
	return fmt.Sprintf("错误代码: %d, 错误信息: %s, 描述: %s", e.ErrorResponse.Code, e.ErrorResponse.Msg, extra)
}

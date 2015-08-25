package sdk

const (
	ErrCodeOK                      = 0
	ErrCodeAccessTokenExpired      = 42001 // access_token 过期(无效)返回这个错误
	ErrCodeSuiteAccessTokenExpired = 42009 // suite_access_token 过期(无效)返回这个错误
//-1 	    系统错误 	系统内部错误，请直接联系技术支持，或邮件给 openapi@youzan.com
//40001 	未指定 AppId 	请求时传入 AppId
//40002 	无效的App 	申请有效的 AppId
//40003 	无效的时间参数 	以当前时间重新发起请求；如果系统时间和服务器时间误差超过10分钟，请调整系统时间
//40004 	请求没有签名 	请使用协议规范对请求中的参数进行签名
//40005 	签名校验失败 	检查 AppId 和 AppSecret 是否正确；如果是自行开发的协议分装，请检查代码
//40006 	未指定请求的 Api 方法 	指定 Api 方法
//40007 	请求非法的方法 	检查请求的方法的值
//40008 	校验团队信息失败 	检查团队是否有效、是否绑定微信
//40009 	未指定 AccessToken 	请求时传入 AccessToken
//40010 	AccessToken不存在或已过期 	传入有效的 AccessToken
//40011 	无效的 AccessToken 	传入有效的 AccessToken
//40012 	请求的 Api 方法不在访问范围内 	检查请求的方法
//41000 	请求方法的应用级输入参数错误 	阅读接口文档，检查调用接口时是否缺少必传参数
//50000 	请求方法时业务逻辑发生错误 	阅读返回的 error_response 里的 msg 字段，做相应的逻辑调整

)

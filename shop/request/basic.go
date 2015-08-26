package request
const (
	MethodBasicGet string = "kdt.shop.basic.get" //获取店铺基本信息
)
type Basic struct {
	Method string `json:"method"`
}

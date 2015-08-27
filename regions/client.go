package regions
//工具接口
import (
	"net/http"

	"github.com/zihuxinyu/youzan"
	"github.com/zihuxinyu/youzan/regions/request"
	"github.com/zihuxinyu/youzan/regions/response"
)
const (
	MethodRegionsGet string = "kdt.regions.get" //获取区域地名列表信息
)

type Client youzan.Client

func NewClient(appId, appSecret string, clt *http.Client) *Client {
	return (*Client)(youzan.NewClient(appId, appSecret, clt))
}

//获取区域地名列表信息
func (clt *Client)  Get(req *request.Region) (resp response.Regions, err error) {

	if req.Method == "" {
		req.Method = MethodRegionsGet
	}

	type result struct {
		response.Regions  `json:"response"`
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

	resp = res.Regions

	return
}

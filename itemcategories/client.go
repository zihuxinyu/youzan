package itemcategories
import (
	"net/http"

	"github.com/zihuxinyu/youzan"


	"github.com/zihuxinyu/youzan/itemcategories/response"
	"github.com/zihuxinyu/youzan/itemcategories/request"
)

const (
	MethodGet string = "kdt.itemcategories.get"//  获取商品分类二维列表
	MethodPromotionsGet  string = "kdt.itemcategories.promotions.get"// 获取商品推广栏目列表
	MethodTagsGet string = "kdt.itemcategories.tags.get"//获取商品自定义标签列表
	MethodTagsGetPage string = "kdt.itemcategories.tags.getpage"// 分页获取商品自定义标签列表
)
type Client youzan.Client

func NewClient(appId, appSecret string, clt *http.Client) *Client {
	return (*Client)(youzan.NewClient(appId, appSecret, clt))
}



//分页获取商品自定义标签列表
func (clt *Client)  TagsGetPage(req *request.TagsGetPage) (resp response.GoodsTagList, err error) {

	if req.Method == "" {
		req.Method = MethodTagsGetPage
	}


	type result struct {
		Response response.GoodsTagList `json:"response"`
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

	resp = res.Response

	return
}


//获取商品自定义标签列表
func (clt *Client)  TagsGet(req *request.TagsGet) (resp []response.GoodsTag, err error) {

	if req.Method == "" {
		req.Method = MethodTagsGet
	}


	type result struct {
		Response struct {
			         Tags []response.GoodsTag `json:"tags"`
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

	resp = res.Response.Tags

	return
}


//获取商品推广栏目列表
func (clt *Client)  PromotionsGet() (resp []response.GoodsPromotionCategory, err error) {


	type itemCategoriesPromotionsGet struct {
		youzan.CommonHeader
	}
	req := new(itemCategoriesPromotionsGet)

	req.Method = MethodPromotionsGet



	type result struct {
		Response struct {
			         Categories []response.GoodsPromotionCategory `json:"categories"`
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

	resp = res.Response.Categories

	return
}


//获取商品分类二维列表
func (clt *Client)  Get() (resp []response.GoodsCategory, err error) {

	type itemCategoriesGet struct {
		youzan.CommonHeader
	}
	req := new(itemCategoriesGet)
	req.Method = MethodGet

	type result struct {
		Response struct {
			         Categories []response.GoodsCategory `json:"categories"`
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

	resp = res.Response.Categories

	return
}


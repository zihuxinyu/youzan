package items
import (
	"net/http"

	"github.com/zihuxinyu/youzan"
	"github.com/zihuxinyu/youzan/items/request"
	"github.com/zihuxinyu/youzan/items/response"
	"errors"
)
const (
	MethodSkuUpdate string = "kdt.item.sku.update"// 更新SKU信息
	MethodSkuCustomGet string = "kdt.skus.custom.get"// 根据外部编号取商品Sku
	MethodOnSaleGet string = "kdt.items.onsale.get"// 获取出售中的商品列表
	MethodInventoryGet string = "kdt.items.inventory.get"// 获取仓库中的商品列表
	MethodCustomGet string = "kdt.items.custom.get"// 根据商品货号获取商品
	MethodUpdateListing string = "kdt.item.update.listing"// 商品上架
	MethodUpdateDelisting string = "kdt.item.update.delisting"// 商品上架
	MethodUpdate string = "kdt.item.update"// 更新单个商品信息
	MethodGet string = "kdt.item.get"// 得到单个商品信息
	MethodDelete string = "kdt.item.delete"// 删除一个商品
	MethodAdd string = "kdt.item.add"// 新增一个商品
)

type Client youzan.Client

func NewClient(appId, appSecret string, clt *http.Client) *Client {
	return (*Client)(youzan.NewClient(appId, appSecret, clt))
}

//更新SKU信息
func (clt *Client)  SkuUpdate(req *request.SkuUpdate) (resp response.GoodsSku, err error) {

	if req.Method == "" {
		req.Method = MethodSkuUpdate
	}

	type result struct {
		Response struct {
			         Sku response.GoodsSku `json:"sku"`
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

	resp = res.Response.Sku

	return
}


//根据外部编号取商品Sku
func (clt *Client)  SkuCustomGet(req *request.SkuCustomGet) (resp []response.GoodsSku, err error) {

	if req.Method == "" {
		req.Method = MethodSkuCustomGet
	}

	type result struct {
		Response struct {
			         Skus []response.GoodsSku `json:"skus"`
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

	resp = res.Response.Skus

	return
}

// 获取出售中的商品列表
func (clt *Client)  OnSaleGet(req *request.OnSaleGet) (resp response.GoodsList, err error) {

	if req.Method == "" {
		req.Method = MethodOnSaleGet
	}

	type result struct {
		Response response.GoodsList `json:"response"`
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


// 获取仓库中的商品列表
func (clt *Client)  InventoryGet(req *request.InventoryGet) (resp response.GoodsList, err error) {

	if req.Method == "" {
		req.Method = MethodInventoryGet
	}

	type result struct {
		Response response.GoodsList `json:"response"`
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

// 根据商品货号获取商品
func (clt *Client)  CustomGet(req *request.CustomGet) (resp []response.GoodsDetail, err error) {

	if req.Method == "" {
		req.Method = MethodCustomGet
	}

	type result struct {
		Response struct {
			         Items [] response.GoodsDetail `json:"items"`
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

	resp = res.Response.Items

	return
}

// 商品上架
func (clt *Client) UpdateListing(req *request.UpdateListing) (resp response.GoodsDetail, err error) {

	if req.Method == "" {
		req.Method = MethodUpdateListing
	}

	type result struct {
		Response struct {
			         Item response.GoodsDetail `json:"item"`
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

	resp = res.Response.Item

	return
}

//商品下架
func (clt *Client) UpdateDelisting(req *request.UpdateDelisting) (resp response.GoodsDetail, err error) {

	if req.Method == "" {
		req.Method = MethodUpdateDelisting
	}

	type result struct {
		Response struct {
			         Item response.GoodsDetail `json:"item"`
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

	resp = res.Response.Item

	return
}

// 更新单个商品信息
func (clt *Client) Update(req *request.ItemUpdate) (resp response.GoodsDetail, err error) {

	if req.Method == "" {
		req.Method = MethodUpdate
	}

	type result struct {
		Response struct {
			         Item response.GoodsDetail `json:"item"`
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

	resp = res.Response.Item

	return
}

// 得到单个商品信息
func (clt *Client) Get(req *request.ItemGet) (resp response.GoodsDetail, err error) {

	if req.Method == "" {
		req.Method = MethodGet
	}

	type result struct {
		Response struct {
			         Item response.GoodsDetail `json:"item"`
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

	resp = res.Response.Item

	return
}

//删除单个商品
func (clt *Client)  Delete(num_iid string) (isSuccess bool, err error) {

	if len(num_iid) == 0 {
		err = errors.New("请输入商品数字编号")
		return
	}
	type OnlineMarkSign struct {
		youzan.CommonHeader
		NumIid string `json:"num_iid"`
	}
	req := new(OnlineMarkSign)

	req.Method = MethodDelete
	req.NumIid = num_iid


	type result struct {
		Response struct {
			         IsSuccess bool `json:"is_success"`
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

	isSuccess = res.Response.IsSuccess

	return
}


//添加商品
func (clt *Client) Add(req *request.ItemAdd) (resp response.GoodsDetail, err error) {

	if req.Method == "" {
		req.Method = MethodAdd
	}

	type result struct {
		Response struct {
			         Item response.GoodsDetail `json:"item"`
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

	resp = res.Response.Item

	return
}

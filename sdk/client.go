package sdk
import (
	"net/http"
	"bytes"
	"sync"
	"encoding/json"
	"net/url"
	"fmt"
	"reflect"
	"github.com/astaxie/beego"
	"github.com/zihuxinyu/GoLibrary"
	"sort"
	"strconv"
	"crypto/md5"
	"encoding/hex"
)


var textBufferPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 16 << 10)) // 16KB
	},
}

var mediaBufferPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 10 << 20)) // 10MB
	},
}


//对请求进行签名

var (
//响应地址
	apiEntry string = "https://open.koudaitong.com/api/entry"

	format string = "json" //指定响应格式
	version string = "1.0" //API协议版本
	signMethod string = "md5" //参数的加密方法
)

type Client struct {
	AppId      string
	AppSecret  string
	HttpClient *http.Client
}

// 创建一个新的 Client.
//  如果 clt == nil 则默认用 http.DefaultClient
func NewClient(appId, appSecret string, clt *http.Client) *Client {
	if appId == "" || appSecret == "" {
		panic("nil appId or appSecret")
	}
	if clt == nil {
		clt = http.DefaultClient
	}

	return &Client{
		AppId               :appId,
		AppSecret           :appSecret,
		HttpClient          :clt,
	}
}

//组合创建的url
func (clt *Client) buildRequestParams(request interface{}) (url string) {


	params := map[string]string{}
	params["app_id"] = clt.AppId
	//组合最基本的
	params["timestamp"] = Library.TimeLocalString()
	params["format"] = "json"
	params["v"] = "1.0"
	params["signMethod"] = "md5"



	req := reflect.ValueOf(request).Elem()



	for x := 0; x < req.NumField(); x++ {

		switch req.Kind() {
		case reflect.String:
			params[req.Type().Field(x).Name] = req.Field(x).Interface().(string)
		case reflect.Int:
			params[req.Type().Field(x).Name]=strconv.Itoa(req.Field(x).Interface().(int))
		case reflect.Int64:
			params[req.Type().Field(x).Name]=strconv.FormatInt(req.Field(x).Interface().(int64),64)
		}
		beego.Error(req.Field(x).Type())
	}



	keys := sort.StringSlice{}
	for k, v := range params {
		keys = append(keys, k)
		beego.Error(k, v)
	}


	keys.Sort()

	lins := clt.AppSecret
	for k, v := range keys {
		lins = fmt.Sprintf("%s%s%s",lins, v, params[v])
		beego.Error(k, v, params[v])
	}
	lins = fmt.Sprintf("%s%s", lins, clt.AppSecret)


	h := md5.New()

	h.Write([]byte(lins))

	sign:= hex.EncodeToString(h.Sum(nil))

	url=fmt.Sprintf("%s?sign=%s",apiEntry,sign)
	for _, v := range keys {
		url = fmt.Sprintf("%s&%s=%s",url, v, params[v])
 	}


	beego.Error(url)



	return ""
}
func (clt *Client) PostJSON(request interface{}, response interface{}) (err error) {
	buf := textBufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer textBufferPool.Put(buf)

	if err = json.NewEncoder(buf).Encode(request); err != nil {
		return
	}
	requestBytes := buf.Bytes()



	finalURL := url.QueryEscape(clt.buildRequestParams(request))

	return
	httpResp, err := clt.HttpClient.Post(finalURL, "application/json; charset=utf-8", bytes.NewReader(requestBytes))
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", httpResp.Status)
	}

	//先判断是否error_response,是的话直接error
	//没有error那么解析到response

	if err = json.NewDecoder(httpResp.Body).Decode(response); err != nil {
		return
	}
	//
	//	var ErrorStructValue reflect.Value // Error
	//
	//	// 下面的代码对 response 有特定要求, 见此函数 NOTE
	//	responseStructValue := reflect.ValueOf(response).Elem()
	//	if v := responseStructValue.Field(0); v.Kind() == reflect.Struct {
	//		ErrorStructValue = v
	//	} else {
	//		ErrorStructValue = responseStructValue
	//	}

	return
}

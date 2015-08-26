package youzan
import (

	"time"
	"net/http"
	"encoding/json"
	"strings"
//	"reflect"
//	"github.com/huandu/xstrings"
	"github.com/astaxie/beego"
	"strconv"
	"sort"
	"fmt"
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego/httplib"
	"crypto/tls"
	"reflect"
)


const (
	apiEntry string = "https://open.koudaitong.com/api/entry"//响应地址
	apiFormat string = "json" //指定响应格式
	apiVersion string = "1.0" //API协议版本
	apiSignMethod string = "md5" //参数的加密方法
)
type CommonHeader struct {
	Method string `json:"method"`
}


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


//查看tag是否指定忽略
func omitempty(tag string) (bool) {
	list := strings.Split(tag, ",")
	for _, v := range list {
		if v == "omitempty" {
			return true
		}
	}
	return false
}


func (clt *Client) Post(request interface{}, response interface{}) (err error) {



	params := map[string]string{}
	//组合签名需要的参数
	params["app_id"] = clt.AppId
	params["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	params["format"] = apiFormat
	params["v"] = apiVersion
	params["sign_method"] = apiSignMethod


	//先对输入进行mashral，过滤调空值
	bytesRequest, err := json.Marshal(request)
	if err != nil {
		return
	}
	//赋值到map中
	var dat map[string]interface{}
	err = json.Unmarshal(bytesRequest, &dat)
	if err != nil {
		return
	}
	beego.Error(dat)
	for key, value := range dat {
		beego.Error(key, value, reflect.TypeOf(value))
		switch value.(type) {
		case string:
			params[key] = value.(string)
		case int:
			params[key] = strconv.Itoa(value.(int))
		case int64:
			params[key] = strconv.FormatInt(value.(int64), 10)
		case float64:
			params[key] = strconv.Itoa(int(value.(float64)))
		}
	}


	//键值按名称升序排列
	keys := sort.StringSlice{}
	for k, _ := range params {
		keys = append(keys, k)
	}
	keys.Sort()

	beego.Info(keys)
	//键值组合为字符串 将 secret 拼接到参数字符串头、尾
	linestr := clt.AppSecret
	for _, v := range keys {
		linestr = fmt.Sprintf("%s%s%s", linestr, v, params[v])
	}
	linestr = fmt.Sprintf("%s%s", linestr, clt.AppSecret)


	//计算字符串MD5,要求小写
	h := md5.New()
	h.Write([]byte(linestr))
	sign := hex.EncodeToString(h.Sum(nil))



	b := httplib.Post(apiEntry)
	b.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	//设置签名
	b.Param("sign", sign)
	for _, v := range keys {
		b.Param(v, params[v])
		beego.Info(v, params[v])
	}





//
//	HttpResp, err := b.Response()
//
//
//	if err != nil {
//		return
//	}
//	if err = json.NewDecoder(HttpResp.Body).Decode(response); err != nil {
//		return
//	}
	//
	bytestr, _ := b.Bytes()
	beego.Info(string(bytestr))

	err = json.Unmarshal(bytestr, response)
	if err != nil {
		return
	}


	return
}

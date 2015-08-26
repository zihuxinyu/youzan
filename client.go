package youzan
import (

	"fmt"
	"time"
	"sort"
	"reflect"
	"strconv"
	"net/http"
	"crypto/tls"
	"crypto/md5"
	"encoding/json"
	"encoding/hex"

	"github.com/huandu/xstrings"
	"github.com/astaxie/beego/httplib"

)


const (
	apiEntry string = "https://open.koudaitong.com/api/entry"//响应地址
	apiFormat string = "json" //指定响应格式
	apiVersion string = "1.0" //API协议版本
	apiSignMethod string = "md5" //参数的加密方法
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


func (clt *Client) Post(request interface{}, response interface{}) (err error) {



	params := map[string]string{}
	//组合签名需要的参数
	params["app_id"] = clt.AppId
	params["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	params["format"] = apiFormat
	params["v"] = apiVersion
	params["sign_method"] = apiSignMethod

	req := reflect.ValueOf(request).Elem()
	for x := 0; x < req.NumField(); x++ {

		//字段名称
		key := xstrings.ToSnakeCase(req.Type().Field(x).Name)
		//字段值
		value := req.Field(x).Interface()

		switch req.Field(x).Kind() {
		case reflect.String:
			params[key] = value.(string)
		case reflect.Int:
			params[key] = strconv.Itoa(value.(int))
		case reflect.Int64:
			params[key] = strconv.FormatInt(value.(int64), 10)
		}
	}



	//键值按名称升序排列
	keys := sort.StringSlice{}
	for k, _ := range params {
		keys = append(keys, k)
	}
	keys.Sort()

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
	}





	bytestr, _ := b.Bytes()
	err = json.Unmarshal(bytestr, response)
	if err != nil {
		return
	}



	return
}

package topsdk

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"

	"topsdk/httpclient"
)

const (
	//TaobaoOpenAPIURL 正式环境 HTTP请求地址
	TaobaoOpenAPIURL = "http://gw.api.taobao.com/router/rest"

	//TaobaoOpenAPIHttpsURL 正式环境 HTTPS请求地址 https://eco.taobao.com/router/rest
	TaobaoOpenAPIHttpsURL = "https://eco.taobao.com/router/rest"
)

//TopClient 阿里API Client
type TopClient struct {
	appKey    string
	appSecret string
}

//NewTopClient 返回新的 TopClient 指针
func NewTopClient(appKey, appSecret string) *TopClient {
	client := new(TopClient)
	client.appKey = appKey
	client.appSecret = appSecret
	return client
}

//UpdateKey 更新使用的KEY
func (m *TopClient) UpdateKey(newAppKey, newAppSecret string) {
	m.appKey = newAppKey
	m.appSecret = newAppSecret
}

//RequestWithKey 使用指定的 appKey appSecret 发送请求
func (m *TopClient) RequestWithKey(appKey, appSecret string, param ITaoBaoParam) (results map[string]interface{}, err error) {
	m.appKey = appKey
	m.appSecret = appSecret
	results, err = m.Request(param)
	return results, err
}

//Request 发送请求
func (m *TopClient) Request(param ITaoBaoParam) (results map[string]interface{}, err error) {
	var p = url.Values{}
	var keys = make([]string, 6, 6)

	p.Add("timestamp", time.Now().Format("2006-01-02 15:04:05"))
	p.Add("format", "json")
	p.Add("v", "2.0")
	p.Add("sign_method", "md5")
	p.Add("app_key", m.appKey)
	p.Add("method", param.APIName())

	keys[0] = "timestamp"
	keys[1] = "format"
	keys[2] = "v"
	keys[3] = "sign_method"
	keys[4] = "app_key"
	keys[5] = "method"

	if len(param.ExtJSONParamName()) > 0 {
		p.Add(param.ExtJSONParamName(), param.ExtJSONParamValue())
		keys = append(keys, param.ExtJSONParamName())
	}

	var ps = param.Params()
	if ps != nil {
		for key, value := range ps {
			p.Add(key, value)
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)
	p.Add("sign", m.sign(m.appSecret, keys, p))

	var req = httpclient.NewRequest("POST", TaobaoOpenAPIURL)
	req.SetHeader("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	req.SetParams(p)

	var rep = req.Exec()
	err = rep.Unmarshal(&results)
	return results, err
}

func (m *TopClient) sign(appSecret string, keys []string, param url.Values) (s string) {
	for _, key := range keys {
		s = s + key + param.Get(key)
	}
	s = fmt.Sprintf("%s%s%s", appSecret, s, appSecret)
	var md5 = md5.New()
	md5.Write([]byte(s))
	s = strings.ToUpper(hex.EncodeToString(md5.Sum(nil)))
	return s
}

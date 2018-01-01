package topsdk

import (
	"encoding/json"
)

//ITaoBaoParam API参数接口
type ITaoBaoParam interface {
	// 用于提供访问的 method
	APIName() string

	// 返回参数列表
	Params() map[string]string

	// 返回扩展 JSON 参数的字段名称
	ExtJSONParamName() string
	// 返回扩展 JSON 参数的字段值
	ExtJSONParamValue() string
}

// TaoBaoParam 示例，别无它用
type TaoBaoParam map[string]interface{}

//APIName 用于提供访问的 method
func (m TaoBaoParam) APIName() string {
	return "taobao.open.sms.sendmsg"
}

//Params 返回参数列表,根据API函数的不同需要修改
func (m TaoBaoParam) Params() map[string]string {

	return nil
}

// ExtJSONParamName 返回扩展 JSON 参数的字段名称
func (m TaoBaoParam) ExtJSONParamName() string {
	return "send_message_request"
}

// ExtJSONParamValue 返回扩展 JSON 参数的字段值
func (m TaoBaoParam) ExtJSONParamValue() string {
	var bytes, err = json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(bytes)
}

// AddParam 附加参数
func (m TaoBaoParam) AddParam(key string, value interface{}) {
	m[key] = value
}

func marshal(obj interface{}) string {
	var bytes, err = json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(bytes)
}

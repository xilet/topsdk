package topsdk

import (
	"strconv"
)

//TbkItemInfoGetParam taobao.tbk.item.info.get (淘宝客商品详情（简版）)
//http://open.taobao.com/doc2/apiDetail.htm?apiId=24518&scopeId=11655
//taobao.tbk.item.info.get (淘宝客商品详情（简版）)
type TbkItemInfoGetParam struct {
	//Fields fields	String	必须	num_iid,title,pict_url,small_images,reserve_price,zk_final_price,user_type,provcity,item_url		需返回的字段列表
	Fields string `json:"fields"`
	//NumIIDs num_iids	String	必须	123,456		商品ID串，用,分割，从taobao.tbk.item.get接口获取num_iid字段，最大40个
	NumIIDs string `json:"num_iids"`
	//platform	Number	可选	1	默认值：1 链接形式：1：PC，2：无线，默认：１
	Platform int `json:"platform,omitempty"`
}

//APIName API名称
func (m TbkItemInfoGetParam) APIName() string {
	return "taobao.tbk.item.info.get"
}

//Params 参数
func (m TbkItemInfoGetParam) Params() map[string]string {
	var params = make(map[string]string)
	if m.Platform == 0 {
		m.Platform = 1
	}
	params["platform"] = strconv.Itoa(m.Platform)
	params["fields"] = m.Fields
	params["num_iids"] = m.NumIIDs
	return params
}

//ExtJSONParamName JSON类型参数的名称
func (m TbkItemInfoGetParam) ExtJSONParamName() string {
	return "" //请求参数的名称
}

//ExtJSONParamValue JSON类型参数的值
func (m TbkItemInfoGetParam) ExtJSONParamValue() string {
	return marshal(m)
}

//AddParam 增加参数.暂时未用到
func (m *TbkItemInfoGetParam) AddParam(key string, value interface{}) {

}

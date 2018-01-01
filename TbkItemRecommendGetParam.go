package topsdk

import (
	"strconv"
)

//TbkItemRecommendGetParam taobao.tbk.item.recommend.get (淘宝客商品关联推荐查询)
//http://open.taobao.com/docs/api.htm?spm=a219a.7395905.0.0.ReBiJd&scopeId=11655&apiId=24517
type TbkItemRecommendGetParam struct {
	//Fields fields	String	必须	num_iid,title,pict_url,small_images,reserve_price,zk_final_price,user_type,provcity,item_url		需返回的字段列表
	Fields string `json:"fields"`
	//NumIID num_iid	Number	必须	123		商品Id
	NumIID int64 `json:"num_iid"`
	//Count count	Number	可选	20	默认值：20 返回数量，默认20，最大值40
	Count int64 `json:"count"`
	//platform	Number	可选	1	默认值：1 链接形式：1：PC，2：无线，默认：１
	Platform int `json:"platform,omitempty"`
}

//APIName taobao.tbk.item.recommend.get(淘宝客商品关联推荐查询)
func (m TbkItemRecommendGetParam) APIName() string {
	return "taobao.tbk.item.recommend.get"
}

//Params 参数
func (m TbkItemRecommendGetParam) Params() map[string]string {
	var params = make(map[string]string)
	if m.Platform == 0 {
		m.Platform = 1
	}
	if m.Count == 0 {
		m.Count = 20
	}
	params["count"] = strconv.FormatInt(m.Count, 10)
	params["platform"] = strconv.Itoa(m.Platform)
	params["fields"] = m.Fields
	params["num_iid"] = strconv.FormatInt(m.NumIID, 10)
	return params
}

//ExtJSONParamName JSON类型参数的名称
func (m TbkItemRecommendGetParam) ExtJSONParamName() string {
	return "" //请求参数的名称
}

//ExtJSONParamValue JSON类型参数的值
func (m TbkItemRecommendGetParam) ExtJSONParamValue() string {
	return marshal(m)
}

//AddParam 增加参数.暂时未用到
func (m *TbkItemRecommendGetParam) AddParam(key string, value interface{}) {

}

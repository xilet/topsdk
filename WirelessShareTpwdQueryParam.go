package topsdk

//WirelessShareTpwdQueryParam 查询淘口令信息
//http://open.taobao.com/doc2/apiDetail.htm?apiId=32461&scopeId=11998
//查询解析淘口令
type WirelessShareTpwdQueryParam struct {
	//PasswordContent password_content	String	必须	【天猫品牌号】，复制这条信息￥sMCl0Yra3Ae￥后打开手机淘宝		淘口令
	PasswordContent string `json:"password_content"`
}

//APIName API名称
func (m WirelessShareTpwdQueryParam) APIName() string {
	return "taobao.wireless.share.tpwd.query"
}

//Params 参数
func (m WirelessShareTpwdQueryParam) Params() map[string]string {
	var params = make(map[string]string)
	params["password_content"] = m.PasswordContent
	return params
}

//ExtJSONParamName JSON类型参数的名称
func (m WirelessShareTpwdQueryParam) ExtJSONParamName() string {
	return "" //请求参数的名称
}

//ExtJSONParamValue JSON类型参数的值
func (m WirelessShareTpwdQueryParam) ExtJSONParamValue() string {
	return marshal(m)
}

//AddParam 增加参数.暂时未用到
func (m *WirelessShareTpwdQueryParam) AddParam(key string, value interface{}) {

}

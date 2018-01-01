package topsdk

//TbkTpwdCreateParam 生成淘口令
//http://open.taobao.com/doc2/apiDetail.htm?spm=0.0.0.0.fnw4Cl&apiId=25598&docType=
//提供淘客生成淘口令接口，淘客提交口令内容、logo、url等参数，生成淘口令关键key如：￥SADadW￥，后续进行文案包装组装用于传播
type TbkTpwdCreateParam struct {
	//UserID user_id	String	可选	123		生成口令的淘宝用户ID
	UserID string `json:"user_id,omitempty"`

	//Text text 口令弹框内容 必须 长度大于5个字符
	Text string `json:"text"`
	//URL url	String	必须	https://uland.taobao.com/		口令跳转目标页
	URL string `json:"url"`

	//Logo logo	String	可选	https://uland.taobao.com/		口令弹框logoURL
	Logo string `json:"logo,omitempty"`
	//ext	String	可选	{}		扩展字段JSON格式
	Ext string `json:"ext,omitempty"`
}

//APIName API名称
func (m TbkTpwdCreateParam) APIName() string {
	return "taobao.tbk.tpwd.create"
}

//Params 参数
func (m TbkTpwdCreateParam) Params() map[string]string {
	var params = make(map[string]string)
	if len(m.UserID) > 0 {
		params["user_id"] = m.UserID
	}
	if len(m.Logo) > 0 {
		params["logo"] = m.Logo
	}
	if len(m.Ext) > 0 {
		params["ext"] = m.Ext
	}
	params["text"] = m.Text
	params["url"] = m.URL
	return params
}

//ExtJSONParamName JSON类型参数的名称
func (m TbkTpwdCreateParam) ExtJSONParamName() string {
	return "" //请求参数的名称
}

//ExtJSONParamValue JSON类型参数的值
func (m TbkTpwdCreateParam) ExtJSONParamValue() string {
	return marshal(m)
}

//AddParam 增加参数.暂时未用到
func (m *TbkTpwdCreateParam) AddParam(key string, value interface{}) {

}

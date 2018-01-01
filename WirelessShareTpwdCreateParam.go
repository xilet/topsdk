package topsdk

//WirelessShareTpwdCreateParam taobao.wireless.share.tpwd.create (生成淘口令)
//http://open.taobao.com/doc2/apiDetail.htm?apiId=26520&scopeId=11998
//提供isv生成淘口令接口，isv提交口令内容、logo、url等参数，生成淘口令关键key如：￥SADadW￥，后续进行文案包装组装用于传播
type WirelessShareTpwdCreateParam struct {
	//UserID user_id Number 可选 24234234234 默认值：0 生成口令的淘宝用户ID
	UserID int64 `json:"user_id,omitempty"`

	//Text text String 必须 超值活动，惊喜活动多多口令弹框内容
	Text string `json:"text"`
	//URL url String 必须 http://m.taobao.com口令跳转url
	URL string `json:"url"`

	//Logo logo String 可选 http://m.taobao.com/xxx.jpg口令弹框logoURL
	Logo string `json:"logo,omitempty"`
	//Ext ext String 可选 {"xx":"xx"}扩展字段JSON格式
	Ext string `json:"ext,omitempty"`
}

//APIName API名称
func (m WirelessShareTpwdCreateParam) APIName() string {
	return "taobao.wireless.share.tpwd.create"
}

//Params 参数
func (m WirelessShareTpwdCreateParam) Params() map[string]string {
	return nil
}

//ExtJSONParamName JSON类型参数的名称
func (m WirelessShareTpwdCreateParam) ExtJSONParamName() string {
	return "tpwd_param" //请求参数的名称
}

//ExtJSONParamValue JSON类型参数的值
func (m WirelessShareTpwdCreateParam) ExtJSONParamValue() string {
	return marshal(m)
}

//AddParam 增加参数.暂时未用到
func (m *WirelessShareTpwdCreateParam) AddParam(key string, value interface{}) {

}

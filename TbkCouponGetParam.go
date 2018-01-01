package topsdk

//TbkCouponGetParam taobao.tbk.coupon.get (阿里妈妈推广券信息查询)
//http://open.taobao.com/doc2/apiDetail.htm?apiId=31106&scopeId=11655
type TbkCouponGetParam struct {
	//Me me	String	必须	nfr%2BYTo2k1PX18gaNN%2BIPkIG2PadNYbBnwEsv6mRavWieOoOE3L9OdmbDSSyHbGxBAXjHpLKvZbL1320ML%2BCF5FRtW7N7yJ056Lgym4X01A%3D		带券ID与商品ID的加密串
	Me string `json:"me"`
}

//APIName API名称
func (m TbkCouponGetParam) APIName() string {
	return "taobao.tbk.coupon.get"
}

//Params 参数
func (m TbkCouponGetParam) Params() map[string]string {
	var params = make(map[string]string)
	params["me"] = m.Me
	return params
}

//ExtJSONParamName JSON类型参数的名称
func (m TbkCouponGetParam) ExtJSONParamName() string {
	return "" //请求参数的名称
}

//ExtJSONParamValue JSON类型参数的值
func (m TbkCouponGetParam) ExtJSONParamValue() string {
	return marshal(m)
}

//AddParam 增加参数.暂时未用到
func (m *TbkCouponGetParam) AddParam(key string, value interface{}) {

}

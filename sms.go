package topsdk

/*
////////////////////////////////////////////////////////////////////////////////
// OpenSMSSendMsgParam 向手机发送短信
// http://open.taobao.com/doc2/apiDetail.htm?spm=0.0.0.0.fnw4Cl&apiId=25598&docType=
type OpenSMSSendMsgParam struct {
	// 必须参数
	TemplateId string                 `json:"template_id,omitempty"` // 必须 模板id
	Context    map[string]interface{} `json:"context"`               // 必须 模板上下文
	Mobile     string                 `json:"mobile"`                // 必须 手机号

	// 可选参数
	SignatureId        string `json:"signature_id,omitempty"`          // 可选 签名id
	ExternalId         string `json:"external_id,omitempty"`           // 可选 外部id
	DeviceLimit        int    `json:"external_id,omitempty"`           // 可选 设备级别次数限制
	SessionLimit       int    `json:"session_limit,omitempty"`         // 可选 session级别次数限制
	DeviceLimitInTime  int    `json:"device_limit_in_time,omitempty"`  // 可选 时间，单位秒
	MobileLimit        int    `json:"mobile_limit,omitempty"`          // 可选 手机号限制
	SessionLimitInTime int    `json:"session_limit_in_time,omitempty"` // 可选 时间，单位秒
	MobileLimitInTime  int    `json:"mobile_limit_in_time,omitempty"`  // 可选 时间，单位秒
	SessionId          string `json:"session_id,omitempty"`            // 可选 sessionId
	Domain             string `json:"domain,omitempty"`                // 可选 业务域
	DeviceId           string `json:"device_id,omitempty"`             // 可选 设备id
}

func (this OpenSMSSendMsgParam) APIName() string {
	return "taobao.open.sms.sendmsg"
}

func (this OpenSMSSendMsgParam) Params() map[string]string {
	return nil
}

func (this OpenSMSSendMsgParam) ExtJSONParamName() string {
	return "send_message_request"
}

func (this OpenSMSSendMsgParam) ExtJSONParamValue() string {
	return marshal(this)
}

func (this *OpenSMSSendMsgParam) AddParam(key string, value interface{}) {
	if this.Context == nil {
		this.Context = make(map[string]interface{})
	}
	this.Context[key] = value
}
*/

package httpclient

import (
	"encoding/json"
	"net/http"
)

//Response Http请求返回内容
type Response struct {
	*http.Response
	data  []byte
	error error
}

func (m *Response) Error() error {
	return m.error
}

//Bytes 返回byte数据
func (m *Response) Bytes() ([]byte, error) {
	return m.data, m.error
}

//MustBytes 返回byte数据
func (m *Response) MustBytes() []byte {
	return m.data
}

//String 返回string数据
func (m *Response) String() (string, error) {
	return string(m.data), m.error
}

//MustString 返回string数据
func (m *Response) MustString() string {
	return string(m.data)
}

//Unmarshal 返回反序列化JSON
func (m *Response) Unmarshal(v interface{}) error {
	if m.error != nil {
		return m.error
	}
	return json.Unmarshal(m.data, v)
}

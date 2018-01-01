package httpclient

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

//Request Http Request
type Request struct {
	url     string
	method  string
	header  http.Header
	params  url.Values
	body    io.Reader
	Client  *http.Client
	cookies []*http.Cookie
	file    *file
}

type file struct {
	name     string
	filename string
	path     string
}

//NewRequest 新的Request指针
func NewRequest(method, urlString string) *Request {
	var r = &Request{}
	r.method = strings.ToUpper(method)
	r.url = urlString
	r.params = url.Values{}
	r.header = http.Header{}
	r.Client = http.DefaultClient
	r.SetContentType("application/x-www-form-urlencoded")
	return r
}

//SetContentType 设定Content-Type
func (m *Request) SetContentType(contentType string) {
	m.SetHeader("Content-Type", contentType)
}

//AddHeader 增加Header头
func (m *Request) AddHeader(key, value string) {
	m.header.Add(key, value)
}

//SetHeader 设定Header头
func (m *Request) SetHeader(key, value string) {
	m.header.Set(key, value)
}

//SetHeaders 设定Header头
func (m *Request) SetHeaders(header http.Header) {
	m.header = header
}

//SetBody 设定POST内容
func (m *Request) SetBody(body io.Reader) {
	m.body = body
	m.params = nil
}

//AddParam 增加Get请求参数
func (m *Request) AddParam(key, value string) {
	m.params.Add(key, value)
	m.body = nil
}

//SetParam 设定Get请求参数
func (m *Request) SetParam(key, value string) {
	m.params.Set(key, value)
	m.body = nil
}

//SetParams 设定Get请求参数
func (m *Request) SetParams(params url.Values) {
	m.params = params
}

//AddFile 增加文件
func (m *Request) AddFile(name, filename, path string) {
	m.file = &file{name, filename, path}
}

//RemoveFile 移除文件
func (m *Request) RemoveFile() {
	m.file = nil
}

//AddCookie 添加COOKIE
func (m *Request) AddCookie(cookie *http.Cookie) {
	m.cookies = append(m.cookies, cookie)
}

//Exec 发送HTTP请求
func (m *Request) Exec() *Response {
	var req *http.Request
	var err error
	var body io.Reader
	var rawQuery string

	if m.method == http.MethodGet || m.method == http.MethodHead || m.method == http.MethodDelete {
		if len(m.params) > 0 {
			rawQuery = m.params.Encode()
		}
	} else {
		if m.body != nil {
			body = m.body
		} else if m.file != nil {
			uploadFile, err := os.Open(m.file.path)
			if err != nil {
				return &Response{nil, nil, err}
			}
			defer uploadFile.Close()

			bodyByte := &bytes.Buffer{}
			writer := multipart.NewWriter(bodyByte)
			part, err := writer.CreateFormFile(m.file.name, m.file.filename)
			if err != nil {
				return &Response{nil, nil, err}
			}
			_, err = io.Copy(part, uploadFile)
			if err != nil {
				return &Response{nil, nil, err}
			}

			for key, values := range m.params {
				for _, value := range values {
					writer.WriteField(key, value)
				}
			}

			err = writer.Close()
			if err != nil {
				return &Response{nil, nil, err}
			}

			m.SetContentType(writer.FormDataContentType())
			body = bodyByte
		} else if m.params != nil {
			body = strings.NewReader(m.params.Encode())
		}
	}

	req, err = http.NewRequest(m.method, m.url, body)
	if len(rawQuery) > 0 {
		req.URL.RawQuery = rawQuery
	}

	if err != nil {
		return &Response{nil, nil, err}
	}
	req.Header = m.header

	for _, cookie := range m.cookies {
		req.AddCookie(cookie)
	}

	resp, err := m.Client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return &Response{nil, nil, err}
	}

	data, err := ioutil.ReadAll(resp.Body)
	return &Response{resp, data, err}
}

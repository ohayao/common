package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	nh "net/http"
	"sync"
	"time"
)

// Context 请求上下文
// 默认超时时间 5分钟
type Context struct {
	URL     string
	request *nh.Request
	client  *nh.Client
	headers map[string]string
	cookies []*nh.Cookie
	Result  *response
	sync.Mutex
}
type response struct {
	Err    error
	Body   []byte
	Header nh.Header
}

// New 初始化
func New(url string) *Context {
	return &Context{
		URL:     url,
		headers: make(map[string]string),
		cookies: make([]*nh.Cookie, 0),
		client:  &nh.Client{Timeout: time.Minute * 5},
		Result:  new(response),
	}
}

// Header 添加请求头
func (that *Context) Header(key, value string) *Context {
	that.Lock()
	defer that.Unlock()
	that.headers[key] = value
	return that
}

// Cookie 添加Cookie
func (that *Context) Cookie(c *nh.Cookie) *Context {
	return that
}

// Timeout 设置超时时间
func (that *Context) Timeout(time time.Duration) *Context {
	that.client.Timeout = time
	return that
}

// createRequest 生成请求
func (that *Context) createRequest(method string, body *bytes.Buffer) *Context {
	if res, err := nh.NewRequest(method, that.URL, body); err == nil {
		that.request = res
	}
	for k, v := range that.headers {
		that.request.Header.Add(k, v)
	}
	for _, v := range that.cookies {
		that.request.AddCookie(v)
	}
	return that
}

// GetBytes 获取原始数据
func (that *Context) GetBytes() []byte {
	that.Lock()
	defer that.Unlock()
	that.createRequest("GET", bytes.NewBuffer(nil))
	return that.Do().Result.Body
}

// GetString 获取字符串
func (that *Context) GetString() string {
	that.GetBytes()
	if that.Result.Err == nil {
		return string(that.Result.Body)
	}
	return ""
}

// GetJSON 获取JSON对象
// v 指针结构体
func (that *Context) GetJSON(v interface{}) {
	that.GetBytes()
	if that.Result.Err == nil {
		_ = json.Unmarshal(that.Result.Body, v)
	}
}

// PostBytes 提交请求
func (that *Context) PostBytes(buf *bytes.Buffer) []byte {
	that.Lock()
	defer that.Unlock()
	that.createRequest("POST", buf)
	return that.Do().Result.Body
}

// PostString 获取字符串
func (that *Context) PostString(buf *bytes.Buffer) string {
	that.PostBytes(buf)
	if that.Result.Err == nil {
		return string(that.Result.Body)
	}
	return ""
}

// PostJSON 获取JSON对象
// v 指针结构体
func (that *Context) PostJSON(v interface{}, buf *bytes.Buffer) {
	that.PostBytes(buf)
	if that.Result.Err == nil {
		_ = json.Unmarshal(that.Result.Body, v)
	}
}

// Do 执行请求
func (that *Context) Do() *Context {
	res, err := that.client.Do(that.request)
	that.Result.Header = res.Header
	if err != nil {
		that.Result.Err = err
		that.Result.Body = nil
	} else {
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		that.Result.Err = err
		that.Result.Body = body
	}
	return that
}

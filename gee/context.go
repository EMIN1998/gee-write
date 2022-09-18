package gee

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Context struct {
	Writer     http.ResponseWriter
	Request    *http.Request
	Path       string
	Method     string
	StatusCode int
	Params     map[string]string
}

type H map[string]interface{}

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		Writer:  w,
		Request: r,
		Path:    r.URL.Path,
		Method:  r.Method,
	}
}

func (c *Context) PostForm(key string) string {
	return c.Request.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c *Context) ParseBody(dst interface{}) (err error) {
	if c.Request == nil {
		err := errors.New("request is empty")
		return err
	}

	if c.Request.Body == nil {
		err = fmt.Errorf(" r.Body为空.URL:%s", c.Request.RequestURI)
		return
	}

	postData, er := ioutil.ReadAll(c.Request.Body)
	if er != nil {
		err = fmt.Errorf(" ReadAll发生异常. URL:%s err:%s", c.Request.RequestURI, er)
		return
	}
	defer c.Request.Body.Close()

	err = json.Unmarshal(postData, &dst)
	if err != nil {
		return err
	}

	return
}

func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) String(code int, format string, value ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.StatusCode = code
	msg := fmt.Sprintf(format, value...)
	c.Writer.Write([]byte(msg))
}

func (c *Context) Json(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.StatusCode = code
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.StatusCode = code
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.StatusCode = code
	c.Writer.Write([]byte(html))
}

func (c *Context) GetParams(key string) string {
	if val, ok := c.Params[key]; ok {
		return val
	}

	return ""
}

package essen

import (
	"net/http"
	"net/url"
)

type GetBody struct {
	body *url.URL
}

type PostBody struct {
	body *http.Request
}

type Param interface {
	Params(name string) (string, EssenError)
}

func (b GetBody) Params(name string) (string, EssenError) {
	v := b.body.Query().Get(name)
	ee := EssenError{nilval: true}
	if v == "" {
		ee.nilval = false
		ee.errortype = "InvalidParam"
		ee.message = `No parameter with key "` + name + `"`
	}
	return v, ee
}

func (b PostBody) Params(name string) (string, EssenError) {
	err := b.body.ParseForm()
	ee := EssenError{nilval: true}
	if err != nil {
		ee.nilval = false
		ee.errortype = "FormParseError"
		ee.message = err.Error()
		return "", ee
	}
	v := b.body.PostFormValue(name)
	if v == "" {
		ee.nilval = false
		ee.errortype = "InvalidParam"
		ee.message = `No parameter with key "` + name + `"`
		return "", ee
	}
	return v, ee
}

func (r Request) Path() string {
	return r.Req.URL.Path
}

func (r Request) Host() string {
	return r.Req.URL.Host
}

func (r Request) Method() string {
	return r.Req.Method
}

func (r Request) HasHeader(key string) bool {
	v := r.Req.Header.Get(key)
	if v == "" {
		return false
	}
	return true
}
func (r Request) Header(key string) (string, EssenError) {
	if r.HasHeader(key) {
		return r.Req.Header.Get(key), EssenError{nilval: true}
	}
	return "", EssenError{message: "No Header Found", errortype: "NoHeader", nilval: false}
}

func (r *Request) requestBody() {
	if r.Method() == "GET" {
		r.Body = GetBody{body: r.Req.URL}
		return
	}
	if r.Method() == "POST" {
		r.Body = PostBody{body: r.Req}
		return
	}
}

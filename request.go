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
	Params(name string) (EssenError, string)
}

func (b GetBody) Params(name string) (EssenError, string) {
	v := b.body.Query().Get(name)
	ee := EssenError{nilval: true}
	if v == "" {
		ee.nilval = false
		ee.errortype = "InvalidParam"
		ee.message = `No parameter with key" ` + name + `"`
	}
	return ee, v
}

func (b PostBody) Params(name string) (EssenError, string) {
	err := b.body.ParseForm()
	ee := EssenError{nilval: true}
	if err != nil {
		ee.nilval = false
		ee.errortype = "FormParseError"
		ee.message = err.Error()
		return ee, ""
	}
	v := b.body.PostFormValue(name)
	if v == "" {
		ee.nilval = false
		ee.errortype = "InvalidParam"
		ee.message = `No parameter with key" ` + name + `"`
		return ee, ""
	}
	return ee, v
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

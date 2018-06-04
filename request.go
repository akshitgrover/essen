package essen

import (
	"log"
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
	Params(name string) string
}

func (b GetBody) Params(name string) string {
	v := b.body.Query().Get(name)
	return v
}

func (b PostBody) Params(name string) string {
	err := b.body.ParseForm()
	if err != nil {
		log.Panic(err)
	}
	return b.body.PostFormValue(name)
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

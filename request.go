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

package essen

import (
	"net/http"
)

type handlerStorage map[string]func(http.ResponseWriter, *http.Request)

type pathCache struct {
	post handlerStorage
	get  handlerStorage
	use  handlerStorage
}

var paths = pathCache{get: make(handlerStorage), post: make(handlerStorage)}

func (e Essen) Get(route string, f func(Response, Request)) {
	ff := func(res http.ResponseWriter, req *http.Request) {
		eres := Response{Res: res}
		ereq := Request{Req: req}
		f(eres, ereq)
	}
	paths.get[route] = ff
}

func (e Essen) Post(route string, f func(Response, Request)) {
	ff := func(res http.ResponseWriter, req *http.Request) {
		eres := Response{Res: res}
		ereq := Request{Req: req}
		f(eres, ereq)
	}
	paths.post[route] = ff
}

func (e Essen) Use(route string, f func(Response, Request)) {
	ff := func(res http.ResponseWriter, req *http.Request) {
		eres := Response{Res: res}
		ereq := Request{Req: req}
		f(eres, ereq)
	}
	paths.use[route] = ff
}

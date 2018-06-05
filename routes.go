package essen

import (
	"net/http"
)

type handlerStorage map[string]func(http.ResponseWriter, *http.Request)

type pathCache struct {
	post handlerStorage
	get  handlerStorage
	head handlerStorage
	use  handlerStorage
}

var paths = pathCache{get: make(handlerStorage), post: make(handlerStorage), head: make(handlerStorage)}

func (e Essen) Head(route string, f func(Response, Request)) {
	ff := func(res http.ResponseWriter, req *http.Request) {

		//Custom Response Fields
		eres := Response{Res: res, ReqMethod: "HEAD"}

		//Custom Request Fields
		ereq := Request{Req: req}
		ereq.requestBody()

		//Call Registered Middleware
		f(eres, ereq)
	}
	paths.head[route] = ff
}

func (e Essen) Get(route string, f func(Response, Request)) {
	ff := func(res http.ResponseWriter, req *http.Request) {

		//Custom Response Fields
		eres := Response{Res: res, ReqMethod: "GET"}

		//Custom Request Fields
		ereq := Request{Req: req}
		ereq.requestBody()

		//Call Registered Middleware
		f(eres, ereq)
	}
	paths.get[route] = ff
}

func (e Essen) Post(route string, f func(Response, Request)) {
	ff := func(res http.ResponseWriter, req *http.Request) {

		//Custom Response Fields
		eres := Response{Res: res, ReqMethod: "POST"}

		//Custom Request Fields
		ereq := Request{Req: req}
		ereq.requestBody()

		//Call Registered Middleware
		f(eres, ereq)
	}
	paths.post[route] = ff
}

func (e Essen) Use(route string, f func(Response, Request)) {
	ff := func(res http.ResponseWriter, req *http.Request) {
		eres := Response{Res: res, ReqMethod: "USE"}
		ereq := Request{Req: req}
		f(eres, ereq)
	}
	paths.use[route] = ff
}

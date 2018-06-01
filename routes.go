package essen

import (
	"net/http"
)

func (e Essen) Get(route string, f func(Response, Request)) {
	ff := func(res http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" && req.URL.Path != route {
			http.NotFound(res, req)
		}
		eres := Response{Res: res}
		ereq := Request{Req: req}
		f(eres, ereq)
	}
	http.HandleFunc(route, ff)
}

func (e Essen) Post(route string, f func(Response, Request)) {
	ff := func(res http.ResponseWriter, req *http.Request) {
		if req.Method != "POST" && req.URL.Path != route {
			http.NotFound(res, req)
		}
		eres := Response{Res: res}
		ereq := Request{Req: req}
		f(eres, ereq)
	}
	http.HandleFunc(route, ff)
}

func (e Essen) Use(route string, f func(Response, Request)) {
	ff := func(res http.ResponseWriter, req *http.Request) {
		if req.URL.Path != route {
			http.NotFound(res, req)
			return
		}
		eres := Response{Res: res}
		ereq := Request{Req: req}
		f(eres, ereq)
	}
	http.HandleFunc(route, ff)
}

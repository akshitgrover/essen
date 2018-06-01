package essen

import (
	"net/http"
	"strconv"
)

type Essen struct{}

type Response struct {
	Res http.ResponseWriter
}

type Request struct {
	Req *http.Request
}

func App() Essen {
	e := Essen{}
	return e
}

func (e Essen) Get(route string, f func(Response, Request)) {
	ff := func(res http.ResponseWriter, req *http.Request) {
		eres := Response{Res: res}
		ereq := Request{Req: req}
		if req.Method == "GET" {
			f(eres, ereq)
		} else {
			http.NotFound(res, req)
		}
	}
	http.HandleFunc(route, ff)
}

func (e Essen) Post(route string, f func(Response, Request)) {
	ff := func(res http.ResponseWriter, req *http.Request) {
		eres := Response{Res: res}
		ereq := Request{Req: req}
		if req.Method == "POST" {
			f(eres, ereq)
		} else {
			http.NotFound(res, req)
		}
	}
	http.HandleFunc(route, ff)
}

func (e Essen) Use(route string, f func(Response, Request)) {
	ff := func(res http.ResponseWriter, req *http.Request) {
		eres := Response{Res: res}
		ereq := Request{Req: req}
		f(eres, ereq)
	}
	http.HandleFunc(route, ff)
}

func (e Essen) Listen(port int) {
	fport := ":" + strconv.Itoa(port)
	http.ListenAndServe(fport, nil)
}

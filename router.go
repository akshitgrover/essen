package essen

import (
	"github.com/zemirco/uid"
	"net/http"
)

type router struct {
	get  handlerStorage
	post handlerStorage
	put  handlerStorage
}

func (r *router) Get(route string, f func(Response, Request)) {
	ff := func(res http.ResponseWriter, req *http.Request) {

		//Custom Response Fields
		eres := Response{Res: res, ReqMethod: "GET"}

		//Custom Request Fields
		ereq := Request{Req: req, Uid: uid.New(7)}
		ereq.requestBody()

		//Call Registered Middleware
		f(eres, ereq)
	}
	r.get[route] = ff
}

func (r *router) Post(route string, f func(Response, Request)) {
	ff := func(res http.ResponseWriter, req *http.Request) {

		//Custom Response Fields
		eres := Response{Res: res, ReqMethod: "POST"}

		//Custom Request Fields
		ereq := Request{Req: req, Uid: uid.New(7)}
		ereq.requestBody()

		//Call Registered Middleware
		f(eres, ereq)
	}
	r.post[route] = ff
}

func (r *router) Put(route string, f func(Response, Request)) {
	ff := func(res http.ResponseWriter, req *http.Request) {

		//Custom Response Fields
		eres := Response{Res: res, ReqMethod: "PUT"}

		//Csutom Request Fields
		ereq := Request{Req: req, Uid: uid.New(7)}
		ereq.requestBody()

		//Call Registered Middleware
		f(eres, ereq)
	}
	r.put[route] = ff
}

func (e Essen) UseRouter(route string, router router) {

	//Register Get Routes
	for k, v := range router.get {
		froute := route + k
		paths.get[froute] = v
	}

	//Register Post Routes
	for k, v := range router.post {
		froute := route + k
		paths.post[froute] = v
	}

	//Register Put Routes
	for k, v := range router.put {
		froute := route + k
		paths.put[froute] = v
	}
}

func (r *router) Done() {
	*r = router{}
}

func (e Essen) Router() router {
	return router{get: make(handlerStorage), post: make(handlerStorage)}
}

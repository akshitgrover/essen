package essen

import (
	"github.com/zemirco/uid"
	"log"
	"net/http"
	"path/filepath"
)

type handlerStorage map[string]func(http.ResponseWriter, *http.Request)
type staticHandlerStorage map[string]http.Handler

type pathCache struct {
	post   handlerStorage
	put    handlerStorage
	get    handlerStorage
	head   handlerStorage
	use    handlerStorage
	static staticHandlerStorage
}

var paths = pathCache{
	get:    make(handlerStorage),
	post:   make(handlerStorage),
	put:    make(handlerStorage),
	head:   make(handlerStorage),
	use:    make(handlerStorage),
	static: make(staticHandlerStorage),
}

func (e Essen) Head(route string, f func(Response, Request)) {
	ff := func(res http.ResponseWriter, req *http.Request) {

		//Custom Response Fields
		eres := Response{Res: res, ReqMethod: "HEAD"}

		//Custom Request Fields
		ereq := Request{Req: req, Uid: uid.New(7)}
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
		ereq := Request{Req: req, Uid: uid.New(7)}
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
		ereq := Request{Req: req, Uid: uid.New(7)}
		ereq.requestBody()

		//Call Registered Middleware
		f(eres, ereq)
	}
	paths.post[route] = ff
}

func (e Essen) Put(route string, f func(Response, Request)) {
	ff := func(res http.ResponseWriter, req *http.Request) {

		//Custom Response Fields
		eres := Response{Res: res, ReqMethod: "PUT"}

		//Custom Request Fields
		ereq := Request{Req: req, Uid: uid.New(7)}
		ereq.requestBody()

		//Call Registered Middleware
		f(eres, ereq)
	}
	paths.put[route] = ff
}

func (e Essen) Use(route string, f func(Response, Request)) {
	ff := func(res http.ResponseWriter, req *http.Request) {

		//Custom Response Field
		eres := Response{Res: res, ReqMethod: req.Method}

		//Custom Request Fields
		ereq := Request{Req: req, Uid: uid.New(7)}
		ereq.requestBody()

		//Call Registered Middleware
		f(eres, ereq)
	}
	paths.use[route] = ff
}

func (e Essen) Static(route string, path string) {

	//Create error instance
	ee := EssenError{nilval: true}

	//Generate absolute path
	path, err := filepath.Abs(path)

	//Check for absolute path generation error
	if err != nil {
		ee.nilval = false
		ee.errortype = "PathError"
		ee.message = "String provided cannot be used for absolute path conversion"
		log.Panic(ee.Error())
	}

	//Create file serving handler
	paths.static[route] = http.FileServer(http.Dir(path))
}

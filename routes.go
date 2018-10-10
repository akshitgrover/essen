package essen

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/zemirco/uid"
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

//Head is a function to register a route against HTTP HEAD request method
//
//  app := essen.App()
//
//  app.Head("/", func(res app.Response, req app.Request){
//  	//do something
//  })
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

//Get is a function to register a route against HTTP GET request method
//
//  app := essen.App()
//
//  app.Get("/", func(res app.Response, req app.Request){
//  	//do something
//  })
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

//Post is a function to register a route against HTTP POST request method
//
//  app := essen.App()
//
//  app.Post("/", func(res app.Response, req app.Request){
//  	//do something
//  })
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

//Put is a function to register a route against HTTP PUT request method
//
//  app := essen.App()
//
//  app.Put("/", func(res app.Response, req app.Request){
//  	//do something
//  })
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

//Use is a function to register a route against any HTTP request method.
//
//  app := essen.App()
//
//  app.Use("/", func(res app.Response, req app.Request){
//  	//do something
//  })
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

//Static is function to serve static files on reciept HTTP request.
//
//`path` parameter can either be a directory path or a file path.
//
//  app := essen.App()
//  app.Static("/static", "./static_dir")
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

package essen

import (
	"net/http"

	"github.com/zemirco/uid"
)

type router struct {
	get  handlerStorage
	post handlerStorage
	put  handlerStorage
}

//Get function is used to add a reqeust handler in a router group to repond against GET HTTP request method.
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

//Post function is used to add a request handler in a router group to respond against POST HTTP request method.
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

//Put function is used to add a request handler in a router group to respond against PUT HTTP request method.
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

//UseRouter function is called to register router against a parent route.
//
//Look at Router() method for an example.
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

//Done function is used to clear all router related data once it is passed in UseRouter function.
func (r *router) Done() {
	*r = router{}
}

//Router function is used to get router instance which can then be used to group different request handlers.
//
//  app := essen.App()
//  router := essen.Router()
//
//  router.Get("/index", func(res essen.Response, req essen.Request){
//		//do something
//  })
//
//	router.Post("/form", func(res essen.Response, req, essen.Request){
//		//do something
//	})
//
//	app.UseRouter("/user", router)
//  router.Done() //Call Done after using a router, Used to achieve memory efficiency.
func (e Essen) Router() router {
	return router{get: make(handlerStorage), post: make(handlerStorage), put: make(handlerStorage)}
}

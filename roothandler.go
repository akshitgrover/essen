package essen

import (
	"essen/jobqueue"
	"net/http"
)

func rootHandler(res http.ResponseWriter, req *http.Request) {

	//Handler For Static Methods
	staticPath := matchStaticUrl(req.URL.Path)
	if staticPath == "" {
		http.NotFound(res, req)
		return
	}
	vStatic, ok := paths.static[staticPath]
	if ok {
		jobqueue.QueuePush(func() {
			http.StripPrefix(staticPath+"/", vStatic).ServeHTTP(res, req)
			jobqueue.QueueNext()
		})
		return
	}

	//Handler For Use Methods
	vUse, ok := paths.use[req.URL.Path]
	if ok {
		jobqueue.QueuePush(func() {
			vUse(res, req)
			jobqueue.QueueNext()
		})
		return
	}
	switch req.Method {

	//Handler Head Requests
	case "HEAD":
		v, ok := paths.head[req.URL.Path]
		if !ok {
			jobqueue.QueuePush(func() {
				http.NotFound(res, req)
				jobqueue.QueueNext()
			})
			return
		}
		jobqueue.QueuePush(func() {
			v(res, req)
			jobqueue.QueueNext()
		})
		break

	//Handle Get Requests
	case "GET":
		v, ok := paths.get[req.URL.Path]
		if !ok {
			jobqueue.QueuePush(func() {
				http.NotFound(res, req)
				jobqueue.QueueNext()
			})
			return
		}
		jobqueue.QueuePush(func() {
			v(res, req)
			jobqueue.QueueNext()
		})
		break

	//Handle Post Requests
	case "POST":
		v, ok := paths.post[req.URL.Path]
		if !ok {
			jobqueue.QueuePush(func() {
				http.NotFound(res, req)
				jobqueue.QueueNext()
			})
			return
		}
		jobqueue.QueuePush(func() {
			v(res, req)
			jobqueue.QueueNext()
		})
		break

	//Handle Any Other Request Method
	default:
		jobqueue.QueuePush(func() {
			http.NotFound(res, req)
			jobqueue.QueueNext()
		})
		break
	}
}

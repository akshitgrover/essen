package essen

import (
	"net/http"

	"github.com/akshitgrover/essen/jobqueue"
)

func rootHandler(res http.ResponseWriter, req *http.Request) {

	done := make(chan bool)

	// Handler For Static Methods
	staticPath := matchStaticUrl(req.URL.Path)
	vStatic, ok := paths.static[staticPath]
	if ok {
		jobqueue.QueuePush(func() {
			http.StripPrefix(staticPath+"/", vStatic).ServeHTTP(res, req)
			done <- true
			jobqueue.QueueNext()
		})
		<-done
		return
	}

	//Handler For Use Methods
	vUse, useOk := paths.use[req.URL.Path]
	if useOk {
		jobqueue.QueuePush(func() {
			vUse(res, req)
			done <- true
			jobqueue.QueueNext()
		})
		<-done
		return
	}

	switch req.Method {

	//Handler Head Requests
	case "HEAD":
		v, ok := paths.head[req.URL.Path]
		if !ok {
			jobqueue.QueuePush(func() {
				http.NotFound(res, req)
				done <- true
				jobqueue.QueueNext()
			})
			break
		}
		jobqueue.QueuePush(func() {
			v(res, req)
			done <- true
			jobqueue.QueueNext()
		})
		break

	//Handle Get Requests
	case "GET":
		v, ok := paths.get[req.URL.Path]
		if !ok {
			jobqueue.QueuePush(func() {
				http.NotFound(res, req)
				done <- true
				jobqueue.QueueNext()
			})
			break
		}
		jobqueue.QueuePush(func() {
			v(res, req)
			done <- true
			jobqueue.QueueNext()
		})
		break

	//Handle Post Requests
	case "POST":
		v, ok := paths.post[req.URL.Path]
		if !ok {
			jobqueue.QueuePush(func() {
				http.NotFound(res, req)
				done <- true
				jobqueue.QueueNext()
			})
			break
		}
		jobqueue.QueuePush(func() {
			v(res, req)
			done <- true
			jobqueue.QueueNext()
		})
		break

	//Handle Put Requests
	case "PUT":
		v, ok := paths.put[req.URL.Path]
		if !ok {
			jobqueue.QueuePush(func() {
				http.NotFound(res, req)
				done <- true
				jobqueue.QueueNext()
			})
			break
		}
		jobqueue.QueuePush(func() {
			v(res, req)
			done <- true
			jobqueue.QueueNext()
		})
		break

	//Handle Any Other Request Method
	default:
		jobqueue.QueuePush(func() {
			http.NotFound(res, req)
			done <- true
			jobqueue.QueueNext()
		})
		break
	}
	<-done

}

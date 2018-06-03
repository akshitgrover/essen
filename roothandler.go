package essen

import (
	"net/http"
)

func rootHandler(res http.ResponseWriter, req *http.Request) {
	v, ok := paths.use[req.URL.Path]
	if ok {
		v(res, req)
		return
	}
	switch req.Method {

	//Handler Head Requests
	case "HEAD":
		v, ok := paths.head[req.URL.Path]
		if !ok {
			http.NotFound(res, req)
			return
		}
		v(res, req)
		break

	//Handle Get Requests
	case "GET":
		v, ok = paths.get[req.URL.Path]
		if !ok {
			http.NotFound(res, req)
			return
		}
		v(res, req)
		break

	//Handle Post Requests
	case "POST":
		v, ok = paths.post[req.URL.Path]
		if !ok {
			http.NotFound(res, req)
			return
		}
		v(res, req)
		break

	//Handle Any Other Request Method
	default:
		http.NotFound(res, req)
		break
	}
}

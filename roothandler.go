package essen

import (
	"net/http"
	"strings"
)

func rootHandler(res http.ResponseWriter, req *http.Request) {

	//Handler For Static Methods
	staticPath := "/" + strings.Split(req.URL.Path, "/")[1]
	vStatic, ok := paths.static[staticPath]
	if ok {
		http.StripPrefix(staticPath+"/", vStatic).ServeHTTP(res, req)
		return
	}

	//Handler For Use Methods
	vUse, ok := paths.use[req.URL.Path]
	if ok {
		vUse(res, req)
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
		v, ok := paths.get[req.URL.Path]
		if !ok {
			http.NotFound(res, req)
			return
		}
		v(res, req)
		break

	//Handle Post Requests
	case "POST":
		v, ok := paths.post[req.URL.Path]
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

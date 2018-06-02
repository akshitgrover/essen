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
	case "GET":
		v, ok = paths.get[req.URL.Path]
		if !ok {
			http.NotFound(res, req)
			return
		}
		v(res, req)
	case "POST":
		v, ok = paths.post[req.URL.Path]
		if !ok {
			http.NotFound(res, req)
			return
		}
		v(res, req)
	default:
		http.NotFound(res, req)
	}
}

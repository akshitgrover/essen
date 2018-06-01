package essen

import (
	"log"
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

func (e Essen) Listen(port int) {
	sport := strconv.Itoa(port)
	fport := ":" + sport
	println("Lifting On Port: " + sport)
	log.Fatal(http.ListenAndServe(fport, nil))
}

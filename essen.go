package essen

import (
	"log"
	"net/http"
	"strconv"
)

var essenGlobal Essen

//Essen struct
type Essen struct {
	Locals map[string]interface{}
}

//Essen Response type
type Response struct {
	Res       http.ResponseWriter
	ReqMethod string
}

//Essen Request Type
type Request struct {
	Req  *http.Request
	App  map[string]interface{}
	Body Param
	Uid  string
}

//Return Essen struct
func App() Essen {
	e := Essen{Locals: make(map[string]interface{})}
	essenGlobal = e
	return e
}

//Listener Port Handler
func (e Essen) Listen(port int) {
	sport := strconv.Itoa(port)
	fport := ":" + sport
	println("Lifting On Port: " + sport)

	//Register Request/Response Root Handler
	log.Fatal(http.ListenAndServe(fport, http.HandlerFunc(rootHandler)))
}

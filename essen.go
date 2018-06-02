package essen

import (
	"log"
	"net/http"
	"strconv"
)

//Essen struct
type Essen struct{}

//Essen Response type
type Response struct {
	Res http.ResponseWriter
}

//Essen Request Type
type Request struct {
	Req  *http.Request
	Body Param
}

//Return Essen struct
func App() Essen {
	e := Essen{}
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

/*
Package essen is a micro web framework for golang, It resembles express.js (node.js web framework) functions with nearly one to one mapping.
*/
package essen

import (
	"log"
	"net/http"
	"strconv"
)

var essenGlobal Essen

//Essen is the root struct, It's instance provide the main entrypoint to access all other functions.
//
//Expressjs similarity
//
//let app = express();
type Essen struct {

	//To share values between every request handler, Accessible in request handler.
	//Checkout Request type for more.
	//Lifetime: Until app is running
	Locals map[string]interface{}
}

//Response type is a wrapper around native http.ResponseWriter.
//
//This type is used to access response functions built in essen.
type Response struct {

	//Copy of native http.ResponseWriter instance.
	Res       http.ResponseWriter
	ReqMethod string
}

//Request type is a wrapper around native http.Request.
//
//This type is used to access request functions built in essen.
//Essen Request Type
type Request struct {

	//Reference to native http.Request instance.
	Req *http.Request

	//Copy of Essen.Locals.
	//Properties set in Essen.Locals are accessed in request handlers through req.App.
	App map[string]interface{}

	//Map to pass values within middlewares.
	//Lifetime: Until request is responded.
	Locals map[string]interface{}

	//Field to access request body parameters.
	//
	//  val, err := req.Body.Params("key")
	//
	//  if err.IsNil(){
	//		fmt.Pritln(val)
	//  }
	Body param

	//Unique id per request. (Used for request data cleanup)
	Uid string
}

//App is an entrypoint to essen functions.
//
//  app := essen.App()
func App() Essen {
	e := Essen{Locals: make(map[string]interface{})}
	essenGlobal = e
	return e
}

//Listen function lets you lift HTTP server.
//
//  app := essen.App()
//  app.Listen(8080) //Will listen for requests on port 8080.
func (e Essen) Listen(port int) {
	sport := strconv.Itoa(port)
	fport := ":" + sport
	println("Lifting On Port: " + sport)

	//Register Request/Response Root Handler
	log.Fatal(http.ListenAndServe(fport, http.HandlerFunc(rootHandler)))
}

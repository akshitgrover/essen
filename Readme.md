# Essen

[![GoDoc](https://godoc.org/github.com/akshitgrover/essen?status.svg)](https://godoc.org/github.com/akshitgrover/essen)
[![Go Report Card](https://goreportcard.com/badge/github.com/akshitgrover/essen)](https://goreportcard.com/report/github.com/akshitgrover/essen)

Essen is a golang based micro web framework, inspired by express.js framework for node.js.
It is in a  very early stage, with some request and response mapping, With time it will ripe.

## Why the name?

The time I started developing it I had that sensation of eating something, in german **essen** is **To eat** so the name :P :'|)

## 10 seconds to code:

```go
package main

import(
	"github.com/akshitgrover/essen"
)

func main(){
	e := essen.App()

	e.Use("/getpath",func(res essen.Response,req essen.Request){
		res.Send(200, "Hello World")
	}) // Use Method is for any request method

	// Similar Methods For Post And Get Request

	e.Listen(<port>) // Specify The PORT
}
```

## P.S.

This project is in very early stage, will be developing it further.

Will love any suggestions and ideas for the same.

## Copyright & License

[MIT License](./LICENSE)

Copyright (c) 2018 Akshit Grover
# Essen

Essen is a golang based micro web framework, inspired by express.js framework for node.js.
It is in a  very early stage, with some request and response mapping, With time it will ripe.

## Why this name?

The time I started developing it I had that sensation of eating something :stuck_out_tongue: in german **essen** is **To eat** so the name :stuck_out_tongue: :bowtie:

## Usage:

```go
package main

import(
	"github.com/akshitgrover/essen"
)

func main(){
	e := essen.App()

	e.Use("/getpath",func(res essen.Response,req essen.Request){
		res.Res.Write([]byte("Hello World"))
	}) // Use Method is for any request method

	// Similar Methods For Post And Get Request

	e.Listen(<port>) // Specify The PORT
}
```

## P.S.

This project is in very early stage, will be developing it further.

Will :heart: any suggestions and ideas for the same.
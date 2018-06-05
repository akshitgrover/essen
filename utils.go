package essen

import (
	"os"
)

func CreateFileIfNotExist(path string) (*os.File, EssenError) {
	ee := EssenError{nilval: true}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		ee.nilval = false
		ee.errortype = "PathError"
		ee.message = err.Error()
		return f, ee
	}
	return f, ee
}

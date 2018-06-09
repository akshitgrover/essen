package essen

import (
	"os"
)

const (
	Minute = 60
	Hour   = 60 * Minute
	Day    = 24 * Hour
	Week   = 7 * Day
	Month  = 4 * Week
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

func CreateDirIfNotExist(path string) EssenError {
	ee := EssenError{nilval: true}
	err := os.MkdirAll(path, 0777)
	if err != nil {
		ee.nilval = false
		ee.errortype = "PathError"
		ee.message = err.Error()
	}
	return ee
}

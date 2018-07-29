package essen

import (
	"essen/jobqueue"
	"os"
	"regexp"
)

const (
	Minute = 60
	Hour   = 60 * Minute
	Day    = 24 * Hour
	Week   = 7 * Day
	Month  = 4 * Week
)

type TemplateFunc map[string]interface{}

func GetTemplateFunc() TemplateFunc {
	return make(TemplateFunc)
}

func (t TemplateFunc) Push(key string, f interface{}) {
	t[key] = f
}

func CreateFileIfNotExist(path string) (*os.File, EssenError) {
	ee := EssenError{nilval: true}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0666)
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

func SetConcurrencyLimit(n int) {
	jobqueue.SetConcurrecny(n)
}

func matchStaticUrl(url string) string {
	l := 0
	mString := ""
	for k := range paths.static {
		matched, _ := regexp.MatchString("^"+k, url)
		if matched && len(k) > l {
			l = len(k)
			mString = k
		}
	}
	return mString
}

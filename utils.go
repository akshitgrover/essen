package essen

import (
	"os"
	"regexp"

	"github.com/akshitgrover/essen/jobqueue"
)

const (

	//Minute constant to get numerical value of time unit
	Minute = 60

	//Hour constant to get numerical value of time unit
	Hour = 60 * Minute

	//Day constant to get numerical value of time unit
	Day = 24 * Hour

	//Week constant to get numerical value of time unit
	Week = 7 * Day

	//Month constant to get numerical value of time unit
	Month = 4 * Week
)

//TemplateFunc type instance is used to pass functions in go templates
type TemplateFunc map[string]interface{}

//GetTemplateFunc returns an instance of TemplateFunc type
//
//  tf := essen.GetTemplateFunc()
func GetTemplateFunc() TemplateFunc {
	return make(TemplateFunc)
}

//Push method is used to add function to TemplateFunc instance
func (t TemplateFunc) Push(key string, f interface{}) {
	t[key] = f
}

//CreateFileIfNotExist is used to create a file if does not exists.
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

//CreateDirIfNotExist is used to create a directory if does not exists.
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

//SetConcurrencyLimit is used to set limit on concurrently running request handlers
func SetConcurrencyLimit(n int) {
	jobqueue.SetConcurrency(n)
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

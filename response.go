package essen

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

//Json function to send respose in JSON format
//
//  res.Json(200, map[string]string{"msg":"Hello Essen"})
func (r Response) Json(status int, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	r.Res.Header().Set("Content-Type", "application/json")
	r.Res.WriteHeader(status)
	r.Res.Write(b)
	return nil
}

//Send function to send response in HTML format
//
//  res.Send(200, "Hello World")
func (r Response) Send(status int, v string) {
	r.Res.Header().Set("Content-Type", "text/html")
	r.Res.WriteHeader(status)
	r.Res.Write([]byte(v))
}

//SendFile function to send a file in http response
//
//  res.Set("Content-Type", "test/plain")
//  res.SendFile(200, "./notes.txt")
func (r Response) SendFile(status int, path string) (int64, EssenError) {
	ee := EssenError{nilval: true}
	f, err := os.Open(path)
	if err != nil {
		ee.nilval = false
		ee.errortype = "PathError"
		ee.message = err.Error()
		return 0, ee
	}
	r.Res.WriteHeader(status)
	n, err := io.Copy(r.Res, f)
	if err != nil {
		ee.nilval = false
		ee.errortype = "FileCopy"
		ee.message = err.Error()
		return 0, ee
	}
	return n, ee
}

//Cookie function to set cookies
//
//  res.Cookie("key", "val", 60, false, false)
//`Age` parameter is given in seconds.
//
//Check https://golang.org/pkg/net/http/#Cookie for more details
func (r Response) Cookie(key string, val string, age int, secure bool, httpOnly bool) {
	c := &http.Cookie{Name: key, Value: val, MaxAge: age, Secure: secure, HttpOnly: httpOnly}
	http.SetCookie(r.Res, c)
}

//Set function is used to set headers
//
//'k' parameter key
//
//'v' parameter value
func (r Response) Set(k string, v string) {
	r.Res.Header().Set(k, v)
}

//SendStatus function is used to write headers.
//
//Used mostly is case to reply with no body (HEAD request - response)
func (r Response) SendStatus(status int) {
	r.Res.WriteHeader(status)
	r.Res.Write([]byte(""))
}

//Render is used to send a go template as response
//
//  res.Render(200, "./templates/index.gohtml", map[string]string{"hello":"world"}, nil)
func (r Response) Render(status int, filename string, data interface{}, f TemplateFunc) {

	//Create Custom Error
	ee := EssenError{nilval: true}
	base := filepath.Base(filename)
	abs, err := filepath.Abs(filename)

	//Check Absolute Path Conversion Error
	if err != nil {
		ee.nilval = false
		ee.errortype = "PathError"
		ee.message = "Template absolute path conversion error"
		log.Panic(ee)
	}

	//Execute Template
	err = template.Must(template.New(base).Funcs(template.FuncMap(f)).ParseFiles(abs)).ExecuteTemplate(r.Res, base, data)

	//Check Template Execution Error
	if err != nil {
		ee.nilval = false
		ee.errortype = "TemplateError"
		ee.message = "Error while executing template"
		log.Panic(ee)
	}
}

//Redirect function is used to redirect to another URL/Route
//
//  res.Redirect(302, '/')
//  res.Redirect(302, "google.com")
func (r Response) Redirect(status int, url string) {
	r.Set("Location", url)
	r.SendStatus(status)
}

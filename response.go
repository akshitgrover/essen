package essen

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

//Send JSON Response
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

//Send HTML Response
func (r Response) Send(status int, v string) {
	r.Res.Header().Set("Content-Type", "text/html")
	r.Res.WriteHeader(status)
	r.Res.Write([]byte(v))
}

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

//Set Cookie
func (r Response) Cookie(key string, val string, age int, secure bool, httpOnly bool) {
	c := &http.Cookie{Name: key, Value: val, MaxAge: age, Secure: secure, HttpOnly: httpOnly}
	http.SetCookie(r.Res, c)
}

//Set Headers
func (r Response) Set(k string, v string) {
	r.Res.Header().Set(k, v)
}

//Send Response With Empty Body
func (r Response) SendStatus(status int) {
	r.Res.WriteHeader(status)
	r.Res.Write([]byte(""))
}

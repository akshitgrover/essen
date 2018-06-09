package essen

import (
	"encoding/json"
	"net/http"
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

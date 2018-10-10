package essen

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type getBody struct {
	body *url.URL
}

type postBody struct {
	body *http.Request
}

type multiPartBody struct {
	body *http.Request
	uid  string
}

type param interface {
	Params(name string) (string, EssenError)
}

func (b getBody) Params(name string) (string, EssenError) {
	v := b.body.Query().Get(name)
	ee := EssenError{nilval: true}
	if v == "" {
		ee.nilval = false
		ee.errortype = "InvalidParam"
		ee.message = `No parameter with key "` + name + `"`
	}
	return v, ee
}

func (b postBody) Params(name string) (string, EssenError) {
	ee := EssenError{nilval: true}
	v := b.body.PostFormValue(name)
	if v == "" {
		ee.nilval = false
		ee.errortype = "InvalidParam"
		ee.message = `No parameter with key "` + name + `"`
		return "", ee
	}
	return v, ee
}

func (b multiPartBody) Params(name string) (string, EssenError) {
	ee := EssenError{nilval: true}
	uid := b.uid
	v, ok := uploadedPaths[uid][name]
	if ok {
		return v, ee
	}
	file, fileHeader, err := b.body.FormFile(name)
	if err != nil && err.Error() != "http: no such file" {
		ee.nilval = false
		ee.errortype = "FormParseError"
		ee.message = err.Error()
		return "", ee
	} else if err == nil {
		UploadDir := MultiPartConfig["UploadDir"]
		path := UploadDir + "/" + fileHeader.Filename
		f, ee := CreateFileIfNotExist(path)
		if !ee.IsNil() {
			return "", ee
		}
		n, err := io.Copy(f, file)
		log.Println(n, err)
		m := uploadedPaths[uid].push(name, path)
		if m != nil {
			uploadedPaths[uid] = m
		}
		return path, ee
	}
	v = b.body.FormValue(name)
	if v == "" {
		ee.nilval = false
		ee.errortype = "InvalidParam"
		ee.message = `No parameter with key "` + name + `"`
		return "", ee
	}
	return v, ee
}

//Path functions is used to get URL path
func (r Request) Path() string {
	return r.Req.URL.Path
}

//Host is used to get host address from the url
func (r Request) Host() string {
	return r.Req.URL.Host
}

//Method is used to check HTTP request method.
func (r Request) Method() string {
	return r.Req.Method
}

//HasHeader is used to check if a header is sent for the request.
func (r Request) HasHeader(key string) bool {
	v := r.Req.Header.Get(key)
	if v == "" {
		return false
	}
	return true
}

//HasCookie is used to check if a cookie exists in a request
func (r Request) HasCookie(key string) bool {
	_, err := r.Req.Cookie(key)
	if err != nil {
		return false
	}
	return true
}

//CookieVal is used to get value of a cookie
func (r Request) CookieVal(key string) (string, EssenError) {
	ee := EssenError{nilval: true}
	if !r.HasCookie(key) {
		ee.nilval = false
		ee.errortype = "CookieError"
		ee.message = "Cookie does not exist"
		return "", ee
	}
	cookie, _ := r.Req.Cookie(key)
	return cookie.Value, ee
}

//Header function is used to get value of request headers
func (r Request) Header(key string) (string, EssenError) {
	if r.HasHeader(key) {
		hval := r.Req.Header.Get(key)
		ok := strings.HasPrefix(hval, "multipart")
		if ok {
			hval = strings.Split(hval, ";")[0]
		}
		return hval, EssenError{nilval: true}
	}
	return "", EssenError{message: "No Header Found", errortype: "NoHeader", nilval: false}
}

func (r *Request) requestBody() {
	contentType, ee := r.Header("Content-Type")
	r.App = essenGlobal.Locals
	r.Locals = make(map[string]interface{})
	if !ee.IsNil() {
		contentType = "application/x-www-form-urlencoded"
	}
	if contentType == "multipart/form-data" {
		if !mConfigIsSet() {
			setDefaultConfig()
		}
		r.Body = multiPartBody{body: r.Req, uid: r.Uid}
		return
	}
	if r.Method() == "GET" || r.Method() == "HEAD" {
		r.Body = getBody{body: r.Req.URL}
		return
	}
	if r.Method() == "POST" || r.Method() == "PUT" {
		err := r.Req.ParseForm()
		if err != nil {
			ee := EssenError{nilval: false, errortype: "FormParseError", message: err.Error()}
			log.Panic(ee.Error())
		}
		r.Body = postBody{body: r.Req}
		return
	}
}

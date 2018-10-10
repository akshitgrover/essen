package essen

import (
	"log"
	"path/filepath"
)

//Paths for uploaded files
type uploadParam map[string]string

func (u uploadParam) push(key string, value string) uploadParam {
	if u != nil {
		u[key] = value
	} else {
		u = make(uploadParam)
		u[key] = value
		return u
	}
	return nil
}

//Close method is used to clear request related data (Mostly Multipart Data)
//
//  req.Close()
func (r Request) Close() {
	_, ok := uploadedPaths[r.Uid]
	if ok {
		delete(uploadedPaths, r.Uid)
	}
}

var uploadedPaths = make(map[string]uploadParam)

//isSet
var isSet = false

//MultiPartConfig map is used to store configuration data for multipart requests.
var MultiPartConfig = map[string]string{"UploadDir": Defaults.UploadDir} //optional

//SetMultiPartConfig function is used to set multipart requests configuration
//
//It can be set directly by changing MultiPartConfig map but calling this function is recommended way of doing so.
//
//  app := essen.App()
//  app.SetMultiPartConfig(map[string]string{"UploadDir": "./uploadsFolder"}) //optional
func (e Essen) SetMultiPartConfig(configMap map[string]string) bool {
	if configMap["UploadDir"] == "" {
		configMap["UploadDir"] = Defaults.UploadDir
	}
	absPath, err := filepath.Abs(configMap["UploadDir"])
	if err != nil {
		log.Fatal(err.Error())
	}
	configMap["UploadDir"] = absPath
	ee := CreateDirIfNotExist(configMap["UploadDir"])
	if !ee.IsNil() {
		log.Fatal(ee.Error())
	}
	MultiPartConfig = configMap
	isSet = true
	return true
}

//Config is set
func mConfigIsSet() bool {
	return isSet
}

//Create Upload Directory
func setDefaultConfig() {
	ee := CreateDirIfNotExist(Defaults.UploadDir)
	if !ee.IsNil() {
		log.Panic(ee.Error())
	}
	isSet = true
}

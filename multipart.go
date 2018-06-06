package essen

import (
	"log"
)

//isSet
var isSet = false

//Multipart Config
var MultiPartConfig = map[string]string{"UploadDir": Defaults.UploadDir}

func (e Essen) SetMultiPartConfig(configMap map[string]string) bool {
	if configMap["UploadDir"] == "" {
		configMap["UploadDir"] = Defaults.UploadDir
	}
	f, ee := CreateFileIfNotExist(configMap["UploadDir"])
	defer f.Close()
	if !ee.IsNil() {
		log.Panic(ee.Error())
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

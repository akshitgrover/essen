package essen

type Default struct {
	UploadDir string
}

var Defaults = Default{UploadDir: "./uploads"}

package jobqueue

type Default struct {
	concurrencyLimit int
}

var Defaults = Default{concurrencyLimit: 25}

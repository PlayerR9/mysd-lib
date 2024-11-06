package common

// commonT is for private use only.
type commonT struct{}

// SD is the namespace for SD-like functions.
var SD commonT

func init() {
	SD = commonT{}
}

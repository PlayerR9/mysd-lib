package pointers

// pointersT is for private use only.
type pointersT struct{}

// SD is the namespace for SD-like functions.
var SD pointersT

func init() {
	SD = pointersT{}
}

package slices

// slicesT is for private use only.
type slicesT struct{}

// SD is the namespace for SD-like functions.
var SD slicesT

func init() {
	SD = slicesT{}
}

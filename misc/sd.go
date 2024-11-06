package misc

// miscT is for private use only.
type miscT struct{}

// SD is the namespace for SD-like functions.
var SD miscT

func init() {
	SD = miscT{}
}

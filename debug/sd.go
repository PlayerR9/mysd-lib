package debug

// debugT is for private use only.
type debugT struct{}

// SD is the namespace for SD-like functions.
var SD debugT

func init() {
	SD = debugT{}
}

package bytes

// bytesT is for private use only.
type bytesT struct{}

// SD is the namespace for SD-like functions.
var SD bytesT

func init() {
	SD = bytesT{}
}

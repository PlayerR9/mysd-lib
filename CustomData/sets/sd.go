package sets

// setsT is for private use only.
type setsT struct{}

// SD is the namespace for SD-like functions.
var SD setsT

func init() {
	SD = setsT{}
}

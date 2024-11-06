package listlike

// listlikeT is for private use only.
type listlikeT struct{}

// SD is the namespace for SD-like functions.
var SD listlikeT

func init() {
	SD = listlikeT{}
}

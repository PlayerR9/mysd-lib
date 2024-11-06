package colors

// colorsT is for private use only.
type colorsT struct{}

// SD is the namespace for SD-like functions.
var SD colorsT

func init() {
	SD = colorsT{}
}

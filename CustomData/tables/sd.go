package tables

// tablesT is for private use only.
type tablesT struct{}

// SD is the namespace for SD-like functions.
var SD tablesT

func init() {
	SD = tablesT{}
}

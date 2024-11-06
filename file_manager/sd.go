package file_manager

// file_managerT is for private use only.
type file_managerT struct{}

// SD is the namespace for SD-like functions.
var SD file_managerT

func init() {
	SD = file_managerT{}
}

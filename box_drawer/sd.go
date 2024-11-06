package box_drawer

// box_drawerT is for private use only.
type box_drawerT struct{}

// SD is the namespace for SD-like functions.
var SD box_drawerT

func init() {
	SD = box_drawerT{}
}

package tree

// treeT is for private use only.
type treeT struct{}

// SD is the namespace for SD-like functions.
var SD treeT

func init() {
	SD = treeT{}
}

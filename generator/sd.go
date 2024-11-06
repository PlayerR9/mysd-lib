package generator

// generatorT is for private use only.
type generatorT struct{}

// SD is the namespace for SD-like functions.
var SD generatorT

func init() {
	SD = generatorT{}
}

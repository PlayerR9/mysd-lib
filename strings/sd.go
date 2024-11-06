package strings

// stringsT is for private use only.
type stringsT struct{}

// SD is the namespace for SD-like functions.
var SD stringsT

func init() {
	SD = stringsT{}
}

// IndicesOf is the SD version of IndicesOf that, given a separator string, returns a function that
// returns the indices of the separator in the given slice of strings.
//
// Parameters:
//   - sep: The separator.
//
// Returns:
//   - func([]string) []uint: The function that returns the indices. Never returns nil.
func (stringsT) IndicesOf(sep string) func(slice []string) []uint {
	fn := func(slice []string) []uint {
		sliceSize := uint(len(slice))
		if sliceSize == 0 {
			return nil
		}

		var count uint

		for i := uint(0); i < sliceSize; i++ {
			if slice[i] != sep {
				continue
			}

			count++
		}

		if count == 0 {
			return nil
		}

		indices := make([]uint, 0, count)
		var lenIndices uint

		for i := uint(0); i < sliceSize && lenIndices < count; i++ {
			if slice[i] != sep {
				continue
			}

			indices = append(indices, i)
			lenIndices++
		}

		return indices
	}

	return fn
}

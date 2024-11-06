package strings

import (
	"strconv"
)

// QuoteInPlace quotes each string in the given slice of strings in-place.
//
// Parameters:
//   - elems: The slice of strings to quote.
//
// Example:
//
//	QuoteInPlace([]string{"a", "b", "c"}) // => []string{"\"a\"", "\"b\"", "\"c\""}
func Quote(elems []string) {
	if len(elems) == 0 {
		return
	}

	for i := 0; i < len(elems); i++ {
		elems[i] = strconv.Quote(elems[i])
	}
}

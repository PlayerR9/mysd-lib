package slices

import (
	"fmt"
	"slices"

	"github.com/PlayerR9/mysd-lib/common"
)

// NewErrNotAsExpected is a convenience function that creates a new ErrNotAsExpected error with
// the specified kind, got value, and expected values.
//
// See common.NewErrNotAsExpected for more information.
func NewErrNotAsExpected[T any](quote bool, kind string, got any, expecteds ...T) error {
	unique := make([]string, 0, len(expecteds))

	for _, expected := range expecteds {
		str := fmt.Sprint(expected)

		pos, ok := slices.BinarySearch(unique, str)
		if !ok {
			unique = slices.Insert(unique, pos, str)
		}
	}

	unique = unique[:len(unique):len(unique)]

	var got_str string

	if got != nil {
		got_str = fmt.Sprint(got)
	}

	return common.NewErrNotAsExpected(quote, kind, got_str, unique...)
}

package file_manager

import (
	"cmp"
	"slices"

	"github.com/PlayerR9/mysd-lib/common"
)

// MayInsert attempts to insert an element into a sorted slice if it is not already present.
//
// Parameters:
//   - slice: A pointer to a slice of ordered elements.
//   - elem: The element to insert.
//
// Returns:
//   - error: Returns ErrBadParam if slice is nil.
//
// If the element is not found in the slice, it is inserted in the correct position to maintain order.
func MayInsert[T cmp.Ordered](slice *[]T, elem T) error {
	if slice == nil {
		return common.NewErrNilParam("slice")
	}

	pos, ok := slices.BinarySearch(*slice, elem)
	if ok {
		return nil
	}

	*slice = slices.Insert(*slice, pos, elem)

	return nil
}

package slices

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
//   - bool: Returns true if the element was inserted into the slice, false otherwise.
//   - error: Returns ErrBadParam if slice is nil.
//
// If the element is not found in the slice, it is inserted in the correct position to maintain order.
func MayInsert[T cmp.Ordered](slice *[]T, elem T) (bool, error) {
	if slice == nil {
		return false, common.NewErrNilParam("slice")
	}

	pos, ok := slices.BinarySearch(*slice, elem)
	if ok {
		return false, nil
	}

	*slice = slices.Insert(*slice, pos, elem)

	return true, nil
}

// Uniquefy removes duplicate elements from a slice, in-place, while also sorting the slice in
// ascending order.
//
// Parameters:
//   - slice: A pointer to the slice where duplicate elements will be removed from.
//
// Returns:
//   - int: The number of elements removed from the slice.
//
// The slice is also resized while clearing the removed elements.
func Uniquefy[T cmp.Ordered](slice *[]T) int {
	if slice == nil {
		return 0
	}

	unique := make([]T, 0, len(*slice))

	for _, elem := range *slice {
		pos, ok := slices.BinarySearch(unique, elem)
		if ok {
			continue
		}

		unique = slices.Insert(unique, pos, elem)
	}

	clear((*slice)[len(unique):])

	n := cap(unique) - len(unique)

	unique = unique[:len(unique):len(unique)]

	*slice = unique

	return n
}

// Merge inserts elements from the 'from' slice into the 'dest' slice, maintaining order and ensuring no duplicates.
//
// Parameters:
//   - dest: A pointer to the destination slice where elements will be inserted. This slice must be sorted
//     and free of duplicates. If not, Uniquefy must be called on it first.
//   - from: The slice of elements to merge into the destination.
//
// Returns:
//   - int: The number of elements that were not inserted or could not be merged.
//   - error: Returns ErrBadParam if dest is nil.
//
// If 'from' is empty, the function does nothing. Each element from 'from' is inserted into 'dest' in the correct position,
// ensuring that 'dest' remains sorted and free of duplicates.
func Merge[T cmp.Ordered](dest *[]T, from []T) (int, error) {
	if len(from) == 0 {
		return 0, nil
	} else if dest == nil {
		return len(from), common.NewErrNilParam("dest")
	}

	var n int

	for _, elem := range from {
		pos, ok := slices.BinarySearch(*dest, elem)
		if ok {
			n++
		} else {
			*dest = slices.Insert(*dest, pos, elem)
		}
	}

	return n, nil
}

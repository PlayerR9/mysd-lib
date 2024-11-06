package pointers

// Pointer is an interface for pointer-like types.
type Pointer interface {
	// IsNil checks whether the pointer is nil.
	//
	// Returns:
	//   - bool: True if the pointer is nil, false otherwise.
	IsNil() bool
}

// RejectNils removes all nil elements from the given slice of pointer-like
// types that implement the Pointer interface.
//
// Parameters:
//   - slice: The slice of pointer-like types to remove nils from.
//
// Returns:
//   - []T: The slice of pointer-like types without nils. Nil if all the
//     elements are nil or no elements were specified.
func RejectNils[T Pointer](slice *[]T) {
	if slice == nil || len(*slice) == 0 {
		return
	}

	var top int

	for _, elem := range *slice {
		ok := elem.IsNil()
		if !ok {
			(*slice)[top] = elem
			top++
		}
	}

	if top == 0 {
		clear(*slice)
		*slice = nil

		return
	}

	clear((*slice)[top:])
	*slice = (*slice)[:top:top]
}

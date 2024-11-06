package pointers

// Set assigns the specified value to the variable pointed to by the pointer. Does
// nothing if the pointer is nil.
//
// Parameters:
//   - p: A pointer to the variable to set.
//   - v: The value to assign to the variable.
func Set[T any](p *T, v T) {
	if p == nil {
		return
	}

	*p = v
}

// Get returns the value of the pointer if it is not nil, otherwise it returns the
// zero value of the type of the pointer.
//
// Parameters:
//   - p: A pointer to the variable to get.
//
// Returns:
//   - T: The value of the pointer.
func Get[T any](p *T) T {
	if p == nil {
		return *new(T)
	} else {
		return *p
	}
}

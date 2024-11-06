package listlike

// Lister is an interface that can be used by list-like data structures.
type Lister interface {
	// Size returns the number of elements in the list-like data structure.
	//
	// Returns:
	//   - int: The number of elements in the list-like data structure. Never negative.
	Size() int

	// IsEmpty checks whether the list-like data structure is empty.
	//
	// Returns:
	//   - bool: True if the list-like data structure is empty, false otherwise.
	IsEmpty() bool

	// Reset resets the list-like data structure for reuse.
	Reset()
}

// baseList is the base implementation of Lister.
type baseList[T any] struct {
	// slice is the underlying list-like data structure.
	slice []T
}

// Size implements the Lister interface.
func (l baseList[T]) Size() int {
	return len(l.slice)
}

// IsEmpty implements the Lister interface.
func (l baseList[T]) IsEmpty() bool {
	return len(l.slice) == 0
}

// Reset resets the list-like data structure for reuse.
func (l *baseList[T]) Reset() {
	if l == nil {
		return
	}

	if len(l.slice) > 0 {
		clear(l.slice)
		l.slice = nil
	}
}

// New creates a new generic list-like data structure.
//
// Parameters:
//   - elems: The elements to add to the list-like data structure.
//
// Returns:
//   - Lister: The new list-like data structure. Never returns nil.
func New[T any](elems ...T) Lister {
	return &baseList[T]{
		slice: elems,
	}
}

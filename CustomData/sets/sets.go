package sets

import (
	"iter"
	"slices"

	"github.com/PlayerR9/mysd-lib/common"
)

// Set is an interface used by any set-like data structure.
type Set[T any] interface {
	// Size returns the number of elements in the set.
	//
	// Returns:
	//   - int: The number of elements in the set. Never negative.
	Size() int

	// IsEmpty checks whether the set is empty.
	//
	// Returns:
	//   - bool: True if the set is empty, false otherwise.
	IsEmpty() bool

	// Reset resets the set for reuse.
	Reset()

	// Add inserts an element into the set if it is not already present.
	//
	// Parameters:
	//   - elem: The element to add to the set.
	//
	// Returns:
	//   - error: Returns an error if the receiver is nil.
	Add(elem T) error

	// AddMany adds multiple elements to the set. It must be equal to calling Add
	// multiple times but can be more efficient than adding each element individually.
	//
	// Parameters:
	//   - elems: The elements to add to the set.
	//
	// Returns:
	//   - error: Returns an error if the receiver is nil.
	AddMany(elems []T) error

	// Contains checks whether the specified element is present in the set.
	//
	// Parameters:
	//   - elem: The element to check for presence in the set.
	//
	// Returns:
	//   - bool: True if the element is present in the set, false otherwise.
	Contains(elem T) bool

	// Elem iterates through the elements in the set. The order of elements
	// is not guaranteed to be the same as the order in which they were added.
	//
	// Returns:
	//   - iter.Seq[T]: The elements in the set. Never returns nil.
	Elem() iter.Seq[T]
}

// Merge merges the elements of another set into the specified set. The elements will
// be added to the set in the order they are returned by the other set's Elem method.
//
// Parameters:
//   - from: The set that the elements will be added to.
//   - other: The set whose elements will be merged into from.
//
// Returns:
//   - error: Returns an error if the sets cannot be merged.
//
// Errors:
//   - common.ErrBadParam: If other has at least one element and from is nil.
func Merge[T any](from, other Set[T]) error {
	if other == nil {
		return nil
	}

	slice := slices.Collect(other.Elem())
	if len(slice) == 0 {
		return nil
	} else if from == nil {
		return common.NewErrNilParam("from")
	}

	_ = from.AddMany(slice)

	return nil
}

// baseSet is the base implementation of the Set interface.
type baseSet[T comparable] struct {
	// elems is the underlying map of elems elements.
	elems map[T]struct{}
}

// Size implements the Set interface.
func (s baseSet[T]) Size() int {
	return len(s.elems)
}

// IsEmpty implements the Set interface.
func (s baseSet[T]) IsEmpty() bool {
	return len(s.elems) == 0
}

// Reset implements the Set interface.
func (s *baseSet[T]) Reset() {
	if s == nil {
		return
	}

	if len(s.elems) > 0 {
		clear(s.elems)
		s.elems = nil
	}
}

// Add implements the Set interface.
func (s *baseSet[T]) Add(elem T) error {
	if s == nil {
		return common.ErrNilReceiver
	}

	if s.elems == nil {
		s.elems = make(map[T]struct{})
	}

	s.elems[elem] = struct{}{}

	return nil
}

// AddMany implements the Set interface.
func (s *baseSet[T]) AddMany(elems []T) error {
	if s == nil {
		return common.ErrNilReceiver
	}

	if s.elems == nil {
		s.elems = make(map[T]struct{})
	}

	for _, k := range elems {
		s.elems[k] = struct{}{}
	}

	return nil
}

// Contains implements the Set interface.
func (s baseSet[T]) Contains(elem T) bool {
	if len(s.elems) == 0 {
		return false
	}

	_, ok := s.elems[elem]
	return ok
}

// Elem implements the Set interface.
func (s baseSet[T]) Elem() iter.Seq[T] {
	return func(yield func(T) bool) {
		for k := range s.elems {
			if !yield(k) {
				return
			}
		}
	}
}

// New creates a new set of comparable elements from the provided elements.
//
// Parameters:
//   - elems: The elements to add to the set.
//
// Returns:
//   - Set[T]: The new set. Never returns nil.
func New[T comparable](elems ...T) Set[T] {
	if len(elems) == 0 {
		return &baseSet[T]{}
	}

	s := &baseSet[T]{
		elems: make(map[T]struct{}),
	}

	for _, k := range elems {
		s.elems[k] = struct{}{}
	}

	return s
}

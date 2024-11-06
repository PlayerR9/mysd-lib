package slices

import "github.com/PlayerR9/mysd-lib/common"

// Builder is a builder for slices. While this is normally not needed, it can have
// some uses when making many slices one after the other.
type Builder[T any] struct {
	// slice is the slice being built.
	slice []T
}

// Append appends an element to the slice being built.
//
// Parameters:
//   - elem: The element to append.
//
// Returns:
//   - error: An error if the receiver is nil.
func (b *Builder[T]) Append(elem T) error {
	if b == nil {
		return common.ErrNilReceiver
	}

	b.slice = append(b.slice, elem)

	return nil
}

// Build builds the slice being built.
//
// Returns:
//   - []T: The slice being built. Nil if no elements were appended.
func (b Builder[T]) Build() []T {
	if len(b.slice) == 0 {
		return nil
	}

	slice := make([]T, len(b.slice), len(b.slice))
	copy(slice, b.slice)

	return slice
}

// Reset resets the builder for reuse.
func (b *Builder[T]) Reset() {
	if b == nil {
		return
	}

	if len(b.slice) > 0 {
		zero := *new(T)

		for i := range b.slice {
			b.slice[i] = zero
		}

		b.slice = nil
	}
}

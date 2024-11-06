package common

// Pair is a pair of two values.
type Pair[A, B any] struct {
	// First is the first value.
	First A

	// Second is the second value.
	Second B
}

// NewPair creates a new pair.
//
// Parameters:
//   - first: The first value.
//   - second: The second value.
//
// Returns:
//   - Pair[A, B]: The new pair.
func NewPair[A, B any](first A, second B) Pair[A, B] {
	return Pair[A, B]{
		First:  first,
		Second: second,
	}
}

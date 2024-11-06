package tree

import "fmt"

// Infoer is the information of the node.
type Infoer interface {
	// Equals checks if the information is equal to another information.
	//
	// Parameters:
	//   - other: The other information to compare with. Assumed to be non-nil.
	//
	// Returns:
	//   - bool: True if the two informations are equal, false otherwise.
	Equals(other Infoer) bool

	fmt.Stringer
}

// baseInfo is the base implementation of the Infoer interface.
type baseInfo[T comparable] struct {
	// v is the information of the node.
	v T
}

// Equals implements the Infoer interface.
func (b baseInfo[T]) Equals(other Infoer) bool {
	v, ok := other.(*baseInfo[T])
	return ok && b.v == v.v
}

// String implements the Stringer interface.
func (b baseInfo[T]) String() string {
	return fmt.Sprint(b.v)
}

// New creates a new node with the given information. This is only
// used for comparable types such as int, string, etc.
//
// Parameters:
//   - v: The information of the node.
//
// Returns:
//   - *Node: The new node. Never returns nil.
func New[T comparable](v T) *Node {
	info := &baseInfo[T]{
		v: v,
	}

	return &Node{
		Info: info,
	}
}

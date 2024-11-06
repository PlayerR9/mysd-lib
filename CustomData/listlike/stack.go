package listlike

import (
	"slices"

	"github.com/PlayerR9/mysd-lib/common"
)

type stackT struct{}

var Stack stackT

func init() {
	Stack = stackT{}
}

type Stacker interface {
	Size() int
	IsEmpty() bool
	Reset()
	Push(elem any)
	PushMany(elems []any)
	Pop() (any, error)
	Peek() (any, error)
}

func (stackT) New(elems ...any) []any {
	if len(elems) == 0 {
		return nil
	}

	slices.Reverse(elems)

	return elems
}

func (stackT) Push(stack *[]any, elem any) {
	if stack == nil {
		panic("no destination was provided")
	}

	*stack = append(*stack, elem)
}

func (stackT) PushMany(stack *[]any, elems ...any) {
	if stack == nil {
		panic("no destination was provided")
	}

	if len(elems) == 0 {
		return
	}

	slices.Reverse(elems)

	*stack = append(*stack, elems...)
}

func (stackT) Pop(stack *[]any) (any, error) {
	if stack == nil || len(*stack) == 0 {
		return 0, ErrEmptyStack
	}

	elem := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]

	return elem, nil
}

func (stackT) Peek(stack []any) (any, error) {
	if len(stack) == 0 {
		return 0, ErrEmptyStack
	}

	return stack[len(stack)-1], nil
}

// ArrayStack is a simple implementation of a stack. An empty stack can either be
// created with the `var stack ArrayStack[T]` syntax or with the `new(ArrayStack[T])`
// constructor.
type ArrayStack[T any] struct {
	// slice is the internal slice.
	slice []T
}

// Size implements the Lister interface.
func (s ArrayStack[T]) Size() int {
	return len(s.slice)
}

// IsEmpty implements the Lister interface.
func (s ArrayStack[T]) IsEmpty() bool {
	return len(s.slice) == 0
}

// Reset implements the Lister interface.
func (s *ArrayStack[T]) Reset() {
	if s == nil {
		return
	}

	if len(s.slice) > 0 {
		clear(s.slice)
		s.slice = nil
	}
}

// NewStack creates a new stack from a slice.
//
// Parameters:
//   - elems: The elements to add to the stack.
//
// Returns:
//   - *Stack[T]: The new stack. Never returns nil.
//
// WARNING: As a side-effect, the original list will be reversed.
func NewStack[T any](elems []T) *ArrayStack[T] {
	if len(elems) == 0 {
		return &ArrayStack[T]{}
	}

	slices.Reverse(elems)

	return &ArrayStack[T]{
		slice: elems,
	}
}

// Push adds an element to the stack.
//
// Parameters:
//   - elem: The element to add.
//
// Returns:
//   - error: An error if the receiver is nil.
func (s *ArrayStack[T]) Push(elem T) error {
	if s == nil {
		return common.ErrNilReceiver
	}

	s.slice = append(s.slice, elem)

	return nil
}

// PushMany adds multiple elements to the stack. If it has at least
// one element but the receiver is nil, an error is returned.
//
// Parameters:
//   - elems: The elements to add.
//
// Returns:
//   - error: An error if the receiver is nil.
//
// WARNING: As a side-effect, the original list will be reversed.
func (s *ArrayStack[T]) PushMany(elems []T) error {
	if len(elems) == 0 {
		return nil
	} else if s == nil {
		return common.ErrNilReceiver
	}

	slices.Reverse(elems)

	s.slice = append(s.slice, elems...)

	return nil
}

// Pop removes an element from the stack.
//
// Returns:
//   - T: The element that was removed.
//   - error: An error if the element could not be removed from the stack.
//
// Errors:
//   - ErrEmptyStack: If the stack is empty.
func (s *ArrayStack[T]) Pop() (T, error) {
	if s == nil || len(s.slice) == 0 {
		return *new(T), ErrEmptyStack
	}

	elem := s.slice[len(s.slice)-1]
	s.slice = s.slice[:len(s.slice)-1]

	return elem, nil
}

// Peek returns the element at the top of the stack.
//
// Returns:
//   - T: The element at the top of the stack.
//   - error: An error if the stack is empty.
//
// Errors:
//   - ErrEmptyStack: If the stack is empty.
func (s ArrayStack[T]) Peek() (T, error) {
	if len(s.slice) == 0 {
		return *new(T), ErrEmptyStack
	}

	return s.slice[len(s.slice)-1], nil
}

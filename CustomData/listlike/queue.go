package listlike

import (
	"github.com/PlayerR9/mysd-lib/common"
)

// Queue is a simple implementation of a queue. An empty queue can either be
// created with the `var queue Queue[T]` syntax or with the `new(Queue[T])`
type Queue[T any] struct {
	// slice is the internal slice.
	slice []T
}

// Size implements the Lister interface.
func (q Queue[T]) Size() int {
	return len(q.slice)
}

// IsEmpty implements the Lister interface.
func (q Queue[T]) IsEmpty() bool {
	return len(q.slice) == 0
}

// Reset implements the Lister interface.
func (q *Queue[T]) Reset() {
	if q == nil {
		return
	}

	if len(q.slice) > 0 {
		clear(q.slice)
		q.slice = nil
	}
}

// NewQueue creates a new queue from a slice.
//
// Parameters:
//   - elems: The elements to add to the queue.
//
// Returns:
//   - *Queue[T]: The new queue. Never returns nil.
func NewQueue[T any](elems []T) *Queue[T] {
	if len(elems) == 0 {
		return &Queue[T]{}
	}

	return &Queue[T]{
		slice: elems,
	}
}

// Enqueue adds an element to the queue.
//
// Parameters:
//   - elem: The element to add.
//
// Returns:
//   - error: An error if the receiver is nil.
func (q *Queue[T]) Enqueue(elem T) error {
	if q == nil {
		return common.ErrNilReceiver
	}

	q.slice = append(q.slice, elem)

	return nil
}

// EnqueueMany adds multiple elements to the queue. If it has at least
// one element but the receiver is nil, an error is returned.
//
// Parameters:
//   - elems: The elements to add.
//
// Returns:
//   - error: An error if the receiver is nil.
func (q *Queue[T]) EnqueueMany(elems []T) error {
	if len(elems) == 0 {
		return nil
	} else if q == nil {
		return common.ErrNilReceiver
	}

	q.slice = append(q.slice, elems...)

	return nil
}

// Dequeue removes the first element from the queue.
//
// Returns:
//   - T: The element that was removed.
//   - error: An error if the element could not be removed from the queue.
//
// Errors:
//   - ErrEmptyQueue: If the queue is empty.
func (q *Queue[T]) Dequeue() (T, error) {
	if q == nil || len(q.slice) == 0 {
		return *new(T), ErrEmptyQueue
	}

	elem := q.slice[0]
	q.slice = q.slice[1:]

	return elem, nil
}

// First returns the element at the start of the queue.
//
// Returns:
//   - T: The element at the start of the queue.
//   - error: An error if the queue is empty.
//
// Errors:
//   - ErrEmptyQueue: If the queue is empty.
func (q Queue[T]) First() (T, error) {
	if len(q.slice) == 0 {
		return *new(T), ErrEmptyQueue
	}

	return q.slice[len(q.slice)-1], nil
}

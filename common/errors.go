package common

import "fmt"

// ErrMust occurs when something must be true.
type ErrMust struct {
	// Inner is the inner error.
	Inner error
}

// Error implements the error interface.
func (e ErrMust) Error() string {
	var msg string

	if e.Inner == nil {
		msg = "something went wrong"
	} else {
		msg = e.Inner.Error()
	}

	return "must: " + msg
}

// NewErrMust creates a new ErrMust error.
//
// Parameters:
//   - inner: The inner error.
//
// Returns:
//   - error: The new error. Never returns nil.
//
// Format:
//
//	"must: <inner>"
//
// Where, <inner>: The inner error. If nil, "something went wrong" is used instead.
func NewErrMust(inner error) error {
	return &ErrMust{
		Inner: inner,
	}
}

// Unwrap returns the inner error.
//
// Returns:
//   - error: The inner error.
func (e ErrMust) Unwrap() error {
	return e.Inner
}

// ErrAt occurs when an error occurs at a specific index.
type ErrAt struct {
	// Idx is the index at which the error occurred.
	Idx int

	// Inner is the inner error.
	Inner error
}

// Error implements the error interface.
func (e ErrAt) Error() string {
	var reason string

	if e.Inner == nil {
		reason = "something went wrong"
	} else {
		reason = e.Inner.Error()
	}

	return fmt.Sprintf("at index %d: %s", e.Idx, reason)
}

// NewErrAt returns a new ErrAt from the given index and inner error.
//
// Parameters:
//   - idx: The index at which the error occurred.
//   - inner: The inner error.
//
// Returns:
//   - error: The new error. Never returns nil.
//
// Format:
//
//	"at index <idx>: <reason>"
//
// Where:
//   - <idx>: The index at which the error occurred.
//   - <reason>: The reason for the error. If nil, "something went wrong" is used instead.
func NewErrAt(idx int, inner error) error {
	return &ErrAt{
		Idx:   idx,
		Inner: inner,
	}
}

// Unwrap implements the errors.Wrapper interface.
//
// Returns:
//   - error: The inner error.
func (e ErrAt) Unwrap() error {
	return e.Inner
}

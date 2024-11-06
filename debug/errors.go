package debug

import (
	"fmt"
)

// ErrPrintFailed is an error that is returned when printing failed.
type ErrPrintFailed struct {
	// Idx is the index of the element that failed to print.
	Idx int

	// Reason is the reason for the failure.
	Reason error
}

// Error implements the error interface.
func (e ErrPrintFailed) Error() string {
	var msg string

	if e.Reason == nil {
		msg = "something went wrong"
	} else {
		msg = e.Reason.Error()
	}

	return fmt.Sprintf("could not print element at index %d: %s", e.Idx, msg)
}

// NewErrPrintFailed creates a new ErrPrintFailed error.
//
// Parameters:
//   - idx: The index of the element that failed to print.
//   - reason: The reason for the failure.
//
// Returns:
//   - error: The new error. Never returns nil.
//
// Format:
//
//	"could not print the <ordinal> element: <reason>"
//
// Where:
// - <ordinal>: The ordinal of the index + 1 according to humanize.Ordinal
// - <reason>: The reason for the failure. If nil, "something went wrong" is used instead.
func NewErrPrintFailed(idx int, reason error) error {
	return &ErrPrintFailed{
		Idx:    idx,
		Reason: reason,
	}
}

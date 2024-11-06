package common

import (
	"errors"
	"io"
	"os"
)

// ErrAssertFail occurs when an assertion fails.
type ErrAssertFail struct {
	// Inner is the error that occurred.
	Inner error
}

// Error implements the error interface.
func (e ErrAssertFail) Error() string {
	var reason string

	if e.Inner == nil {
		reason = "something went wrong"
	} else {
		reason = e.Inner.Error()
	}

	return "ASSERT FAIL: " + reason
}

// NewErrAssertFail returns a new ErrAssertFail from the given error.
//
// Parameters:
//   - inner: The error.
//
// Returns:
//   - error: The new error. Never returns nil.
//
// Format:
//
//	"ASSERT FAIL: <reason>"
//
// Where, <reason> is the error message of the given error. If nil, "something went wrong" is used.
func NewErrAssertFail(inner error) error {
	return &ErrAssertFail{
		Inner: inner,
	}
}

// Unwrap returns the inner error.
//
// Returns:
//   - error: The inner error.
func (e ErrAssertFail) Unwrap() error {
	return e.Inner
}

// Validate panics if the given element is not valid.
//
// Parameters:
//   - elem: The element to validate.
//
// Panics:
//   - ErrAssertFail: If the element is not valid either because it is nil or the Validate method returns an error.
//
// Note: This function is intended to be used for implementing the assert.Validater interface.
// func Validate(elem interface{ Validate() error }) {
// 	if elem == nil {
// 		panic(NewErrAssertFail(ErrNilReceiver))
// 	}

// 	err := elem.Validate()
// 	if err != nil {
// 		panic(NewErrAssertFail(err))
// 	}
// }

// Assert asserts that the given condition is true. If it is not, a panic is triggered
// with an ErrAssertFail error.
//
// Parameters:
//   - cond: The condition to assert.
//   - msg: The error message for the panic.
func Assert(cond bool, msg string) {
	if cond {
		return
	}

	var inner error

	if msg == "" {
		panic(NewErrAssertFail(nil))
	} else {
		inner = errors.New(msg)
	}

	panic(NewErrAssertFail(inner))
}

// TODO panics with a TODO message. The given message is appended to the
// string "TODO: ". If the message is empty, the message "TODO: Handle this
// case" is used instead.
//
// Parameters:
//   - msg: The message to append to the string "TODO: ".
//
// This function is meant to be used only when the code is being built or
// refactored.
func TODO(msg string) {
	if msg == "" {
		panic("TODO: Handle this case")
	} else {
		panic("TODO: " + msg)
	}
}

// Must is a helper function that wraps a call to a function that returns (T, error) and
// panics if the error is not nil.
//
// This function is intended to be used to handle errors in a way that is easy to read and write.
//
// Parameters:
//   - res: The result of the function.
//   - err: The error returned by the function.
//
// Returns:
//   - T: The result of the function.
func Must[T any](res T, err error) T {
	if err != nil {
		panic(NewErrMust(err))
	}

	return res
}

// Warn prints a warning message to the console.
// The message is prefixed with "[WARNING]:" to indicate its nature.
//
// Parameters:
//   - msg: The warning message to be displayed.
//
// Panics if there is an error writing to the standard output.
func Warn(msg string) {
	data := []byte("[WARNING]: " + msg + "\n")

	n, err := os.Stdout.Write(data)
	if err != nil {
		panic(err)
	} else if n != len(data) {
		panic(io.ErrShortWrite)
	}
}

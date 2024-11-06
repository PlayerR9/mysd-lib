package misc

import (
	"errors"
	"fmt"
)

// try executes the provided function and recovers from any panic that occurs,
// converting it to an error which is then assigned to the provided error pointer.
//
// Parameters:
//   - err: A pointer to an error that will be set if a panic occurs.
//   - fn: The function to execute.
//
// Panics:
//   - If a panic occurs, it will be recovered and converted into an error:
//   - If the panic value is a string, it will be wrapped in a new error.
//   - If the panic value is an error, it will be used directly.
//   - Otherwise, the panic value will be formatted as an error message.
func try(err *error, fn func()) {
	defer func() {
		r := recover()
		if r == nil {
			return
		}

		switch r := r.(type) {
		case string:
			*err = errors.New(r)
		case error:
			*err = r
		default:
			*err = fmt.Errorf("panic: %v", r)
		}
	}()

	fn()
}

// Try executes the provided function and recovers from any panic that occurs,
// returning it as an error.
//
// Returns:
//   - error: An error if a panic occurred, or nil if no panic occurred.
func Try(fn func()) error {
	if fn == nil {
		return nil
	}

	var err error

	try(&err, fn)

	return err
}

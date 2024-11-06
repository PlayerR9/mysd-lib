package runes

import (
	"errors"
	"slices"
	"strconv"

	"github.com/PlayerR9/mysd-lib/common"
)

var (
	// ErrBadEncoding occurs when an input (slice of bytes or string) is not valid utf-8.
	// This error can be checked using the == operator.
	//
	// Format:
	// 	"invalid utf-8"
	ErrBadEncoding error
)

func init() {
	ErrBadEncoding = errors.New("invalid utf-8")
}

// NewErrNotAsExpected is a convenience function that creates a new ErrNotAsExpected error with
// the specified kind, got value, and expected values.
//
// See common.NewErrNotAsExpected for more information.
func NewErrNotAsExpected(quote bool, kind string, got *rune, expecteds ...rune) error {
	var got_str string

	if got != nil {
		got_str = string(*got)
	}

	unique := make([]string, 0, len(expecteds))

	for _, expected := range expecteds {
		str := string(expected)

		pos, ok := slices.BinarySearch(unique, str)
		if !ok {
			unique = slices.Insert(unique, pos, str)
		}
	}

	unique = unique[:len(unique):len(unique)]

	return common.NewErrNotAsExpected(quote, kind, got_str, unique...)
}

// ErrAfter is an error that occurs after another error.
type ErrAfter struct {
	// Quote is a flag that indicates that the error should be quoted.
	Quote bool

	// Previous is the previous value.
	Previous *rune

	// Inner is the inner error.
	Inner error
}

// Error implements the error interface.
func (e ErrAfter) Error() string {
	var previous string

	if e.Previous == nil {
		previous = "at the start"
	} else if e.Quote {
		previous = strconv.QuoteRune(*e.Previous)
		previous = "after " + previous
	} else {
		previous = string(*e.Previous)
		previous = "after " + previous
	}

	var reason string

	if e.Inner == nil {
		reason = "something went wrong"
	} else {
		reason = e.Inner.Error()
	}

	return previous + ": " + reason
}

// NewErrAfter creates a new ErrAfter error.
//
// Parameters:
//   - quote: A flag indicating whether the previous value should be quoted.
//   - previous: The previous value associated with the error. If not provided, "at the start" is used.
//   - inner: The inner error that occurred. If not provided, "something went wrong" is used.
//
// Returns:
//   - error: The newly created ErrAfter error. Never returns nil.
//
// Format:
//
//	"after <previous>: <inner>"
func NewErrAfter(quote bool, previous *rune, inner error) error {
	return &ErrAfter{
		Quote:    quote,
		Previous: previous,
		Inner:    inner,
	}
}

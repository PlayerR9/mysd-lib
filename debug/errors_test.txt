package debug

import (
	"errors"
	"fmt"
	"testing"

	"github.com/PlayerR9/go-verify/test"
	gers "github.com/PlayerR9/mygo-lib/errors"
	"github.com/dustin/go-humanize"
)

// TestErrMsgOf tests the ErrMsgOf function.
func TestNewErrPrintFailed(t *testing.T) {
	type args struct {
		idx      int
		reason   error
		expected string
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			err := NewErrPrintFailed(args.idx, args.reason)
			if err == nil {
				t.Error("want error, got nil")
			} else {
				msg := err.Error()
				if msg != args.expected {
					t.Errorf("want %q, got %q", args.expected, msg)
				}
			}
		}
	})

	// 1. Test that, if reason is not provided, the reason of the error is the default error
	// in errors.DefaultErrMsg.
	// Here, the case idx is greater than or equal to zero is also tested.
	_ = tests.AddTest("no reason, idx >= 0", args{
		idx:      0,
		reason:   nil,
		expected: fmt.Sprintf("could not print the %s element: %s", humanize.Ordinal(1), gers.DefaultErr.Error()),
	})

	// 2. Test that, if reason is provided, the reason of the error is the provided reason.
	// Here, the case idx is greater than or equal to zero is also tested.
	_ = tests.AddTest("reason provided, idx >= 0", args{
		idx:      3,
		reason:   errors.New("foo"),
		expected: fmt.Sprintf("could not print the %s element: foo", humanize.Ordinal(4)),
	})

	// 3. Test that, if idx is less than zero, the error message is as expected.
	_ = tests.AddTest("idx < 0", args{
		idx:      -1,
		reason:   nil,
		expected: fmt.Sprintf("could not print the %s element: %s", humanize.Ordinal(0), gers.DefaultErr.Error()),
	})

	_ = tests.Run(t)
}

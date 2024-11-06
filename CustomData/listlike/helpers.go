package listlike

import (
	"errors"
	"fmt"
	"iter"
)

type SizeFunc struct {
	lastElem *any
	hasError bool
	result   uint
}

func (f SizeFunc) GetError() error {
	if f.lastElem == nil {
		return fmt.Errorf("automaton not in a terminal state")
	} else {
		return fmt.Errorf("function Size() is not implemented for %T", *f.lastElem)
	}
}

func (f SizeFunc) HasError() bool {
	return f.hasError
}

func (f *SizeFunc) Call(arg any) bool {
	if f == nil {
		panic("nil receiver")
	} else if f.lastElem != nil {
		return true
	}

	if arg == nil {
		f.result = 0
		*f.lastElem = nil

		return false
	}

	switch ll := arg.(type) {
	case []any:
		f.result = uint(len(ll))
	case *[]any:
		f.result = uint(len(*ll))
	case interface{ Size() uint }:
		f.result = ll.Size()
	default:
		f.hasError = true
	}

	*f.lastElem = arg

	return false
}

func (f *SizeFunc) Reset() {
	if f == nil {
		return
	}

	f.hasError = false
	f.result = 0
	f.lastElem = nil
}

func Size() *SizeFunc {
	return &SizeFunc{
		lastElem: nil,
		hasError: false,
		result:   0,
	}
}

func (f SizeFunc) Result() uint {
	return f.result
}

type Aut[O any] interface {
	HasError() bool
	GetError() error
	Call(arg any) bool
	Reset()
	Result() O
}

var (
	ErrNoAutomaton      error
	ErrBadlyImplemented error
)

func init() {
	ErrNoAutomaton = errors.New("no automaton was provided")
	ErrBadlyImplemented = errors.New("automaton is implemented incorrectly")
}

// ERROR checks if the given automaton has an error and returns it.
//
// Parameters:
//   - aut: The automaton to check for an error.
//
// Returns:
//   - error: The error associated with the automaton if it has one, or nil if there is no error.
//
// Panics:
//   - ErrNoAutomaton: If the provided automaton is nil.
func ERROR[O any](aut Aut[O]) error {
	if aut == nil {
		panic(ErrNoAutomaton)
	}

	if !aut.HasError() {
		return nil
	} else {
		return aut.GetError()
	}
}

// TRY calls the given automaton with the given argument and returns its result if there is no error.
// If there is an error, it calls the given error handler with the error and returns its result.
// If the error handler is nil, it will panic with the error.
//
// Parameters:
//   - aut: The automaton to call.
//   - arg: The argument to pass to the automaton.
//   - handle: The error handler to call if there is an error.
//
// Returns:
//   - O: The result of either the automaton or the error handler.
//
// Panics:
//   - ErrNoAutomaton: If the provided automaton is nil.
//   - error: If there is an error and the error handler is nil.
func TRY[O any](aut Aut[O], arg any, handle func(err error) O) O {
	if aut == nil {
		panic(ErrNoAutomaton)
	}

	_ = aut.Call(arg)

	if !aut.HasError() {
		return aut.Result()
	}

	err := aut.GetError()

	if handle == nil {
		panic(err)
	}

	return handle(err)
}

// CALL executes the given automaton with the provided argument and returns its result.
// If the automaton reaches a terminal state, it resets and tries again. If the automaton
// is implemented incorrectly, it panics.
//
// Parameters:
//   - aut: The automaton to execute.
//   - arg: The argument to pass to the automaton.
//
// Returns:
//   - O: The result of the automaton execution.
//
// Panics:
//   - ErrNoAutomaton: If the provided automaton is nil.
//   - ErrBadlyImplemented: If the automaton reaches a terminal state on the second call.
//   - error: If the automaton has an error.
func CALL[O any](aut Aut[O], arg any) O {
	if aut == nil {
		panic(ErrNoAutomaton)
	}
	defer aut.Reset()

	isDone := aut.Call(arg)
	if !isDone {
		// Successfully called the automaton.
		if aut.HasError() {
			panic(aut.GetError())
		}

		return aut.Result()
	}

	// The automaton is in a terminal state; try to recall it.
	aut.Reset()

	isDone = aut.Call(arg)
	if isDone {
		panic(ErrBadlyImplemented)
	}

	if aut.HasError() {
		panic(aut.GetError())
	}

	return aut.Result()
}

func EACH[O any](aut Aut[O], arg any) iter.Seq[O] {
	if aut == nil {
		panic(ErrNoAutomaton)
	}

	return func(yield func(O) bool) {
		defer aut.Reset()

		var isDone bool

		for !isDone {
			isDone = aut.Call(arg)

			if aut.HasError() {
				panic(aut.GetError())
			}

			if !yield(aut.Result()) {
				return
			}
		}

		if aut.HasError() {
			panic(aut.GetError())
		}

		_ = yield(aut.Result())
	}
}

func IsEmpty(ll any) bool {
	if ll == nil {
		return true
	}

	switch ll := ll.(type) {
	case []any:
		return len(ll) == 0
	case *[]any:
		return len(*ll) == 0
	case interface{ IsEmpty() bool }:
		return ll.IsEmpty()
	case interface{ Size() int }:
		return ll.Size() == 0
	default:
		panic(fmt.Sprintf("IsEmpty() is not implemented for %T", ll))
	}
}

func Reset(ll any) {
	if ll == nil {
		return
	}

	switch ll := ll.(type) {
	case []any:
		clear(ll)
	case *[]any:
		clear(*ll)
		*ll = nil
	case interface{ Reset() }:
		ll.Reset()
	default:
		panic(fmt.Sprintf("Reset() is not implemented for %T", ll))
	}
}

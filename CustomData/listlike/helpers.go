package listlike

import (
	"errors"
	"fmt"
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

type Aut interface {
	HasError() bool
	GetError() error
	Call(arg any) bool
	Reset()
}

var (
	ErrNoAutomaton error
)

func init() {
	ErrNoAutomaton = errors.New("no automaton was provided")
}

func ERROR(aut Aut) error {
	if aut == nil {
		return ErrNoAutomaton
	} else if !aut.HasError() {
		return nil
	} else {
		return aut.GetError()
	}
}

func CALL[O any](aut interface {
	Result() O

	Aut
}, arg any) O {
	if aut == nil {
		panic(ErrNoAutomaton)
	}
	defer aut.Reset()

	ok := aut.Call(arg)
	if ok {
		aut.Reset()

		ok = aut.Call(arg)
		if !ok {
			panic(errors.New("automaton is implemented incorrectly"))
		}
	}

	if aut.HasError() {
		panic(aut.GetError())
	}

	return aut.Result()
}

func FOR(f interface {
	Result() uint

	Aut
}, arg any) uint {
	if f == nil {
		return 0, ErrNoAutomaton
	}
	defer f.Reset()

	isDone := f.Call(arg)
	if isDone {
		f.Reset()

		isDone = f.Call(arg)
		if !isDone {
			panic(errors.New("automaton is implemented incorrectly"))
		}
	}

	for {
		isDone := f.Call(arg)
		if isDone {
			break
		}
	}

	res := f.Result()
	return res, ERROR(f)
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

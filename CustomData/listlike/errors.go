package listlike

import "errors"

var (
	// ErrEmptyStack occurs when a pop or peek operation is called on an empty stack.
	// This error can be checked with the == operator.
	//
	// Format:
	// 	"empty stack"
	ErrEmptyStack error

	// ErrEmptyQueue occurs when a pop or peek operation is called on an empty queue.
	// This error can be checked with the == operator.
	//
	// Format:
	// 	"empty queue"
	ErrEmptyQueue error

	// ErrCannotPush occurs when a push operation is called on a refusable stack that
	// was not accepted nor refused yet. This error can be checked with the == operator.
	//
	// Format:
	// 	"cannot push elements: stack not accepted nor refused"
	ErrCannotPush error
)

func init() {
	ErrEmptyStack = errors.New("empty stack")
	ErrEmptyQueue = errors.New("empty queue")
	ErrCannotPush = errors.New("cannot push elements: stack not accepted nor refused")
}

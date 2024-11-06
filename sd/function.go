package sd

// Func is an interface for functions that can be used in the SD package.
type Func[I, O any] interface {
	// HasError checks if the function has an error.
	//
	// Returns:
	//   - bool: True if the function has an error, false otherwise.
	HasError() bool

	// GetError returns the error associated with the function.
	//
	// Returns:
	//   - error: The error associated with the function. Nil must not be expected even
	// 	when HasError() method returns false.
	GetError() error

	// Call calls the function with the given input.
	//
	// Parameters:
	//   - input: The input to pass to the function.
	//
	// Returns:
	//   - O: The output of the function.
	Call(input I) O
}

// baseFunc is a base implementation of the Func interface.
type baseFunc[I, O any] struct {
	// err is the error that is associated with the function.
	err error

	// fn is the function that is associated with the function. Never nil.
	fn func(err *error, input I) O
}

// HasError implements the Func interface.
func (f baseFunc[I, O]) HasError() bool {
	return f.err != nil
}

// GetError implements the Func interface.
func (f baseFunc[I, O]) GetError() error {
	return f.err
}

// Call implements the Func interface.
func (f *baseFunc[I, O]) Call(input I) O {
	res := f.fn(&f.err, input)
	return res
}

// NewFunc creates a new instance of a function that implements the Func interface.
//
// Parameters:
//   - fn: A function that takes an error pointer and an input of type I, returning an output of type O.
//     Must not be nil.
//
// Returns:
//   - Func[I, O]: A new instance of a function that implements the Func interface, or nil if fn is nil.
func NewFunc[I, O any](fn func(err *error, input I) O) Func[I, O] {
	if fn == nil {
		return nil
	}

	return &baseFunc[I, O]{
		err: nil,
		fn:  fn,
	}
}

// NewErrFunc creates a new instance of a function that implements the Func interface, which
// always returns the given error.
//
// Parameters:
//   - err: The error to return. If nil, the function returns nil.
//
// Returns:
//   - Func[I, O]: A new instance of a function that implements the Func interface, which
//     always returns the given error. If err is nil, the function returns nil.
func NewErrFunc[I, O any](err error) Func[I, O] {
	if err == nil {
		return nil
	}

	return &baseFunc[I, O]{
		err: err,
		fn:  func(err *error, input I) O { return *new(O) },
	}
}

// NewNoopFunc creates a new instance of a function that implements the Func interface, which
// does nothing and does not return an error.
//
// Returns:
//   - Func[I, O]: A new instance of a function that implements the Func interface, which
//     does nothing and does not return an error.
func NewNoopFunc[I, O any]() Func[I, O] {
	return &baseFunc[I, O]{
		err: nil,
		fn:  func(err *error, input I) O { return *new(O) },
	}
}

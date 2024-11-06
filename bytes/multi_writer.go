package bytes

import (
	"fmt"
	"io"

	"github.com/PlayerR9/mysd-lib/common"
)

// MultiWriter is a writer that writes multiple data to the underlying io.Writer. Useful
// for writing many bytes at once.
type MultiWriter struct {
	// w is the underlying io.Writer.
	w io.Writer

	// written is the number of bytes written so far.
	written int
}

// Write implements io.Writer.
func (mw *MultiWriter) Write(data []byte) (int, error) {
	if len(data) == 0 {
		return 0, nil
	} else if mw == nil {
		return 0, common.ErrNilReceiver
	}

	n, err := mw.w.Write(data)
	mw.written += n

	if err == nil && n != len(data) {
		err = io.ErrShortWrite
	}

	return n, err
}

// New creates a new MultiWriter for the given io.Writer.
//
// Parameters:
//   - w: The underlying io.Writer.
//
// Returns:
//   - *MultiWriter: The new MultiWriter.
//   - error: An error if w is nil.
func New(w io.Writer) (*MultiWriter, error) {
	if w == nil {
		return nil, common.NewErrNilParam("w")
	}

	return &MultiWriter{
		w: w,
	}, nil
}

// Written returns the total number of bytes written.
//
// Returns:
//   - int: The total number of bytes written.
func (w MultiWriter) Written() int {
	return w.written
}

// WriteBytes writes the data to the underlying io.Writer. This does the same thing as
// Write, but does not return the number of bytes written.
//
// Parameters:
//   - data: The data to write.
//
// Returns:
//   - error: An error if writing failed.
//
// Errors:
//   - io.ErrShortWrite: If the data is not fully written.
//   - any other error returned by the underlying io.Writer.
func (w *MultiWriter) WriteBytes(data []byte) error {
	if len(data) == 0 {
		return nil
	} else if w == nil {
		return common.ErrNilReceiver
	}

	n, err := w.w.Write(data)
	w.written += n

	if err == nil && n != len(data) {
		err = io.ErrShortWrite
	}

	return err
}

// WriteNewline writes a newline character to the underlying io.Writer.
//
// Returns:
//   - error: An error if writing failed.
//
// Errors:
//   - io.ErrShortWrite: If the data is not fully written.
//   - any other error returned by the underlying io.Writer.
func (w *MultiWriter) WriteNewline() error {
	if w == nil {
		return common.ErrNilReceiver
	}

	n, err := w.w.Write(Newline)
	w.written += n

	if err == nil && n != NewlineLen {
		err = io.ErrShortWrite
	}

	return err
}

// WriteMany writes many data to the underlying io.Writer. This is a convenience
// function that acts like WriteBytes for many data in a more efficient way.
//
// Parameters:
//   - datas: The datas to write.
//
// Returns:
//   - error: An error if writing failed.
//
// Errors:
//   - io.ErrShortWrite: If the data is not fully written.
//   - any other error returned by the underlying io.Writer.
func (w *MultiWriter) WriteMany(datas ...[]byte) error {
	var total int

	for _, data := range datas {
		total += len(data)
	}

	if total == 0 {
		return nil
	} else if w == nil {
		return io.ErrShortWrite
	}

	final := make([]byte, total)
	var prev int

	for _, data := range datas {
		copy(final[prev:], data)
		prev += len(data)
	}

	n, err := w.w.Write(final)
	w.written += n

	if err == nil && n != total {
		err = io.ErrShortWrite
	}

	return err
}

// WriteString writes the given string to the underlying io.Writer.
//
// Parameters:
//   - str: The string to write.
//
// Returns:
//   - error: An error if writing failed.
//
// Errors:
//   - io.ErrShortWrite: If the data is not fully written.
//   - any other error returned by the underlying io.Writer.
func (w *MultiWriter) WriteString(str string) error {
	if str == "" {
		return nil
	} else if w == nil {
		return io.ErrShortWrite
	}

	data := []byte(str)

	n, err := w.w.Write(data)
	w.written += n

	if err == nil && n != len(data) {
		err = io.ErrShortWrite
	}

	return err
}

// Printf formats the given arguments according to the given format string and
// writes the result to the underlying io.Writer.
//
// Parameters:
//   - format: The format string.
//   - args: The arguments to format.
//
// Returns:
//   - error: An error if writing failed.
//
// Errors:
//   - io.ErrShortWrite: If the data is not fully written.
//   - any other error returned by the underlying io.Writer.
func (w *MultiWriter) Printf(format string, args ...any) error {
	data := []byte(fmt.Sprintf(format, args...))
	if len(data) == 0 {
		return nil
	}

	n, err := w.w.Write(data)
	w.written += n

	if err == nil && n != len(data) {
		err = io.ErrShortWrite
	}

	return err
}

// Print writes the given arguments to the underlying io.Writer.
//
// Parameters:
//   - args: The arguments to write.
//
// Returns:
//   - error: An error if writing failed.
//
// Errors:
//   - io.ErrShortWrite: If the data is not fully written.
//   - any other error returned by the underlying io.Writer.
func (w *MultiWriter) Print(args ...any) error {
	data := []byte(fmt.Sprint(args...))
	if len(data) == 0 {
		return nil
	}

	n, err := w.w.Write(data)
	w.written += n

	if err == nil && n != len(data) {
		err = io.ErrShortWrite
	}

	return err
}

// Println writes the given arguments to the underlying io.Writer, followed by a newline.
//
// Parameters:
//   - args: The arguments to write.
//
// Returns:
//   - error: An error if writing failed.
//
// Errors:
//   - io.ErrShortWrite: If the data is not fully written.
//   - any other error returned by the underlying io.Writer.
func (w *MultiWriter) Println(args ...any) error {
	data := []byte(fmt.Sprintln(args...))
	if len(data) == 0 {
		return nil
	}

	n, err := w.w.Write(data)
	w.written += n

	if err == nil && n != len(data) {
		err = io.ErrShortWrite
	}

	return err
}

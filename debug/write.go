package debug

import (
	"io"
	"iter"
	"log"

	common "github.com/PlayerR9/mysd-lib/common"
)

// LogSeq prints a title and a sequence of strings. Each element of the sequence is
// written in a new line.
//
// However, only the title uses the logger's flags since the sequence itself uses logger.Writer()
// writer instead.
//
// Parameters:
//   - logger: The logger to use.
//   - title: The title to print.
//   - seq: The sequence to print.
//
// Returns:
//   - error: An error if printing failed.
//
// Errors:
//   - errors.ErrNilParameter: when the logger is nil.
//   - ErrPrintFailed: when the logger failed to print an element.
//
// Behaviors:
//   - Prints the title if it is not empty. Same goes for seq when it is nil.
//   - To add an empty line at the end, use _ = yield("\n") in seq.
func LogSeq(logger *log.Logger, title string, seq iter.Seq[string]) error {
	if logger == nil {
		return common.NewErrNilParam("logger")
	}

	if title != "" {
		logger.Println(title)
	}

	if seq == nil {
		return nil
	}

	w := logger.Writer()

	var idx int

	for str := range seq {
		data := []byte(str + "\n")

		n, err := w.Write(data)
		if err != nil {
			return NewErrPrintFailed(idx, err)
		} else if n < len(data) {
			return NewErrPrintFailed(idx, io.ErrShortWrite)
		}

		idx++
	}

	return nil
}

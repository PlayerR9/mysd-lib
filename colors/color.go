package colors

import (
	"io"
	"strconv"

	"github.com/PlayerR9/mysd-lib/common"
)

// Color represents an RGB color.
type Color struct {
	// Red is the red component of the RGB color.
	Red uint8

	// Green is the green component of the RGB color.
	Green uint8

	// Blue is the blue component of the RGB color.
	Blue uint8
}

// New creates a new RGB color from the given red, green, and blue values.
//
// Parameters:
//   - rule: The red component of the RGB color.
//   - green: The green component of the RGB color.
//   - blue: The blue component of the RGB color.
//
// Returns:
//   - Color: The new RGB color. Never returns nil.
func New(red, green, blue uint8) Color {
	return Color{
		Red:   red,
		Green: green,
		Blue:  blue,
	}
}

// Bytes returns the bytes of the color.
//
// Returns:
//   - []byte: The bytes of the color. Never returns nil.
func (c Color) Bytes() []byte {
	return []byte{
		c.Red,
		c.Green,
		c.Blue,
	}
}

// Foreground writes an ANSI escape code to set the terminal's text color to the RGB
// values of the Color receiver.
//
// Parameters:
//   - w: The writer to which the ANSI escape code is written.
//
// Returns:
//   - error: An error if writing to the writer fails or if the writer is nil.
func (c Color) Foreground(w io.Writer) error {
	if w == nil {
		return common.NewErrNilParam("w")
	}

	data := []byte("\x1b[38;2;")
	data = strconv.AppendUint(data, uint64(c.Red), 10)
	data = append(data, ';')
	data = strconv.AppendUint(data, uint64(c.Green), 10)
	data = append(data, ';')
	data = strconv.AppendUint(data, uint64(c.Blue), 10)
	data = append(data, 'm')

	n, err := w.Write(data)
	if err == nil && n != len(data) {
		err = io.ErrShortWrite
	}

	return err
}

// Background writes an ANSI escape code to set the terminal's background color to the RGB
// values of the Color receiver.
//
// Parameters:
//   - w: The writer to which the ANSI escape code is written.
//
// Returns:
//   - error: An error if writing to the writer fails or if the writer is nil.
func (c Color) Background(w io.Writer) error {
	if w == nil {
		return common.NewErrNilParam("w")
	}

	data := []byte("\x1b[48;2;")
	data = strconv.AppendUint(data, uint64(c.Red), 10)
	data = append(data, ';')
	data = strconv.AppendUint(data, uint64(c.Green), 10)
	data = append(data, ';')
	data = strconv.AppendUint(data, uint64(c.Blue), 10)
	data = append(data, 'm')

	n, err := w.Write(data)
	if err == nil && n != len(data) {
		err = io.ErrShortWrite
	}

	return err
}

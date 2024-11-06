package tables

import (
	"iter"

	"github.com/PlayerR9/mysd-lib/common"
)

// Table is a boundless table. This means that any operation done
// to out-of-bounds cells will not cause any error.
type Table[T any] struct {
	// table is the underlying table.
	table [][]T

	// width is the width of the table.
	width int

	// height is the height of the table.
	height int
}

// NewTable creates a new Table with a width and height.
//
// Parameters:
//   - width: The width of the table.
//   - height: The height of the table.
//
// Returns:
//   - *Table: The new Table.
//   - error: If the table could not be created.
//
// Errors:
//   - errors.BadParameterError: If width or height is negative.
func NewTable[T any](width, height int) (*Table[T], error) {
	if width < 0 {
		return nil, common.NewErrBadParam("width", "must be non-negative")
	} else if height < 0 {
		return nil, common.NewErrBadParam("height", "must be non-negative")
	}

	table := make([][]T, 0, height)

	for i := 0; i < height; i++ {
		table = append(table, make([]T, width, width))
	}

	return &Table[T]{
		table:  table,
		width:  width,
		height: height,
	}, nil
}

// Height returns the height of the table.
//
// Returns:
//   - int: The height of the table.
func (t Table[T]) Height() int {
	return t.height
}

// Width returns the width of the table.
//
// Returns:
//   - int: The width of the table.
func (t Table[T]) Width() int {
	return t.width
}

// CellAt returns the cell at the specified position.
//
// Parameters:
//   - x: The x position of the cell.
//   - y: The y position of the cell.
//
// Returns:
//   - T: The cell at the specified position. The zero value if the position
//     is out of bounds.
func (t Table[T]) CellAt(x, y int) T {
	if x < 0 || x >= t.width || y < 0 || y >= t.height {
		return *new(T)
	}

	return t.table[y][x]
}

// ResizeWidth resizes the width of the table. The width is not
// resized if the receiver is nil or the new width is the same as the
// current width.
//
// Parameters:
//   - new_width: The new width of the table.
//
// Returns:
//   - error: If the table could not be resized.
//
// Errors:
//   - gers.BadParameterError: If new_width is negative.
func (t *Table[T]) ResizeWidth(new_width int) error {
	if t == nil {
		return common.ErrNilReceiver
	} else if new_width < 0 {
		return common.NewErrBadParam("new_width", "must be non-negative")
	}

	if new_width == t.width {
		return nil
	}

	if new_width < t.width {
		for i := 0; i < t.height; i++ {
			t.table[i] = t.table[i][:new_width:new_width]
		}
	} else {
		extension := make([]T, new_width-t.width)

		for i := 0; i < t.height; i++ {
			t.table[i] = append(t.table[i], extension...)
		}
	}

	return nil
}

// ResizeHeight resizes the height of the table. The height is not
// resized if the receiver is nil or the new height is the same as the
// current height.
//
// Parameters:
//   - new_height: The new height of the table.
//
// Returns:
//   - error: If the table could not be resized.
//
// Errors:
//   - gers.BadParameterError: If new_height is negative.
func (t *Table[T]) ResizeHeight(new_height int) error {
	if t == nil {
		return common.ErrNilReceiver
	} else if new_height < 0 {
		return common.NewErrBadParam("new_height", "must be non-negative")
	}

	if new_height == t.height {
		return nil
	}

	if new_height < t.height {
		t.table = t.table[:new_height:new_height]
	} else {
		for i := t.height; i < new_height; i++ {
			t.table = append(t.table, make([]T, t.width))
		}
	}

	return nil
}

// SetCellAt sets the cell at the specified position. The cell is not
// set if the receiver is nil or the position is out of bounds.
//
// Parameters:
//   - cell: The cell to set.
//   - x: The x position of the cell.
//   - y: The y position of the cell.
func (t *Table[T]) SetCellAt(cell T, x, y int) {
	if t == nil || y < 0 || y >= t.height || x < 0 || x >= t.width {
		return
	}

	t.table[y][x] = cell
}

// Row returns an iterator over the rows in the table.
//
// Returns:
//   - iter.Seq2[int, []T]: An iterator over the rows in the table. Never returns nil.
func (t Table[T]) Row() iter.Seq2[int, []T] {
	return func(yield func(int, []T) bool) {
		for i := 0; i < t.height; i++ {
			if !yield(i, t.table[i]) {
				return
			}
		}
	}
}

// Cleanup cleans up the table. Does nothing if the receiver is nil or if
// is already cleaned up.
func (t *Table[T]) Cleanup() {
	if t == nil {
		return
	}

	if len(t.table) > 0 {
		for i := 0; i < t.height; i++ {
			for j := 0; j < t.width; j++ {
				t.table[i][j] = *new(T)
			}

			t.table[i] = t.table[i][:0]
		}

		t.table = t.table[:0]
	}

	t.width = 0
	t.height = 0
}

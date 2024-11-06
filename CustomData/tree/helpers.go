package tree

import (
	"io"

	"github.com/PlayerR9/mysd-lib/common"
)

var (
	nonterminal_data, terminal_data, indent_pipe_data, indent_empty_data []byte
	nonterminal_size, terminal_size                                      int
)

func init() {
	nonterminal_data = []byte("├── ")
	nonterminal_size = len(nonterminal_data)

	terminal_data = []byte("└── ")
	terminal_size = len(terminal_data)

	indent_pipe_data = []byte("│   ")

	indent_empty_data = []byte("    ")
}

// wtInfo is used by WriteTree to print the token tree in a form that is easier to read
// and understand.
type wtInfo struct {
	// indent is the current indentation.
	indent []byte

	// has_next is true if there is a next sibling, false otherwise.
	has_next bool

	// is_first is true if this is the root node, false otherwise.
	is_first bool
}

// printFn is a helper function to print the tree.
//
// Parameters:
//   - w: The writer to use.
//   - node: The node to print.
//   - info_table: A map of nodes to their MakeWTInfo.
//
// Returns:
//   - error: An error if printing failed.
//
// Errors:
//   - any error returned by io.Writer.Write
var printFn func(w io.Writer, node *Node, info_table map[*Node]wtInfo) error

func init() {
	printFn = func(w io.Writer, node *Node, info_table map[*Node]wtInfo) error {
		info := info_table[node]

		if !info.is_first {
			data := append([]byte("\n"), info.indent...)
			expected_size := len(data)

			if !info.has_next {
				data = append(data, terminal_data...)
				expected_size += terminal_size
			} else {
				data = append(data, nonterminal_data...)
				expected_size += nonterminal_size
			}

			n, err := w.Write(data)
			if err != nil {
				return err
			} else if n != expected_size {
				return io.ErrShortWrite
			}
		}

		data := []byte(node.String())
		n, err := w.Write(data)
		if err != nil {
			return err
		} else if n != len(data) {
			return io.ErrShortWrite
		}

		if node.FirstChild == nil {
			return nil
		}

		var new_indent []byte

		if info.is_first {
			new_indent = nil
		} else if !info.has_next {
			new_indent = append([]byte(info.indent), indent_empty_data...)
		} else {
			new_indent = append([]byte(info.indent), indent_pipe_data...)
		}

		for c := node.FirstChild; c.NextSibling != nil; c = c.NextSibling {
			info_table[c] = wtInfo{
				indent:   new_indent,
				has_next: true,
				is_first: false,
			}
		}

		info_table[node.LastChild] = wtInfo{
			indent:   new_indent,
			has_next: false,
			is_first: false,
		}

		return nil
	}
}

// Equals checks if two trees are equal.
//
// The two trees are considered equal if they have the same number of nodes and the same structure.
// The contents of the nodes are compared with the Equals method of the node type.
//
// Parameters:
//   - other: The other tree to compare with.
//
// Returns:
//   - bool: True if the two trees are equal, false otherwise.
func Equals(tree, other *Tree) bool {
	if tree == nil || other == nil || tree.Size() != other.Size() {
		return false
	}

	queue := [][2]*Node{{tree.Root(), other.Root()}}

	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]

		if !first[0].Equals(first[1]) {
			return false
		}

		var s1, s2 []*Node

		for c := first[0].FirstChild; c != nil; c = c.NextSibling {
			s1 = append(s1, c)
		}

		for c := first[1].FirstChild; c != nil; c = c.NextSibling {
			s2 = append(s2, c)
		}

		if len(s1) != len(s2) {
			return false
		}

		slice := make([][2]*Node, 0, len(s1))

		for i, elem := range s1 {
			slice = append(slice, [2]*Node{elem, s2[i]})
		}

		queue = append(queue, slice...)
	}

	return true
}

// Get returns the value of the node if it is of type T, or an error if the node is nil or the
// information of the node is not of type T.
//
// Parameters:
//   - node: The node to get the value of.
//
// Returns:
//   - T: The value of the node if the node is not nil and the information of the node is of type T.
//   - error: An error if the node is nil, or the information of the node is not of type T.
//
// Errors:
//   - common.ErrBadParam: If the node is nil.
//   - common.ErrInvalidType: If the information of the node is not of type T, including if node.Info is nil.
func Get[T Infoer](node *Node) (T, error) {
	if node == nil {
		return *new(T), common.NewErrNilParam("node")
	}

	info := node.Info
	if info == nil {
		return *new(T), common.NewErrInvalidType(nil, *new(T))
	}

	v, ok := info.(T)
	if !ok {
		return *new(T), common.NewErrInvalidType(info, v)
	}

	return v, nil
}

// MustGet returns the value of the node if it is of type T, or panics if the node is nil or the
// information of the node is not of type T.
//
// Parameters:
//   - node: The node to get the value of.
//
// Panics:
//   - common.ErrNilParam: If the node is nil.
//   - common.ErrInvalidType: If the information of the node is not of type T, including if node.Info is nil.
//
// Returns:
//   - T: The value of the node if the node is not nil and the information of the node is of type T.
func MustGet[T Infoer](node *Node) T {
	if node == nil {
		panic(common.NewErrNilParam("node"))
	}

	info := node.Info
	if info == nil {
		panic(common.NewErrInvalidType(nil, *new(T)))
	}

	v, ok := info.(T)
	if !ok {
		panic(common.NewErrInvalidType(info, v))
	}

	return v
}

package tree

import (
	"strings"
)

// Tree is a tree data structure.
type Tree struct {
	// root is the root node of the tree.
	root *Node

	// leaves are the leaf nodes of the tree.
	leaves []*Node

	// size is the number of nodes in the tree.
	size int
}

// String implements the fmt.Stringer interface.
func (t Tree) String() string {
	var builder strings.Builder

	table := make(map[*Node]wtInfo)
	table[t.root] = wtInfo{
		indent:   nil,
		has_next: false,
		is_first: true,
	}

	fn := func(node *Node) error {
		err := printFn(&builder, node, table)
		return err
	}

	err := View.DFS(&t, fn)
	if err != nil {
		panic(err)
	}

	return builder.String()
}

// NewTree creates a new Tree given the root node.
//
// Parameters:
//   - root: The root node of the tree.
//
// Returns:
//   - *Tree[T]: The tree. Nil if root is nil.
func NewTree(root *Node) *Tree {
	if root == nil {
		return nil
	}

	var leaves []*Node
	stack := []*Node{root}
	var size int

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		size++

		if top.FirstChild == nil {
			leaves = append(leaves, top)

			continue
		}

		for c := top.LastChild; c != nil; c = c.PrevSibling {
			stack = append(stack, c)
		}
	}

	return &Tree{
		root:   root,
		leaves: leaves,
		size:   size,
	}
}

// Root returns the root node of the tree.
//
// Returns:
//   - *Node: The root node of the tree. Never returns nil.
func (t Tree) Root() *Node {
	return t.root
}

// Size returns the number of nodes in the tree.
//
// Returns:
//   - int: The number of nodes in the tree.
func (t Tree) Size() int {
	return t.size
}

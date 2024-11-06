package tree

import (
	"github.com/PlayerR9/mysd-lib/common"
	"github.com/PlayerR9/mysd-lib/slices"
)

// Node is the node in the tree.
type Node struct {
	// Parent, FirstChild, LastChild, NextSibling, and PrevSibling are pointers of
	// the node.
	Parent, FirstChild, LastChild, NextSibling, PrevSibling *Node

	// Info contains the position, type, and data of the node.
	Info Infoer
}

// String implements the fmt.Stringer interface.
func (n Node) String() string {
	if n.Info == nil {
		return "Node[nil]"
	} else {
		return "Node[" + n.Info.String() + "]"
	}
}

// NewNode creates a new node with the given information.
//
// Parameters:
//   - info: The information of the node.
//
// Returns:
//   - *Node: The new node. Never returns nil.
func NewNode(info Infoer) *Node {
	return &Node{
		Info: info,
	}
}

// Equals checks if the node is equal to another node.
//
// The two nodes are considered equal if the information stored in the nodes
// are equal.
//
// Parameters:
//   - other: The other node to compare with.
//
// Returns:
//   - bool: True if the two nodes are equal, false otherwise.
func (n Node) Equals(other *Node) bool {
	return other != nil && other.Info != nil && n.Info.Equals(other.Info)
}

// link_nodes links the given children nodes to the specified parent node,
// setting up the parent, next sibling, and previous sibling pointers.
//
// Parameters:
//   - parent: The parent node to which the children will be linked.
//   - children: A slice of nodes to be linked as children.
//
// Returns:
//   - []*Node: The linked children nodes.
func link_nodes(parent *Node, children []*Node) []*Node {
	for _, c := range children {
		c.Parent = parent
	}

	prev := children[0]

	for _, c := range children[1:] {
		prev.NextSibling = c
		c.PrevSibling = prev
		prev = c
	}

	return children
}

// PrependChildren adds the given children nodes to the beginning of the current node's children list.
//
// Parameters:
//   - children: Variadic parameter of type Node representing the children to be added.
//
// Returns:
//   - error: Returns an error if the operation fails or if the receiver is nil, otherwise returns nil.
func (n *Node) PrependChildren(children ...*Node) error {
	slices.RejectNils(&children)
	if len(children) == 0 {
		return nil
	} else if n == nil {
		return common.ErrNilReceiver
	}

	children = link_nodes(n, children)

	if n.FirstChild == nil {
		n.LastChild = children[len(children)-1]
	} else {
		n.FirstChild.PrevSibling = children[len(children)-1]
		children[len(children)-1].NextSibling = n.FirstChild
	}

	n.FirstChild = children[0]

	return nil
}

// AppendChildren adds the given children nodes to the end of the current node's children list.
//
// Parameters:
//   - children: Variadic parameter of type Node representing the children to be appended.
//
// Returns:
//   - error: Returns an error if the operation fails or if the receiver is nil, otherwise returns nil.
func (n *Node) AppendChildren(children ...*Node) error {
	slices.RejectNils(&children)
	if len(children) == 0 {
		return nil
	} else if n == nil {
		return common.ErrNilReceiver
	}

	children = link_nodes(n, children)

	if n.LastChild == nil {
		n.FirstChild = children[0]
	} else {
		n.LastChild.NextSibling = children[0]
		children[0].PrevSibling = n.LastChild
	}

	n.LastChild = children[len(children)-1]

	return nil
}

package tree

import "errors"

// viewT is for internal use only.
type viewT struct{}

// View is the namespace for view functions.
var View viewT

func init() {
	View = viewT{}
}

var (
	// ErrEarlyExit occurs when the traversal should stop early without error.
	// This can be checked with the == operator.
	//
	// Format:
	//
	// 	"early exit"
	ErrEarlyExit error
)

func init() {
	ErrEarlyExit = errors.New("early exit")
}

// ViewElem is an element of a view.
type ViewElem struct {
	// node is the node of the view.
	node *Node

	// seen is a flag that indicates if the node has been visited.
	seen bool
}

// NewViewElem returns a new view element with the given node and seen set to false.
//
// Parameters:
//   - node: The node of the view.
//
// Returns:
//   - *ViewElem: The new view element. Never returns nil.
func NewViewElem(node *Node) *ViewElem {
	return &ViewElem{
		node: node,
		seen: false,
	}
}

// VisitFn is a function that visits a node.
//
// Parameters:
//   - node: The node to visit.
//
// Returns:
//   - error: An error that indicates if the traversal should immediately stop.
//
// Errors:
//   - ErrEarlyExit: The traversal should stop early without error.
type VisitFn func(node *Node) error

// PreorderView performs a preorder traversal without using recursion of the tree; stopping at the
// first error encountered.
//
// Parameters:
//   - tree: The tree to traverse.
//   - visit: The function to call for each node in the tree.
//
// Returns:
//   - error: The first error encountered during the traversal, or nil if the traversal was successful.
//
// Behaviors:
//   - If tree is nil or if visit is nil, then the traversal won't be performed but a nil error will
//     be returned.
//   - Nil children will be ignored.
func (viewT) Preorder(tree *Tree, visit VisitFn) error {
	if tree == nil || visit == nil {
		return nil
	}

	root := tree.Root()

	stack := []*Node{root}
	var err error

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		err = visit(top)
		if err != nil {
			break
		}

		for c := top.LastChild; c != nil; c = c.PrevSibling {
			stack = append(stack, c)
		}
	}

	if err == ErrEarlyExit {
		err = nil
	}

	return err
}

// PostorderView performs a post-order traversal of the tree without using recursion.
// It processes each node after its children have been processed.
//
// Parameters:
//   - tree: The tree to traverse. If it is nil, the function returns immediately without error.
//   - visit: A function to call for each node visited. If it is nil, the function returns immediately
//     without error.
//
// Returns:
//   - error: The first error encountered during the traversal, or nil if the traversal was successful.
//
// Behaviors:
//   - If the tree is nil or if the visit function is nil, then the traversal won't be performed but a
//     nil error will be returned.
//   - Nil children will be ignored.
//   - Stops traversal early if the visit function returns an error, except for ErrEarlyExit which is
//     ignored.
func (viewT) Postorder(tree *Tree, visit VisitFn) error {
	if tree == nil || visit == nil {
		return nil
	}

	stack := []*ViewElem{
		NewViewElem(tree.Root()),
	}

	var err error

	for len(stack) > 0 {
		top := stack[len(stack)-1]

		if top.seen {
			stack = stack[:len(stack)-1]

			err = visit(top.node)
			if err != nil {
				break
			}

			continue
		}

		stack[len(stack)-1].seen = true

		var elems []*ViewElem

		for c := top.node.LastChild; c != nil; c = c.PrevSibling {
			elem := NewViewElem(c)
			elems = append(elems, elem)
		}

		stack = append(stack, elems...)
	}

	if err == ErrEarlyExit {
		err = nil
	}

	return err
}

// inorderView is a helper function that performs an in-order traversal of a tree node with recursion.
// It visits the left subtree, then the node itself, and finally the right subtree.
//
// Parameters:
//   - node: The current node to process. It should implement the TreeNoder interface.
//   - visit: A function to call for each node visited. If it is nil, traversal won't occur.
//
// Returns:
//   - error: The first error encountered during the traversal, or nil if successful.
//
// Behaviors:
//   - If node is a leaf (no children), it will directly call visit.
//   - If the visit function returns an error, traversal stops immediately.
//   - If the visit function returns ErrEarlyExit, it is ignored and traversal continues.
func inorderView(node *Node, visit VisitFn) error {
	var children []*Node

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		children = append(children, c)
	}

	if len(children) == 0 {
		err := visit(node)
		return err
	}

	mid := len(children) / 2

	left, right := children[:mid], children[mid:]

	for _, c := range left {
		err := inorderView(c, visit)
		if err != nil {
			return err
		}
	}

	err := visit(node)
	if err != nil {
		return err
	}

	for _, c := range right {
		err := inorderView(c, visit)
		if err != nil {
			return err
		}
	}

	return nil
}

// Inorder performs an in-order traversal of the tree by using recursion.
// It visits the left subtree, then the node itself, and finally the right subtree.
//
// Parameters:
//   - tree: The tree to traverse. If it is nil, the function returns immediately without error.
//   - visit: A function to call for each node visited. If it is nil, the function returns immediately
//     without error.
//
// Returns:
//   - error: The first error encountered during the traversal, or nil if the traversal was successful.
//
// Behaviors:
//   - If the tree is nil or if the visit function is nil, then the traversal won't be performed but a
//     nil error will be returned.
//   - Nil children will be ignored.
//   - If the visit function returns an error, traversal stops immediately.
//   - If the visit function returns ErrEarlyExit, it is ignored and traversal continues.
func (viewT) Inorder(tree *Tree, visit VisitFn) error {
	if tree == nil || visit == nil {
		return nil
	}

	err := inorderView(tree.Root(), visit)
	if err == ErrEarlyExit {
		err = nil
	}

	return err
}

// BFS performs a breadth-first traversal of the tree without using recursion.
// It processes each node before its children are processed.
//
// Parameters:
//   - tree: The tree to traverse. If it is nil, the function returns immediately without error.
//   - visit: A function to call for each node visited. If it is nil, the function returns immediately
//     without error.
//
// Returns:
//   - error: The first error encountered during the traversal, or nil if the traversal was successful.
//
// Behaviors:
//   - If the tree is nil or if the visit function is nil, then the traversal won't be performed but a
//     nil error will be returned.
//   - Nil children will be ignored.
//   - Stops traversal early if the visit function returns an error, except for ErrEarlyExit which is
//     ignored.
func (viewT) BFS(tree *Tree, visit VisitFn) error {
	if tree == nil || visit == nil {
		return nil
	}

	queue := []*Node{tree.Root()}
	var err error

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		err = visit(top)
		if err != nil {
			break
		}

		for c := top.FirstChild; c != nil; c = c.NextSibling {
			queue = append(queue, c)
		}
	}

	if err == ErrEarlyExit {
		err = nil
	}

	return err
}

// DFS performs a depth-first traversal of the tree without using recursion; stopping at the
// first error encountered.
//
// Parameters:
//   - tree: The tree to traverse.
//   - visit: The function to call for each node in the tree.
//
// Returns:
//   - error: The first error encountered during the traversal, or nil if the traversal was successful.
//
// Behaviors:
//   - If the tree is nil or if the visit function is nil, then the traversal won't be performed but a
//     nil error will be returned.
//   - Nil children will be ignored.
//   - Stops traversal early if the visit function returns an error, except for ErrEarlyExit which is
//     ignored.
func (viewT) DFS(tree *Tree, visit VisitFn) error {
	if tree == nil || visit == nil {
		return nil
	}

	stack := []*ViewElem{
		NewViewElem(tree.Root()),
	}
	var err error

	for len(stack) > 0 {
		top := stack[len(stack)-1]

		if top.seen {
			stack = stack[:len(stack)-1]

			err = visit(top.node)
			if err != nil {
				break
			}

			continue
		}

		stack[len(stack)-1].seen = true

		var elems []*ViewElem

		for c := top.node.LastChild; c != nil; c = c.PrevSibling {
			elem := NewViewElem(c)
			elems = append(elems, elem)
		}

		stack = append(stack, elems...)
	}

	if err == ErrEarlyExit {
		err = nil
	}

	return err
}

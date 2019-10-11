package node

// Node is a recursive data structure where each node has an integer element (key) and a pointer to another node.
type Node struct {
	key  int
	next *Node
}

// NewNode returns a pointer to a new node initialised by an integer key and a pointer to the next node.
func NewNode(key int, next *Node) *Node {
	return &Node{
		key:  key,
		next: next,
	}
}

// Key returns the integer value of the node's key.
func (n *Node) Key() int {
	return n.key
}

// Key returns the pointer to the next node.
func (n *Node) Next() *Node {
	return n.next
}

// SetKey sets the key of the node to the input integer.
func (n *Node) SetKey(newKey int) {
	n.key = newKey
}

//SetNext sets the next node to the input node pointer.
func (n *Node) SetNext(newNext *Node) {
	n.next = newNext
}

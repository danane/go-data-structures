package bst

import "fmt"

//Node represents a single node that composes the tree.
type Node struct {
	key   int64
	left  *Node
	right *Node
}

//BinarySearchTree represents the tree.
type BinarySearchTree struct {
	root *Node
}

//NewBinarySearchTree creates an empty tree.
func NewBinarySearchTree() *BinarySearchTree {
	tree := &BinarySearchTree{}
	return tree
}

//Insert adds a new node to the tree.
func (tree *BinarySearchTree) Insert(v int64) {
	newNode := &Node{v, nil, nil}
	if tree.root != nil {
		insertNode(tree.root, newNode)
	} else {
		tree.root = newNode
	}
}

func insertNode(n, newNode *Node) {
	if n.key < newNode.key {
		if n.right == nil {
			n.right = newNode
		} else {
			insertNode(n.right, newNode)
		}

	} else {
		if n.left == nil {
			n.left = newNode
		} else {
			insertNode(n.left, newNode)
		}
	}
}

//PreOrderTraverse performs a pre-order traverse of the tree and prints the the keys of the nodes.
func (tree *BinarySearchTree) PreOrderTraverse() {
	fmt.Println("PRE-ORDER TRAVERSE----")
	if tree.root != nil {
		preOrderNode(tree.root)
	}
	defer fmt.Printf("\n\n")
}

func preOrderNode(n *Node) {
	if n != nil {
		fmt.Printf("%d ", n.key)
		preOrderNode(n.left)
		preOrderNode(n.right)
	}
}

//InOrderTraverse performs a in-order traverse of the tree and prints the the keys of the nodes.
func (tree *BinarySearchTree) InOrderTraverse() {
	fmt.Println("IN-ORDER TRAVERSE----")
	if tree.root != nil {
		inOrderNode(tree.root)
	}
	defer fmt.Printf("\n\n")
}

func inOrderNode(n *Node) {
	if n != nil {
		inOrderNode(n.left)
		fmt.Printf("%d ", n.key)
		inOrderNode(n.right)
	}
}

//PostOrderTraverse performs a post-order traverse of the tree and prints the the keys of the nodes.
func (tree *BinarySearchTree) PostOrderTraverse() {
	fmt.Println("POST-ORDER TRAVERSE----")
	if tree.root != nil {
		postOrderNode(tree.root)
	}
	defer fmt.Printf("\n\n")
}

func postOrderNode(n *Node) {
	if n != nil {
		postOrderNode(n.left)
		postOrderNode(n.right)
		fmt.Printf("%d ", n.key)
	}
}

//Min returns the minimum key in the tree.
func (tree *BinarySearchTree) Min() int64 {
	if tree.root != nil {
		return minNode(tree.root)
	}
	return -1
}

func minNode(n *Node) int64 {
	curr := n

	for curr.left != nil {
		curr = curr.left
	}
	return curr.key
}

//Max returns the maximum key in the tree.
func (tree *BinarySearchTree) Max() int64 {
	if tree.root != nil {
		curr := tree.root

		for curr.right != nil {
			curr = curr.right
		}
		return curr.key
	}
	return -1
}

//Search returns true if the key is in the tree, false otherwise.
func (tree *BinarySearchTree) Search(k int64) bool {
	return search(tree.root, k)
}

func search(n *Node, k int64) bool {
	if n == nil {
		return false
	}

	if n.key < k {
		return search(n.right, k)
	}
	if n.key > k {
		return search(n.left, k)
	}
	return true
}

//Remove deletes the node with key v from the tree.
func (tree *BinarySearchTree) Remove(v int64) *Node {
	if tree.root != nil {
		return removeNode(tree.root, v)
	}
	return nil
}

func removeNode(n *Node, v int64) *Node {
	if n.key < v {
		n.right = removeNode(n.right, v)
	} else if n.key > v {
		n.left = removeNode(n.left, v)
	} else {
		//key == v
		if n.left == nil {
			return n.right
		} else if n.right == nil {
			return n.left
		}

		minKey := minNode(n)
		n = removeNode(n, minKey)
		n.key = minKey
	}
	return n
}

func (tree *BinarySearchTree) String() {
	fmt.Println("---------------------STRUCTURE OF THE BST---------------------------")
	stringify(tree.root, 0)
	fmt.Printf("--------------------------------------------------------------------\n\n")
}

func stringify(n *Node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.right, level)
		fmt.Printf(format+"%d\n", n.key)
		stringify(n.left, level)
	}
}

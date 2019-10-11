package nodelist

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/danane/go-data-structures/node"
)

const (
	ErrEmptyList   = nodeListError("empty list")
	ErrKeyNotFound = nodeListError("key not found")
)

type nodeListError string

func (nle nodeListError) Error() string {
	return string(nle)
}

// NodeList implements a non-thread safe version of the list.List data structure by using a linked list of node.Node.
type NodeList struct {
	head   *node.Node
	length int
}

// Length provides the total length of the list.
func (nl *NodeList) Length() int {
	return nl.length
}

//Insert appends a new node to the NodeList with key equals to e.
func (nl *NodeList) Insert(e int) {
	defer func() { nl.length++ }()

	newNode := node.NewNode(e, nil)

	if nl.head == nil {
		nl.head = newNode
		return
	}

	var n *node.Node
	for n = nl.head; n.Next() != nil; n = n.Next() {
	}

	n.SetNext(newNode)
}

// Remove removes the first element with key equals to e and returns an error if any.
func (nl *NodeList) Remove(e int) error {
	removed := false
	defer func() {
		if removed {
			nl.length--
		}
	}()

	if nl.head == nil {
		return fmt.Errorf("could not remove: %w", ErrEmptyList)
	}

	var (
		n    *node.Node
		prec *node.Node
	)

	for n = nl.head; n.Next() != nil && n.Key() != e; prec, n = n, n.Next() {
	}

	if n.Key() != e {
		return fmt.Errorf("could not remove: %w", ErrKeyNotFound)
	}

	if n == nl.head {
		nl.head = n.Next()
		removed = true
		return nil
	}

	prec.SetNext(n.Next())
	removed = true

	return nil
}

// String returns a string representation of the NodeList.
func (nl *NodeList) String() string {
	var sb strings.Builder
	n := nl.head
	sb.WriteString("<")

	if n != nil {
		sb.WriteString(strconv.Itoa(n.Key()))
		n = n.Next()
	}

	for n != nil {
		sb.WriteString(" -> ")
		sb.WriteString(strconv.Itoa(n.Key()))
		n = n.Next()
	}

	sb.WriteString(">")

	return sb.String()
}

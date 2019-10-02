package bst

import (
	"reflect"
	"testing"
)

func generateTree(values []int64) *BinarySearchTree {
	root := &Node{values[0], nil, nil}

	root.left = &Node{values[1], nil, nil}
	root.right = &Node{values[2], nil, nil}
	return &BinarySearchTree{root}
}

func inOrderSlice(node *Node) []int64 {
	if node != nil {
		leftSlice := inOrderSlice(node.left)
		rightSlice := inOrderSlice(node.right)
		rightSlice = append([]int64{node.key}, rightSlice...)
		toReturn := append(leftSlice, rightSlice...)
		return toReturn
	}
	return []int64{}
}

func TestBinarySearchTree_Insert(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		tree *BinarySearchTree
		v    int64
		want []int64
	}{
		{name: "InsertEmpty", tree: NewBinarySearchTree(), v: 5, want: []int64{5}},
		{name: "InsertLeftSubtree", tree: generateTree([]int64{4, 3, 5}), v: 2, want: []int64{2, 3, 4, 5}},
		{name: "InsertRightSubtree", tree: generateTree([]int64{4, 3, 5}), v: 7, want: []int64{3, 4, 5, 7}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.tree.Insert(tc.v)

			got := inOrderSlice(tc.tree.root)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("want %v got %v", tc.want, got)
			}

		})
	}
}

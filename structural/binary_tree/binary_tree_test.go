package binarytree

import (
	"fmt"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	tree := &Tree{LeafValue: 10}
	fmt.Println(tree)
	
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(2)
	tree.Insert(8)
	tree.Insert(12)
	tree.Insert(18)

	fmt.Printf("In-order traversal (sorted): %v\n", tree.InOrder())

}



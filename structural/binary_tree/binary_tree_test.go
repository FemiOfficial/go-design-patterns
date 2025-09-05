package binarytree

import (
	"fmt"
	"testing"
	"reflect"
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
	if !reflect.DeepEqual(tree.InOrder(), []int{2, 5, 8, 10, 12, 15, 18}) {
		t.Errorf("In-order traversal is not correct")
	}

}



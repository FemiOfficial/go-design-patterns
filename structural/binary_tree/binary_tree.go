package binarytree

type Tree struct {
	Left *Tree
	Right *Tree
	LeafValue int
}

func (t *Tree) Insert(value int) {
	if t.LeafValue == 0 {
		// Empty tree, set the root value
		t.LeafValue = value
		return
	}
	
	if value < t.LeafValue {
		// Insert in left subtree
		if t.Left == nil {
			t.Left = &Tree{LeafValue: value}
		} else {
			t.Left.Insert(value)
		}
	} else if value > t.LeafValue {
		// Insert in right subtree
		if t.Right == nil {
			t.Right = &Tree{LeafValue: value}
		} else {
			t.Right.Insert(value)
		}
	}
	// If value == t.LeafValue, do nothing (no duplicates)
}

func (t *Tree) InOrder() []int {
	if t == nil {
		return []int{}
	}
	
	var result []int
	result = append(result, t.Left.InOrder()...)
	result = append(result, t.LeafValue)
	result = append(result, t.Right.InOrder()...)
	return result
}
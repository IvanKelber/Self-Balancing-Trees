package redblack

import (
	"fmt"
)

type Color bool

var Red Color = true
var Black Color = false

type RBNode struct {
	val    int
	isNil  bool
	color  Color
	parent *RBNode
	left   *RBNode
	right  *RBNode
}

func nilRBNode() *RBNode {
	return &RBNode{-1, true, Black, nil, nil, nil}
}

func RedBlackTree(val int) *RBNode {
	left := nilRBNode()
	right := nilRBNode()
	root := &RBNode{val, false, Black, nil, left, right}
	left.parent = root
	right.parent = root
	return root
}

func (rb *RBNode) String() string {
	return fmt.Sprintf("%v", rb.GetList())
}

func (rb *RBNode) GetList() []int {
	if rb.isNil {
		return []int{}
	}
	return append(append(rb.left.GetList(), rb.val), rb.right.GetList()...)
}

func (rb *RBNode) Contains(val int) bool {
	cur := rb
	for !cur.isNil {
		if cur.val > val {
			cur = cur.left
		} else if cur.val < val {
			cur = cur.right
		} else {
			return true
		}
	}
	return false
}

func (rb *RBNode) Insert(val int) bool {
	cur := rb
	for !cur.isNil {
		if cur.val > val {
			cur = cur.left
		} else if cur.val < val {
			cur = cur.right
		} else {
			//Value already exists so we could not insert the node
			return false
		}
	}

	cur = cur.parent
	n := RedBlackTree(val)
	n.color = Red
	if cur.val > val {
		cur.left = n
	} else {
		//Note that cur.Val != val because we would have already returned
		cur.right = n
	}

	return true
}

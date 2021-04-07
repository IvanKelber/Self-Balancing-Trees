package redblack

import (
	"fmt"
)

type Color bool

var Red Color = true
var Black Color = false

const (
	Left = iota
	Right
	RootNode
)

type RBNode struct {
	val    int
	isNil  bool
	color  Color
	parent *RBNode
	kids   []*RBNode
}

func nilRBNode() *RBNode {
	return &RBNode{-1, true, Black, nil, make([]*RBNode, 2)}
}

func RedBlackTree(val int) *RBNode {
	left := nilRBNode()
	right := nilRBNode()
	root := &RBNode{val, false, Black, nil, []*RBNode{left, right}}
	left.parent = root
	right.parent = root
	return root
}

func (rb *RBNode) String() string {
	return fmt.Sprintf("%v", rb.GetList())
}

func (rb *RBNode) left() *RBNode {
	return rb.kids[Left]
}

func (rb *RBNode) right() *RBNode {
	return rb.kids[Right]
}

func (rb *RBNode) GetList() []int {
	if rb.isNil {
		return []int{}
	}
	return append(append(rb.left().GetList(), rb.val), rb.right().GetList()...)
}

func (rb *RBNode) Contains(val int) bool {
	cur := rb
	for !cur.isNil {
		if cur.val > val {
			cur = cur.left()
		} else if cur.val < val {
			cur = cur.right()
		} else {
			return true
		}
	}
	return false
}

func (rb *RBNode) traverse(val int) *RBNode {
	cur := rb
	for !cur.isNil {
		if cur.val > val {
			cur = cur.left()
		} else if cur.val < val {
			cur = cur.right()
		} else {
			//Value already exists so we could not insert the node
			return cur
		}
	}
	return cur.getParent()
}

func (rb *RBNode) Insert(val int) bool {

	cur := rb.traverse(val)
	if cur.val == val {
		return false
	}
	n := RedBlackTree(val)
	n.color = Red
	n.parent = cur
	if cur.val > val {
		cur.kids[Left] = n
	} else {
		//Note that cur.Val != val because we would have already returned
		cur.kids[Right] = n
	}

	return true
}

func (rb *RBNode) getDirection() int {
	p := rb.getParent()
	if p == nil {
		return RootNode
	}
	if p.left() == rb {
		return Left
	}
	if p.right() == rb {
		return Right
	}
	fmt.Println("Error determining direction")
	return RootNode
}

func (rb *RBNode) getParent() *RBNode {
	return rb.parent
}

func (rb *RBNode) getGrandParent() *RBNode {
	p := rb.getParent()
	if p != nil {
		return p.getParent()
	}
	return nil
}

func (rb *RBNode) getSibling() *RBNode {
	direction := rb.getDirection()
	fmt.Printf("direction of %d is %d\n", rb.val, direction)
	if direction != RootNode {
		fmt.Printf("sibling of %d is %d\n", rb.val, rb.getParent().kids[(direction+1)%2].val)
		return rb.getParent().kids[(direction+1)%2]
	}
	return nil
}

func (rb *RBNode) getAunt() *RBNode {
	parent := rb.getParent()
	if parent != nil {
		fmt.Printf("Parent of %d is %d\n", rb.val, parent.val)

		return parent.getSibling()
	}
	return nil
}

func (rb *RBNode) getDistantNiece() *RBNode {
	sibling := rb.getSibling()
	if sibling != nil && !sibling.isNil {
		direction := rb.getDirection()
		return sibling.kids[(direction+1)%2]
	}
	return nil
}

func (rb *RBNode) getCloseNiece() *RBNode {
	sibling := rb.getSibling()
	if sibling != nil && !sibling.isNil {
		direction := rb.getDirection()
		return sibling.kids[direction]
	}
	return nil
}

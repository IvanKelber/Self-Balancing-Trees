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

func (rb *RBNode) Insert(val int) *RBNode {

	cur := rb.traverse(val)
	if cur.val == val {
		return nil
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

	return n
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
	if direction != RootNode {
		return rb.getParent().kids[(direction+1)%2]
	}
	return nil
}

func (rb *RBNode) getAunt() *RBNode {
	parent := rb.getParent()
	if parent != nil {
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

func (rb *RBNode) getRoot() *RBNode {
	cur := rb
	for cur.parent != nil {
		cur = cur.parent
	}
	return cur
}

//Make sure to check if the returned newSubtree has a parent
//If it doesn't then it is the new root of the entire tree
func (root *RBNode) rotate(subtree *RBNode, direction int) *RBNode {

	parent := subtree.getParent()
	subtreeDirection := subtree.getDirection()

	//get child of the subtree in the opposite direction that we are rotating
	newSubtree := subtree.kids[(direction+1)%2]
	if newSubtree.isNil {
		fmt.Println("Error while rotating: child is null")
		return root
	}
	newSubtree.kids[direction].parent = subtree
	newSubtree.parent = parent
	if parent != nil {
		parent.kids[subtreeDirection] = newSubtree
	}
	subtree.kids[(direction+1)%2] = newSubtree.kids[direction]
	newSubtree.kids[direction] = subtree
	subtree.parent = newSubtree

	return newSubtree
}

//In many cases node is the node that we just inserted into the tree
//However in some cases node refers to a recently updated node that may be causing a violation
func (root *RBNode) fixUp(node *RBNode) *RBNode {
	//case 0:
	if root == node {
		root.color = Black
		return root
	}
	parent := node.getParent() //must not be null otherwise we have case 0
	grandparent := node.getGrandParent()
	if grandparent == nil {
		//grandparent is nil so we have no aunts and therefore no violation
		return root
	}
	if node.color == Red && parent.color == Red {
		aunt := node.getAunt()
		//case 1:
		if aunt.color == Red {
			//push blackness from grandparent
			aunt.color = Black
			parent.color = Black
			grandparent.color = Red
			potentialRoot := root.fixUp(grandparent)
			if potentialRoot.parent == nil {
				return potentialRoot
			}
			return root
		}
		parentDirection := parent.getDirection()
		nodeDirection := node.getDirection()

		//case 2:
		if parentDirection == nodeDirection {
			// straight line
			// rotate in the opposite direction of the parent and node
			newSubtree := root.rotate(grandparent, (parentDirection+1)%2)
			newSubtree.color = Black
			grandparent.color = Red

			return newSubtree
		}
		//case 3:
		if parentDirection != nodeDirection {
			//rotate in the direction of the parent from the parent
			root.rotate(parent, parentDirection)
			//then rotate in the direction of the node from the grandparent
			potentialRoot := root.rotate(grandparent, nodeDirection)
			if potentialRoot.parent == nil {
				root = potentialRoot
			}
			root = root.fixUp(potentialRoot)
		}

	}
	return root
}

func (rb *RBNode) InsertAndBalance(val int) *RBNode {
	node := rb.Insert(val)
	if node == nil {
		return nil
	}
	return rb.fixUp(node)
}

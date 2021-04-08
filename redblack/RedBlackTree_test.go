package redblack

import (
	"math"
	"testing"
)

func TestCreation(t *testing.T) {
	type CreationTestCase struct {
		rootVal  int
		expected string
	}
	cases := []CreationTestCase{
		{0, "[0]"},
		{1, "[1]"},
		{-1, "[-1]"},
	}
	for _, testcase := range cases {
		rb := RedBlackTree(testcase.rootVal)
		if got := rb.String(); got != testcase.expected {
			t.Errorf("RedBlackTree(%d) Expected '%s' but got '%s' instead", testcase.rootVal, testcase.expected, got)
		}
	}
}

func TestInsertion(t *testing.T) {
	type InsertionTestCase struct {
		insertionOrder []int
		expected       string
	}
	cases := []InsertionTestCase{
		{[]int{0, 1, 2, 3, 4, 5}, "[0 1 2 3 4 5]"},
		{[]int{5, 4, 3, 2, 1, 0}, "[0 1 2 3 4 5]"},
		{[]int{3, 1, 5, 2, 0, 4}, "[0 1 2 3 4 5]"},
		{[]int{6, 6, 6, 6}, "[6]"},
		{[]int{6, 6, 6, 7}, "[6 7]"},
		{[]int{-1, -2, -3, -2, -3, -1}, "[-3 -2 -1]"},
	}
	for _, testcase := range cases {
		rb := RedBlackTree(testcase.insertionOrder[0])
		for i := 1; i < len(testcase.insertionOrder); i++ {
			rb.Insert(testcase.insertionOrder[i])
		}
		if got := rb.String(); got != testcase.expected {
			t.Errorf("Inserting(%v) Expected '%s' but got '%s' instead", testcase.insertionOrder, testcase.expected, got)
		}
	}
}

func TestContains(t *testing.T) {
	type ContainsTestCase struct {
		searchVal int
		expected  bool
	}
	treeVals := []int{0, 1, 2, 3, 4, 5, -1, -math.MaxInt32}
	rb := RedBlackTree(treeVals[0])
	for i := range treeVals {
		rb.Insert(treeVals[i])
	}

	cases := []ContainsTestCase{
		{0, true},
		{1, true},
		{2, true},
		{3, true},
		{4, true},
		{5, true},
		{-1, true},
		{6, false},
		{-2, false},
		{-8, false},
		{math.MaxInt32, false},
		{-math.MaxInt32, true},
	}
	for _, testcase := range cases {
		if got := rb.Contains(testcase.searchVal); got != testcase.expected {
			t.Errorf("Contains(%d) Expected '%v' but got '%v' instead", testcase.searchVal, testcase.expected, got)
		}
	}
}

func TestParent(t *testing.T) {
	type TestCase struct {
		nodeVal  int
		isNil    bool
		expected int
	}
	treeVals := []int{10, 5, 15, 8}
	rb := RedBlackTree(treeVals[0])
	for i := range treeVals {
		rb.Insert(treeVals[i])
	}

	cases := []TestCase{
		{10, true, -1},
		{5, false, 10},
		{15, false, 10},
		{8, false, 5},
	}
	for _, testcase := range cases {
		node := rb.traverse(testcase.nodeVal)
		parent := node.getParent()
		if parent == nil && !testcase.isNil {
			t.Errorf("GetParent(%d) Expected '%v' but got '%v' instead", testcase.nodeVal, testcase.expected, parent)
			continue
		}
		if parent == nil && testcase.isNil {
			continue
		}
		if parent != nil && testcase.isNil {
			t.Errorf("GetParent(%d) Expected '%v' but got '%v' instead", testcase.nodeVal, nil, parent.val)
			continue
		}

		if got := parent.val; got != testcase.expected {
			t.Errorf("GetParent(%d) Expected '%v' but got '%v' instead", testcase.nodeVal, testcase.expected, got)
		}
	}
}

func TestGrandparent(t *testing.T) {
	type TestCase struct {
		nodeVal  int
		isNil    bool
		expected int
	}
	treeVals := []int{10, 5, 15, 8, 3, 12, 20}
	rb := RedBlackTree(treeVals[0])
	for i := range treeVals {
		rb.Insert(treeVals[i])
	}

	cases := []TestCase{
		{10, true, -1},
		{5, true, -1},
		{15, true, -1},
		{3, false, 10},
		{8, false, 10},
		{12, false, 10},
		{20, false, 10},
	}
	for _, testcase := range cases {
		node := rb.traverse(testcase.nodeVal)
		grandParent := node.getGrandParent()
		if grandParent == nil && !testcase.isNil {
			t.Errorf("GetGrandParent(%d) Expected '%v' but got '%v' instead", testcase.nodeVal, testcase.expected, grandParent)
			continue
		}
		if grandParent == nil && testcase.isNil {
			continue
		}
		if grandParent != nil && testcase.isNil {
			t.Errorf("GetGrandParent(%d) Expected '%v' but got '%v' instead", testcase.nodeVal, nil, grandParent.val)
			continue
		}

		if got := grandParent.val; got != testcase.expected {
			t.Errorf("GetGrandParent(%d) Expected '%v' but got '%v' instead", testcase.nodeVal, testcase.expected, got)
		}
	}
}

func TestDirection(t *testing.T) {
	type TestCase struct {
		nodeVal  int
		expected int
	}
	treeVals := []int{10, 5, 15}
	rb := RedBlackTree(treeVals[0])
	for i := range treeVals {
		rb.Insert(treeVals[i])
	}

	cases := []TestCase{
		{10, RootNode},
		{5, Left},
		{15, Right},
	}
	for _, testcase := range cases {
		node := rb.traverse(testcase.nodeVal)

		if got := node.getDirection(); got != testcase.expected {
			t.Errorf("GetDirection(%d) Expected '%v' but got '%v' instead", testcase.nodeVal, testcase.expected, got)
		}
	}
}
func TestSibling(t *testing.T) {
	type TestCase struct {
		nodeVal  int
		isNil    bool
		expected int
	}
	treeVals := []int{10, 5, 15, 8, 3, 12}
	rb := RedBlackTree(treeVals[0])
	for i := range treeVals {
		rb.Insert(treeVals[i])
	}

	cases := []TestCase{
		{10, true, -1},
		{5, false, 15},
		{15, false, 5},
		{3, false, 8},
		{8, false, 3},
		{12, true, -1},
	}
	for _, testcase := range cases {
		node := rb.traverse(testcase.nodeVal)
		sibling := node.getSibling()
		if (sibling == nil || sibling.isNil) && !testcase.isNil {
			t.Errorf("GetSibling(%d) Expected '%v' but got '%v' instead", testcase.nodeVal, testcase.expected, sibling)
			continue
		}
		if (sibling == nil || sibling.isNil) && testcase.isNil {
			continue
		}
		if sibling != nil && !sibling.isNil && testcase.isNil {
			t.Errorf("GetSibling(%d) Expected '%v' but got '%v' instead", testcase.nodeVal, nil, sibling.val)
			continue
		}

		if got := sibling.val; got != testcase.expected {
			t.Errorf("GetSibling(%d) Expected '%v' but got '%v' instead", testcase.nodeVal, testcase.expected, got)
		}
	}
}

func TestAunt(t *testing.T) {
	type TestCase struct {
		nodeVal  int
		isNil    bool
		expected int
	}
	treeVals := []int{10, 5, 15, 8, 3, 12, 20}
	rb := RedBlackTree(treeVals[0])
	for i := range treeVals {
		rb.Insert(treeVals[i])
	}

	cases := []TestCase{
		{10, true, -1},
		{5, true, -1},
		{15, true, -1},
		{3, false, 15},
		{8, false, 15},
		{12, false, 5},
		{20, false, 5},
	}
	for _, testcase := range cases {
		node := rb.traverse(testcase.nodeVal)
		aunt := node.getAunt()
		if (aunt == nil || aunt.isNil) && !testcase.isNil {
			t.Errorf("GetAunt(%d) Expected '%v' but got '%v' instead", testcase.nodeVal, testcase.expected, aunt)
			continue
		}
		if (aunt == nil || aunt.isNil) && testcase.isNil {
			continue
		}
		if aunt != nil && !aunt.isNil && testcase.isNil {
			t.Errorf("GetAunt(%d) Expected '%v' but got '%v' instead", testcase.nodeVal, nil, aunt.val)
			continue
		}

		if got := aunt.val; got != testcase.expected {
			t.Errorf("GetAunt(%d) Expected '%v' but got '%v' instead", testcase.nodeVal, testcase.expected, got)
		}
	}
}

func TestDistantNiece(t *testing.T) {
	type TestCase struct {
		nodeVal  int
		isNil    bool
		expected int
	}
	treeVals := []int{10, 5, 15, 8, 3, 12, 20}
	rb := RedBlackTree(treeVals[0])
	for i := range treeVals {
		rb.Insert(treeVals[i])
	}

	cases := []TestCase{
		{10, true, -1},
		{5, false, 20},
		{15, false, 3},
		{3, true, -1},
		{8, true, -1},
		{12, true, -1},
		{20, true, -1},
	}
	for _, testcase := range cases {
		node := rb.traverse(testcase.nodeVal)
		niece := node.getDistantNiece()
		if (niece == nil || niece.isNil) && !testcase.isNil {
			t.Errorf("GetDistantNiece(%d) Expected '%v' but got '%v' instead", testcase.nodeVal, testcase.expected, niece)
			continue
		}
		if (niece == nil || niece.isNil) && testcase.isNil {
			continue
		}
		if niece != nil && !niece.isNil && testcase.isNil {
			t.Errorf("GetDistantNiece(%d) Expected '%v' but got '%v' instead", testcase.nodeVal, nil, niece.val)
			continue
		}

		if got := niece.val; got != testcase.expected {
			t.Errorf("GetDistantNiece(%d) Expected '%v' but got '%v' instead", testcase.nodeVal, testcase.expected, got)
		}
	}
}

func TestCloseNiece(t *testing.T) {
	type TestCase struct {
		nodeVal  int
		isNil    bool
		expected int
	}
	treeVals := []int{10, 5, 15, 8, 3, 12, 20}
	rb := RedBlackTree(treeVals[0])
	for i := range treeVals {
		rb.Insert(treeVals[i])
	}

	cases := []TestCase{
		{10, true, -1},
		{5, false, 12},
		{15, false, 8},
		{3, true, -1},
		{8, true, -1},
		{12, true, -1},
		{20, true, -1},
	}
	for _, testcase := range cases {
		node := rb.traverse(testcase.nodeVal)
		niece := node.getCloseNiece()
		if (niece == nil || niece.isNil) && !testcase.isNil {
			t.Errorf("GetCloseNiece(%d) Expected '%v' but got '%v' instead", testcase.nodeVal, testcase.expected, niece)
			continue
		}
		if (niece == nil || niece.isNil) && testcase.isNil {
			continue
		}
		if niece != nil && !niece.isNil && testcase.isNil {
			t.Errorf("GetCloseNiece(%d) Expected '%v' but got '%v' instead", testcase.nodeVal, nil, niece.val)
			continue
		}

		if got := niece.val; got != testcase.expected {
			t.Errorf("GetCloseNiece(%d) Expected '%v' but got '%v' instead", testcase.nodeVal, testcase.expected, got)
		}
	}
}

func TestRotate(t *testing.T) {
	type TestCase struct {
		treeVals           []int
		direction          int
		subtreeVal         int
		expectedNewSubtree int
	}

	cases := []TestCase{
		{[]int{100, 120, 150}, Left, 100, 120},
		{[]int{100, 50, 30}, Right, 100, 50},
		{[]int{100, 50, 120, 80, 90, 95}, Left, 80, 90},
		{[]int{100, 50, 30, 120, 80, 70, 90, 95}, Left, 50, 80},
		{[]int{100, 80, 50, 30, 120, 70, 90, 95}, Right, 100, 80},
	}
	for _, testcase := range cases {

		rb := RedBlackTree(testcase.treeVals[0])
		for i := range testcase.treeVals {
			rb.Insert(testcase.treeVals[i])
		}
		originalString := rb.String()
		subtree := rb.traverse(testcase.subtreeVal)
		newSubtree := rb.rotate(subtree, testcase.direction)
		if newSubtree == nil {
			t.Errorf("Rotate(%d, %d) resulted in nil new subtree", testcase.subtreeVal, testcase.direction)
		}
		if got := newSubtree.val; got != testcase.expectedNewSubtree {
			t.Errorf("Rotate(%d, %d) Expected '%d' but got '%d' instead", testcase.subtreeVal, testcase.direction,
				testcase.expectedNewSubtree, got)
		}
		if newSubtree.parent == nil {
			rb = newSubtree
		}
		newString := rb.String()
		if newString != originalString {
			t.Errorf("Rotate(%d, %d) Expected '%s' but got '%s' instead", testcase.subtreeVal, testcase.direction,
				originalString, newString)
		}
	}
}

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

func TestSearch(t *testing.T) {
	type InsertionTestCase struct {
		searchVal int
		expected  bool
	}
	treeVals := []int{0, 1, 2, 3, 4, 5, -1, -math.MaxInt32}
	rb := RedBlackTree(treeVals[0])
	for i := range treeVals {
		rb.Insert(treeVals[i])
	}

	cases := []InsertionTestCase{
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

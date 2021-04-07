package redblack

import "testing"

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

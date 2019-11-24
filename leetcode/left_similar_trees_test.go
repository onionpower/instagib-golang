package leetcode

import (
	"testing"
)

func TestSimilar(t *testing.T) {
	if leafSimilar(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
		&TreeNode{Val: 3, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 2}}) {
		t.Error("expected to be false")
	}
}
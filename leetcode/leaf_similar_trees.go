package leetcode

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	seq1 := make([]int, 0)
	traverse(root1, &seq1)
	seq2 := make([]int, 0)
	traverse(root2, &seq2)

	if len(seq1) != len(seq2) {
		return false
	}

	for i := 0; i < len(seq1); i++ {
		if seq1[i] != seq2[i] {
			return false
		}
	}

	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traverse(n *TreeNode, s *[]int) {
	if n == nil {
		return
	}
	if n.Left == nil && n.Right == nil {
		*s = append(*s, n.Val)
		return
	}
	traverse(n.Left, s)
	traverse(n.Right, s)
}

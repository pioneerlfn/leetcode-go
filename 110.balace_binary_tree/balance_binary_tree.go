package _10_balace_binary_tree

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil || root.Left == nil && root.Right == nil {
		return true
	}
	leftDepth := treeDepth(root.Left)
	rightDepth := treeDepth(root.Right)
	if leftDepth - rightDepth > 1 || rightDepth - leftDepth > 1 || !isBalanced(root.Left) || isBalanced(root.Right) {
		return false
	}
	return true
}

func treeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := treeDepth(root.Left)
	rightDepth := treeDepth(root.Right)
	depth := max(leftDepth, rightDepth) + 1
	return depth
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

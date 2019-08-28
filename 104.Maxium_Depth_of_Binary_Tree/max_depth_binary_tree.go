package _04_Maxium_Depth_of_Binary_Tree

import (
	// "container/heap"
	// "math"
)

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

// Depth first Search.
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	depth := max(leftDepth, rightDepth) + 1
	return depth
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package _8_valid_bst

import "math"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if isValidBST(root.Left) && isValidBST(root.Right) {
		leftMax := int(math.MinInt64)
		rightMin := int(math.MaxInt64)
		if root.Left != nil {
			leftMax = maxVal(root.Left)
		}
		if root.Right != nil {
			rightMin = minVal(root.Right)
		}
		if leftMax < root.Val && root.Val < rightMin {
			return true
		}
	}
	return false
}

// root != nil
func maxVal(root *TreeNode) int {
	cur := root
	for cur.Right != nil {
		cur = cur.Right
	}
	return cur.Val
}

// root != nil
func minVal(root *TreeNode) int {
	cur := root
	for cur.Left != nil {
		cur = cur.Left
	}
	return cur.Val
}

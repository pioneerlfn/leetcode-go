package _45_binary_tree_post_traversal

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var results []int
	post(root, &results)
	return results
}

func post(root *TreeNode, results *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		post(root.Left, results)
	}
	if root.Right != nil {
		post(root.Right, results)
	}
	*results = append(*results, root.Val)
	return
}


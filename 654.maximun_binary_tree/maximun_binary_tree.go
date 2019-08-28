package _54_maximun_binary_tree

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{nums[0], nil, nil}
	}
	// nums has at least two elements.
	// maxIndex can be the first or the last of the nums.
	maxIndex := findMax(nums)
	root := &TreeNode{nums[maxIndex], nil, nil}
	var leftTree, rightTree *TreeNode
	if maxIndex > 0 {
		leftTree = constructMaximumBinaryTree(nums[:maxIndex])
	}
	if maxIndex < len(nums)-1 {
		rightTree = constructMaximumBinaryTree(nums[maxIndex+1:])
	}
	root.Left = leftTree
	root.Right = rightTree
	return root
}

func findMax(nums []int) int {
	var index int
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[index] {
			index = i
		}
	}
	return index
}


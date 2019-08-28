package _06_binary_tree_from_inorder_and_postorder

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(inorder) != len(postorder) {
		return nil
	}
	root := &TreeNode{postorder[len(postorder)-1], nil, nil}
	if len(postorder) == 1 {
		return root
	}
	var i int
	for i = 0; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			break
		}
	}

	inorderLeft := inorder[:i] // [9]
	inorderRight := inorder[i+1:] // [15, 20, 7]
	postorderLeft := postorder[:len(inorderLeft)] // [9]
	postorderRight := postorder[len(inorderLeft):len(inorderLeft)+len(inorderRight)]
	leftTree := buildTree(inorderLeft, postorderLeft)
	rightTree := buildTree(inorderRight, postorderRight)
	root.Left = leftTree
	root.Right = rightTree
	return root
}

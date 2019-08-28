package _05_binary_tree_from_preorder_and_inorder

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	if len(preorder) == 1 {
		return root
	}
	var i int
	for i = 0; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			break
		}
	}

	inorderLeft := inorder[:i]
	inorderRight := inorder[i+1:]
	preorderLeft := preorder[1:len(inorderLeft)+1]
	preorderRight := preorder[len(inorderLeft)+1:]
	leftTree := buildTree(preorderLeft, inorderLeft)
	rightTree := buildTree(preorderRight, inorderRight)
	root.Left = leftTree
	root.Right = rightTree
	return root
}

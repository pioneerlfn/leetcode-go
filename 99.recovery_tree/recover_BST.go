package _9_recovery_tree

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func recoverTree(root *TreeNode)  {
	if root == nil {
		return
	}
	var pre, first, second *TreeNode
	var traverse func(*TreeNode)
	// closure
	traverse = func(root *TreeNode) {
		if root.Left != nil {
			traverse(root.Left)
		}
		if pre != nil && pre.Val > root.Val {
			if first == nil {
				first = pre
			}
			if first != nil {
				second = root
			}
		}
		pre = root

		if root.Right != nil {
			traverse(root.Right)
		}
	}
	traverse(root)
	first.Val, second.Val = second.Val, first.Val
	return
}


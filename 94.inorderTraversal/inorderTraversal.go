package _4_inorderTraversal


type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

/*
 * treat right branch as part of the root
 * so we only prune left branch and push the remain part to the stack.
 */
func inorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	for	cur := root; cur != nil; {
		if cur.Left != nil {
			left := cur.Left
			cur.Left = nil
			// cur is the remainder with left branch pruned.
			stack = append(stack, cur)
			cur = left
		} else {
			// cur.Left == nil
			 res = append(res, cur.Val)
			 if cur.Right != nil {
			 	cur = cur.Right
			 } else {
			 	// cur.Left == nil && cur.Right == nil
			 	if len(stack) == 0 {
			 		break
				}
			 	cur = stack[len(stack)-1]
			 	stack = stack[:len(stack)-1]
			 }
		}
	}
	return res
}

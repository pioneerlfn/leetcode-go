package _45_binary_tree_post_traversal

func postIterationTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	var res []int
	cur := root
	for cur != nil {
		// leaf node
		if cur.Left == nil && cur.Right == nil {
			// only a node become a 'leaf node', we append it to result.
			res = append(res, cur.Val)
			if len(stack) == 0 {
				break
			}
			// pop from stack.
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		} else {
			left, right := cur.Left, cur.Right
			// prune, avoid repeatedly visit.
			cur.Left, cur.Right = nil, nil
			// push root to stack
			stack = append(stack, cur)
			if left != nil {
				cur = left
				if right != nil {
					// push right to stack.
					stack = append(stack, right)
				}
			} else {
				cur = right
			}
		}
	}
	return res
}

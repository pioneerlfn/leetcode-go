package _44__binaryTree_preorderTraversal

import "github.com/pingcap/tidb/ast"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func preorderTravesal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	cur := root
	for cur != nil {
		res = append(res, cur.Val)
		if cur.Left != nil {
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
			cur = cur.Left
		} else {
			// cur.Left == nil
			if cur.Right != nil {
				cur = cur.Right
			} else {
				// cur.Left == nil && cur.Right == nil
				// stack is empty.
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

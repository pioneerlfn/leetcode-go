package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// BFS, level traversal.
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// root has at least one node.
	var queue []*TreeNode
	var depth int
	queue = append(queue, root)
	for len(queue) != 0 {
		depth++
		// 注意先用一个变量记下queue的大小.
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:] // if queue has only one element, after this slice, queue become nil.
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return depth
}

func main() {
	n1 := &TreeNode{3, nil, nil}
	n2 := &TreeNode{9, nil, nil}
	n3 := &TreeNode{20, nil, nil}
	n4 := &TreeNode{15, nil, nil}
	n5 := &TreeNode{7, nil, nil}
	n1.Left = n2
	n1.Right = n3
	n3.Left = n4
	n3.Right = n5
	var queue []*TreeNode
	queue = append(queue, n1)
	fmt.Println(queue[0])
	fmt.Println(queue[1:])
}
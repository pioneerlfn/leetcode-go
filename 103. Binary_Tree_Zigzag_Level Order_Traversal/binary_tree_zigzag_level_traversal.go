package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var queue []*TreeNode
	queue = append(queue, root)
	var level int
	var levelVals []int
	for len(queue) != 0 {
		level++
		size := len(queue)
		levelVals = []int{}
		for i := 0; i < size; i++ {
			var node *TreeNode
			node = queue[0]
			queue = queue[1:]
			if level % 2 == 1 {
				levelVals = append(levelVals, node.Val)
			} else {
				temp := make([]int, len(levelVals)+1)
				temp[0] = node.Val
				copy(temp[1:], levelVals)
				levelVals = temp
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, levelVals)
	}
	return res
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
	res := zigzagLevelOrder(n1)
	fmt.Println(res)
}
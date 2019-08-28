package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// initialize.
	sorted := head
	head = head.Next
	sorted.Next = nil

	// pop a node from the unsorted list
	// and insert it to the sorted list.
	for head != nil {
		node := head
		head = head.Next
		node.Next = nil
		sorted = insert(sorted, node)
	}
	return sorted
}

func insert(sorted, node *ListNode) *ListNode {
	cur := sorted
	// insert to the start.
	if node.Val < sorted.Val {
		node.Next = cur
		sorted = node
		return sorted
	}

	var pre *ListNode
	for cur != nil {
		if node.Val >= cur.Val {
			pre, cur = cur, cur.Next
		} else {
			node.Next = cur
			pre.Next = node
			break
		}
	}

	// insert to the end.
	if cur == nil {
		pre.Next = node
		node.Next = nil
	}
	return sorted
}

func PrintList(list *ListNode)  {
	fmt.Println("*************")
	for cur := list; cur != nil; cur = cur.Next {
		fmt.Println(cur.Val)
	}
}

func main() {
	n1 := &ListNode{4, nil}
	n2 := &ListNode{2, nil}
	n3 := &ListNode{1, nil}
	n4 := &ListNode{3, nil}
	n3.Next = n4
	n2.Next = n3
	n1.Next = n2
	head := n1

	insertionSortList(head)
}
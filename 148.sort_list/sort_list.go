package main

import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	sorted := MergeSort(head)
	return sorted
}

func MergeSort(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	left, right := split(head)
	return Merge(MergeSort(left), MergeSort(right))
}

func Merge(l1 *ListNode, l2 *ListNode) *ListNode {
	// fmt.Println("\nComing into merge()...")
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head, node *ListNode
	if l1.Val < l2.Val {
		head = l1
		node = l1
		l1 = l1.Next
	} else {
		head = l2
		node = l2
		l2 = l2.Next
	}

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
		node = node.Next
	}

	if l1 != nil {
		node.Next = l1
	}

	if l2 != nil {
		node.Next = l2
	}
	// fmt.Println("Leaving merge()...")
	return head
}

func Merge_(first, second *ListNode) *ListNode {
	fmt.Println("\ncomming to Merge()...")
	var head *ListNode
	var follower *ListNode
	if first.Val > second.Val {
		head = second
		follower = first
		fmt.Println("head value: ", head.Val)
	} else {
		head = first
		follower = second
	}

	cursor11, cursor12 := head, head

	for cursor11 != nil && follower != nil {
		if cursor11.Val <= follower.Val {
			cursor11 = cursor11.Next
			cursor12 = cursor12.Next
		} else {
			cursor2 := follower
			follower = follower.Next
			cursor2.Next = cursor11
			cursor12.Next = cursor2
		}
	}

	if follower != nil {
		cursor12 = follower
		follower = nil
	}

	cursor11, cursor12 = nil, nil
	fmt.Println("leaving merge()...")
	return head
}

func split(list *ListNode)  (*ListNode, *ListNode) {
	fast, slow, preSlow := list, list, list
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow, preSlow = slow.Next, slow
	}
	preSlow.Next = nil
	return list, slow
}

func MiddleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
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

	sortList(head)
}

func PrintList(list *ListNode)  {
	fmt.Println("***************Print merged list**********")
	for cur := list; cur != nil; cur = cur.Next {
		fmt.Println(cur.Val)
	}
}
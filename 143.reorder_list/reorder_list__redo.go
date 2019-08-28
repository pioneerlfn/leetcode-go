package main

import "fmt"

// Given 1->2->3->4, reorder it to 1->4->2->3.
// Given 1->2->3->4->5, reorder it to 1->5->2->4->3.

type ListNode struct {
	Val int
	Next *ListNode
}

func reorderListII(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	// head has at least two nodes.
	// first equal to head.
	fisrt, second := splitII(head)
	// second has at least one node.
	secondRersed := reverseII(second)
	// fisrt should not be changed.
	mergeII(fisrt, secondRersed)
	return
}

func mergeII(first, second *ListNode)  {
	if second == nil {
		return
	}
	cur1, cur2 := first, second
	for cur1 != nil && cur2 != nil {
		second = second.Next
		cur2.Next = cur1.Next
		cur1.Next = cur2
		cur1 = cur2.Next
		cur2 = second
	}
}

func reverseII(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	posterior := reverseII(head.Next)
	head.Next.Next = head
	head.Next = nil

	return posterior
}

func splitII(head *ListNode) (*ListNode, *ListNode) {
	dummy := &ListNode{}
	dummy.Next = head
	tail, slow, fast := dummy, head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		tail, slow = tail.Next, slow.Next
	}
	tail.Next = nil

	return head, slow
}

func main() {
	n1 := &ListNode{1, nil}
	n2 := &ListNode{2, nil}
	n3 := &ListNode{3, nil}
	n4 := &ListNode{4, nil}
	n5 := &ListNode{5, nil}
	n6 := &ListNode{6, nil}
	n5.Next = n6
	n4.Next = n5
	n3.Next = n4
	n2.Next = n3
	n1.Next = n2
	head := n1
	PrintList(head, "==BEFORE==")
	first, second := splitII(head)
	PrintList(first, "FISRT")
	PrintList(second, "SECOND")
	reversed := reverseII(second)
	PrintList(reversed, "REVERSED")
	// reorderListII(head)
	mergeII(first, reversed)
	PrintList(first, "==AFTER==")
}

func PrintList(list *ListNode, name string)  {
	fmt.Printf("print list %s\n", name)
	for cur := list; cur != nil; cur = cur.Next {
		fmt.Print("-->", cur.Val)
	}
	fmt.Println()
}
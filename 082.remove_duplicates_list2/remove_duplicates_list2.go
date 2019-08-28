package main

import (
	"fmt"
	"math"
)

type ListNode struct {
    Val int
    Next *ListNode
}

// Input: 1->2->3->3->4->4->5
// Output: 1->2->5
// Input: 1->1->1->2->3
// Output: 2->3

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{math.MinInt64, nil}
	dummy.Next = head

	p := dummy
	for p.Next != nil && p.Next.Next != nil {
		if p.Next.Val == p.Next.Next.Val {
			dup := p.Next.Val
			for p.Next != nil && p.Next.Val == dup {
				p = p.Next
			}
		} else {
			p = p.Next
		}

	return dummy.Next
}








}


func main() {
	n1 := &ListNode{1, nil}
	n2 := &ListNode{1, nil}
	//n3 := &ListNode{3, nil}
	//n4 := &ListNode{3, nil}
	//n5 := &ListNode{4, nil}
	//n6 := &ListNode{4, nil}
	//n7 := &ListNode{5, nil}

	//n6.Next = n7
	//
	// n5.Next = n6
	//n4.Next = n5
	//n3.Next = n4
	//n2.Next = n3
	n1.Next = n2
	head := n1
	PrintList(head, "list before deduplicate...")
	newHead := deleteDuplicates(head)
	PrintList(newHead, "List At last...")
}

func PrintList(list *ListNode, name string)  {
	fmt.Printf("print list %s\n", name)
	for cur := list; cur != nil; cur = cur.Next {
		fmt.Print("-->", cur.Val)
	}
	fmt.Println()
}
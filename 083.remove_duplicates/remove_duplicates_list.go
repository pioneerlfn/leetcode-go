package main

import (
	"fmt"
	"os"
)

type ListNode struct {
    Val int
    Next *ListNode
}


func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fast, slow := head.Next, head
	for fast != nil {
		if fast.Val == slow.Val {
			slow.Next = fast.Next
			fast = fast.Next
		} else {
			slow = slow.Next
			fast = fast.Next
		}
	}

	return head
}

func main() {
	n1 := &ListNode{1, nil}
	n2 := &ListNode{2, nil}
	n3 := &ListNode{2, nil}
	n4 := &ListNode{3, nil}
	n5 := &ListNode{3, nil}
	// n6 := &ListNode{6, nil}
	// n5.Next = n6
	n4.Next = n5
	n3.Next = n4
	n2.Next = n3
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
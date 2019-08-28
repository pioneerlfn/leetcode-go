package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

// Input: head = 1->4->3->2->5->2, x = 3
// Output: 1->2->2->4->3->5
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy, dummy2 := &ListNode{}, &ListNode{}
	dummy.Next = head
	p, cur := dummy, dummy2

	// find the first item not small than x.
	for p = dummy; p.Next != nil; {
		if p.Next.Val < x {
			p = p.Next
		} else {
			cur.Next = p.Next
			cur = cur.Next
			p.Next = p.Next.Next
		}
	}
	cur.Next = nil
	p.Next = dummy2.Next

	return dummy.Next
}

func main() {
	n1 := &ListNode{1, nil}
	n2 := &ListNode{4, nil}
	n3 := &ListNode{3, nil}
	n4 := &ListNode{2, nil}
	n5 := &ListNode{5, nil}
	n6 := &ListNode{2, nil}
	//n7 := &ListNode{5, nil}

	//n6.Next = n7
	//
	 n5.Next = n6
	n4.Next = n5
	n3.Next = n4
	n2.Next = n3
	n1.Next = n2
	head := n1
	PrintList(head, "list before")
	newHead := partition(head, 3)
	PrintList(newHead, "List At last...")
}

func PrintList(list *ListNode, name string)  {
	fmt.Printf("print list %s\n", name)
	for cur := list; cur != nil; cur = cur.Next {
		fmt.Print("-->", cur.Val)
	}
	fmt.Println()
}
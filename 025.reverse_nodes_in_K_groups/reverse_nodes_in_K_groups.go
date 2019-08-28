package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 1 {
		return head
	}

	cur := head
	var i int
	for i = 0; cur.Next != nil; i++	{
		if i == k-1 {
			break
		}
		cur = cur.Next
	}
	nextPart := cur.Next
	cur.Next = nil
	var newHead, newTail *ListNode
	if i == k-1 {
		newHead, newTail = reverse(head, cur)
	} else {
		newHead, newTail = head, cur
	}

	newTail.Next = reverseKGroup(nextPart, k)

	return newHead
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	if head == nil || head.Next == nil {
		return head, head
	}

	newHead, _ := reverse(head.Next, tail)
	head.Next.Next = head
	head.Next = nil
	// head now is new tail
	return newHead, head
}

func main() {
	n1 := &ListNode{1, nil}
	n2 := &ListNode{2, nil}
	n3 := &ListNode{3, nil}
	n4 := &ListNode{4, nil}
	n5 := &ListNode{5, nil}
	// n6 := &ListNode{6, nil}
	// n5.Next = n6
	n4.Next = n5
	n3.Next = n4
	n2.Next = n3
	n1.Next = n2
	head := n1
	// PrintList(head, "list before reverse...")
	newHead := reverseKGroup(head, 3)
	PrintList(newHead, "List At last...")
}

func PrintList(list *ListNode, name string)  {
	fmt.Printf("print list %s\n", name)
	for cur := list; cur != nil; cur = cur.Next {
		fmt.Print("-->", cur.Val)
	}
	fmt.Println()
}
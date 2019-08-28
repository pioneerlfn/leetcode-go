package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	fast := head
	length := 1
	for fast.Next != nil {
		length++
		fast = fast.Next
	}

	fmt.Println("length: ", length)

	if k %  length == 0 {
		return head
	}

	fast = head
	for i := 0; i < k % length; i++ {
		fast = fast.Next
	}
	fmt.Println("fast value: ", fast.Val)

	slow := head
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	fmt.Println("fast value: ", fast.Val)
	fmt.Println("fast value: ", slow.Val)


	newHead := slow.Next
	slow.Next = nil
	fast.Next = head

	return newHead
}

func main() {
	n1 := &ListNode{1, nil}
	n2 := &ListNode{2, nil}

	n1.Next = n2
	head := n1

	PrintList(head, "list before")
	newHead := rotateRight(head, 1)
	PrintList(newHead, "list After rotate.")

}

func PrintList(list *ListNode, name string)  {
	fmt.Printf("print list %s\n", name)
	for cur := list; cur != nil; cur = cur.Next {
		fmt.Print("-->", cur.Val)
	}
	fmt.Println()
}
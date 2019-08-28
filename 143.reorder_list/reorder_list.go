package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	// head has at least two element.
	first, second := split(head)
	second = reverse(second)
	head = merge(first, second)
	return
}

/** It's important to have a return value when we manipulate linked list.
 * this reverse is based on stack and it's tediously long
 * maybe we can find a more concise version.
 */

func reverse(l *ListNode) *ListNode {
	if l == nil || l.Next == nil {
		return l
	}
	var stack []*ListNode
	for cur := l; cur != nil; {
		before := cur
		cur = cur.Next
		before.Next = nil
		stack = append(stack, before)
	}

	l = stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	var tmp *ListNode
	for cur := l; len(stack) != 0; {
		tmp = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur.Next = tmp
		cur = cur.Next
	}
	return l
}

/**
 * reverse linked list based on recursive.
 */
func reverseRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}


func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	// assume l1 is not nil.
	if l2 == nil {
		return l1
	}
	// l2 is not nil too.
	node1, node2 := l1, l2
	for node1 != nil && node2 != nil {
		l2 = l2.Next
		node2.Next = node1.Next
		node1.Next = node2
		node1 = node2.Next
		node2 = l2
	}
	return l1
}

func split(head *ListNode) (*ListNode, *ListNode) {
	// head has at least two elements.
	fast, slow := head.Next, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	before := slow
	slow = slow.Next
	before.Next = nil
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
	first, second := split(head)
	second = reverse(second)
	merge(first, second)
	PrintList(first, "after merge...")
}

func PrintList(list *ListNode, name string)  {
	fmt.Printf("print list %s\n", name)
	for cur := list; cur != nil; cur = cur.Next {
		fmt.Print("-->", cur.Val)
	}
	fmt.Println()
}
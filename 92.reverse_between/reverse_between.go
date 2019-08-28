package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

// 1 <= m <= n <= length
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil || m == n {
		return head
	}
	cur := head
	for i := 1; i < m; i++ {
		cur = cur.Next
	}
	length := n - m + 1
	middle := make([]*ListNode, length)
	for i, j := 0, m; j <= n; j++ {
		middle[i] = cur
		cur = cur.Next
		i++
	}

	fmt.Println(middle)
	for i := 0; i < length; i++ {
		fmt.Printf("middle[%d]: %d", i, middle[i].Val)
	}
	for i := 0; i < (length)/2; i++ {
		middle[i].Val, middle[length-1-i].Val = middle[length-1-i].Val, middle[i].Val
	}
	return head
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
	reverseBetween(head, 2, 4)
}

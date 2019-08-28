package _06_reverse_linked_list

type ListNode struct {
    Val int
    Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	reversed := reverseList(head.Next) // keypoint.
	head.Next.Next = head
	head.Next = nil
	return reversed
}

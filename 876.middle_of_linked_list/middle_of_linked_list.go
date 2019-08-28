package _76_middle_of_linked_list

type ListNode struct {
     Val int
     Next *ListNode
}

func MiddleNode(head *ListNode) *ListNode {
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow



}

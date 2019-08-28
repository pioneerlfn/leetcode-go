package _21_merge_two_sorted_list

type ListNode struct {
    Val int
    Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	// l1 not nil.
	if l2 == nil {
		return l1
	}
	// neither l1 or l2 is nil.
	var head, node *ListNode
	if l1.Val < l2.Val {
		head = l1
		l1 = l1.Next
		node = head
	} else {
		head = l2
		l2 = l2.Next
		node = head
	}

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
		node = node.Next
	}

	if l1 != nil {
		node.Next = l1
	}
	if l2 != nil {
		node.Next = l2
	}

	return head
}
package _41_linked_list_cycle

func detectCycle(head *ListNode) *ListNode {
	// no cycle.
	if head == nil || head.Next == nil {
		return nil
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			break
		}
	}
	// no cycle.
	if fast == nil || fast.Next == nil {
		return nil
	}
	// has cycle
	// cycle start from position 0, handle it specially.
	if slow == head {
		return head
	}
	// two points, one start from the head
	// the other start from the meeting node.
	// start going ahead at the same time, each with one step a time.
	// when they meet, we find the node where cycle begins.
	ahead, behind := slow, head
	for {
		ahead = ahead.Next
		behind = behind.Next
		if ahead == behind {
			break
		}
	}
	return ahead
}

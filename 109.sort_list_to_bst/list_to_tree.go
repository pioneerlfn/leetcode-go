package _09_sort_list_to_bst

type ListNode struct {
    Val int
    Next *ListNode
}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	// has only one element, convert it to type of TreeNode.
	if head.Next == nil {
		return &TreeNode{head.Val, nil, nil}
	}
	// has at least two elements.
	// find the mid node, split it to two part.
	left, mid, right := split(head)
	midNode := &TreeNode{mid.Val, nil, nil}
	leftTree := sortedListToBST(left)
	rightTree := sortedListToBST(right)
	midNode.Left = leftTree
	midNode.Right = rightTree
	return midNode
}

// assume list has at least two elements.
func split(head *ListNode) (*ListNode, *ListNode, *ListNode) {
	fast, slow, follower := head, head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow, follower = slow.Next, slow
	}
	follower.Next = nil

	// slow is mid node.
	// only two element, mid = slow, right = nil.
	if slow.Next == nil {
		return head, slow, nil
	}
	// more that two element, mid and right can be different.
	second := slow.Next
	slow.Next = nil
	return head, slow, second
}
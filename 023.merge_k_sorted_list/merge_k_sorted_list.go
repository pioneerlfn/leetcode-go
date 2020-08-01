package _23_merge_k_sorted_list

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {

	if len(lists) == 0 {
		return nil
	}
	return merge(lists)
}

func merge(lists []*ListNode) *ListNode {
	length := len(lists)

	if length == 1 {
		return lists[0]
	}

	if length == 2 {
		var (
			l0, l1   = lists[0], lists[1]
			res, cur *ListNode
		)

		if l0 == nil {
			return l1
		}
		if l1 == nil {
			return l0
		}
		// both l0 and l1 != nil.
		if l0.Val < l1.Val {
			res, cur, l0 = l0, l0, l0.Next
		} else {
			res, cur, l1 = l1, l1, l1.Next
		}

		for l0 != nil && l1 != nil {
			if l0.Val < l1.Val {
				cur.Next, l0 = l0, l0.Next
			} else {
				cur.Next, l1 = l1, l1.Next
			}
			cur = cur.Next
		}

		if l0 != nil {
			cur.Next = l0
		}
		if l1 != nil {
			cur.Next = l1
		}

		return res
	}

	half := length / 2
	return mergeKLists([]*ListNode{mergeKLists(lists[half:]), mergeKLists(lists[:half])})
}

// 2020-08-01
// 解法1:递归
func mergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) == 2 {
		return merge2Lists(lists[0], lists[1])
	}
	k := len(lists)
	return merge2Lists(mergeKLists2(lists[:k/2]), mergeKLists2(lists[k/2:]))
}

func merge2Lists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var m *ListNode
	if l1.Val < l2.Val {
		m = l1
		l1 = l1.Next
	} else {
		m = l2
		l2 = l2.Next
	}
	p := m
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}
	return m
}

// 解法2:最小堆
type PriorityQueue []int

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i] < pq[j]
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	it := x.(int)
	*pq = append(*pq, it)
}
func (pq *PriorityQueue) Pop() interface{} {
	item := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return item
}

func mergeKLists3(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	var pq PriorityQueue
	for _, l := range lists {
		if l != nil {
			for l != nil {
				pq = append(pq, l.Val)
				l = l.Next
			}
		}
	}
	heap.Init(&pq)
	res := new(ListNode)
	cur := res
	for len(pq) != 0 {
		cur.Next = &ListNode{heap.Pop(&pq).(int), nil}
		cur = cur.Next
	}
	return res.Next
}

package swapnodesinpairs

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}
	pre := dummy
	cur := head

	for cur != nil && cur.Next != nil {
		temp1 := cur.Next
		temp2 := cur.Next.Next

		pre.Next = temp1
		temp1.Next = cur
		cur.Next = temp2

		pre = cur
		cur = cur.Next
	}

	return dummy.Next
}

func swapPairs2(head *ListNode) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}
	cur := dummy

	for cur.Next != nil && cur.Next.Next != nil {
		temp1 := cur.Next
		temp2 := cur.Next.Next.Next

		cur.Next = cur.Next.Next
		cur.Next.Next = temp1
		temp1.Next = temp2

		cur = cur.Next.Next
	}

	return dummy.Next
}

package intersectionoftwolinkedlistslcci

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB := 0, 0
	curA, curB := headA, headB
	for curA != nil {
		curA = curA.Next
		lenA++
	}
	for curB != nil {
		curB = curB.Next
		lenB++
	}

	var fast, slow *ListNode
	step := 0
	if lenA > lenB {
		fast, slow = headA, headB
		step = lenA - lenB
	} else {
		fast, slow = headB, headA
		step = lenB - lenA
	}
	for i := 0; i < step; i++ {
		fast = fast.Next
	}

	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

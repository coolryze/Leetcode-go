package intersectionoftwolinkedlistslcci

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	listA *ListNode
	listB *ListNode
}

type output struct {
	expected *ListNode
}

func linkedListToArray(head *ListNode) []int {
	arr := make([]int, 0)
	cur := head
	for cur != nil {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}
	return arr
}

func arrayToLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := &ListNode{Val: arr[0]}
	current := head
	for i := 1; i < len(arr); i++ {
		current.Next = &ListNode{Val: arr[i]}
		current = current.Next
	}
	return head
}

func TestGetIntersectionNode(t *testing.T) {
	tc1InterNode := arrayToLinkedList([]int{8, 4, 5})
	tc1NodeA := arrayToLinkedList([]int{4, 1})
	tc1NodeA.Next = tc1InterNode
	tc1NodeB := arrayToLinkedList([]int{5, 0, 1})
	tc1NodeB.Next = tc1InterNode
	tcs := []testcase{
		testcase{
			input{tc1NodeA, tc1NodeB},
			output{tc1InterNode},
		},
		testcase{
			input{arrayToLinkedList([]int{2, 6, 4}), arrayToLinkedList([]int{1, 5})},
			output{nil},
		},
	}

	for _, tc := range tcs {
		got := getIntersectionNode(tc.input.listA, tc.input.listB)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}

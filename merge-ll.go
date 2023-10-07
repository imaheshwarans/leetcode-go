package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
// 	if l1 == nil {
// 		return l2
// 	}
// 	if l2 == nil {
// 		return l1
// 	}
// 	if l1.Val < l2.Val {
// 		l1.Next = mergeTwoLists(l1.Next, l2)
// 		return l1
// 	}
// 	l2.Next = mergeTwoLists(l1, l2.Next)
// 	return l2
// }

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	curr1, curr2 := list1, list2

	for curr2.Next != nil {
		if curr1.Val < curr2.Val {
			tmp := curr1.Next
			curr1.Next = curr2
			curr1.Next.Next = tmp
			curr1 = curr1.Next
			curr2 = curr2.Next
		} else if curr1.Val == curr2.Val {
			// curr1 = curr1.Next
			curr2 = curr2.Next
		} else if curr1.Val > curr2.Val {

		}

	}
	return nil
}

// func main() {

// }

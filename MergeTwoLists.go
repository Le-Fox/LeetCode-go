package leetcodego

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// if list1 happen to be NULL
	// we will simply return list2.
	if list1 == nil {
		return list2
	}
	// if list2 happen to be NULL
	// we will simply return list1.

	if list2 == nil {
		return list1
	}

	// if value pointend by l1 pointer is less than equal to value pointed by l2 pointer
	// we wall call recursively l1 -> next and whole l2 list.
	if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
		// we will call recursive l1 whole list and l2 -> next
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

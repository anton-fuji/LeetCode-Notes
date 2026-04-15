// Problem: 83. Remove Duplicates from Sorted List
// Difficulty: Easy
// Tags: Singly Linked List
// URL: https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/
// Date: 2026-04-15

package removedup

type ListNode struct {
	Val  int
	Next *ListNode
}

// before [1, 1, 2] → after [1,2]
func deleteDuplicates(head *ListNode) *ListNode {
	curr := head

	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return head
}

/**
* 19. Remove Nth Node From End of List
**/

package main

/**
* Definition for singly-linked list.
**/
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	sentinel := &ListNode{Next: head}
	slow, fast := sentinel, sentinel

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

    slow.Next = slow.Next.Next

    return sentinel.Next
}

/**
* 206. Reverse Linked List (easy)
**/

package main

/**
* Definition for singly-linked list.
**/
type ListNode struct {
	Val  int
	Next *ListNode
}

// iterative
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}

	return prev
}

// recursive
func reverseList2(head *ListNode) *ListNode {
	_, root := helper(head)

	return root
}

func helper(head *ListNode) (curr, root *ListNode) {
	if head == nil {
		return nil, nil
	}

	prev, root := helper(head.Next)
	if prev == nil {
		return head, head
	}

	curr = &ListNode{Val: head.Val}
	prev.Next = curr

	return curr, root
}

// recursive not really go style
func reverseList3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prev := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return prev
}

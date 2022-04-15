/**
* Get Node Value
**/

package main

import (
	"fmt"
)

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

// **wip
func getNode(llist *SinglyLinkedListNode, positionFromTail int32) int32 {
	slow, fast := llist, llist
	fmt.Println(slow, fast)
	for i := int32(0); i < positionFromTail+1; i++ {
		if fast.next != nil {
			fast = fast.next
		}
	}

	for fast != nil {
		fast = fast.next
		slow = slow.next
	}

	return slow.data
}


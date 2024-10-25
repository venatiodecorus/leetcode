package main

import (
	"fmt"
)

func main() {
	// Given the head of a linked list, reverse the nodes of the list k at a time, and return the modified list.
	// k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.
	// You may not alter the values in the list's nodes, only nodes themselves may be changed.

	// Test cases
	fmt.Println("Should be [3 2 1 4 5]:", reverseKGroup(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}, 3))
	printList(reverseKGroup(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}, 3))
	fmt.Println("Should be [2 1 4 3 5]:", reverseKGroup(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}, 2))
	printList(reverseKGroup(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}, 2))
	fmt.Println("Should be [1 2 3 4 5]:", reverseKGroup(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}, 1))
	printList(reverseKGroup(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}, 1))

}

func printList(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val, " ")
		l = l.Next
	}
	fmt.Println()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}

	dummy := &ListNode{Next: head}
	prevGroupEnd := dummy

	for {
		kth := getKthNode(prevGroupEnd, k)
		if kth == nil {
			break
		}

		groupStart := prevGroupEnd.Next
		groupEnd := kth.Next

		reverseList(groupStart, groupEnd)

		prevGroupEnd.Next = kth
		groupStart.Next = groupEnd

		prevGroupEnd = groupStart
	}

	return dummy.Next
}

func getKthNode(node *ListNode, k int) *ListNode {
	for node != nil && k > 0 {
		node = node.Next
		k--
	}

	return node
}

func reverseList(start, end *ListNode) {
	var prev *ListNode
	curr := start

	for curr != end {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
}

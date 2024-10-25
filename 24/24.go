package main

import (
	"fmt"
)

func main() {
	// Given a linked list, swap every two adjacent nodes and return its head. You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)

	// Test cases
	fmt.Println("Should be [2 1 4 3 5]:", swapPairs(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}))
	printList(swapPairs(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}))
	fmt.Println("Should be [1 2]:", swapPairs(&ListNode{Val: 2, Next: &ListNode{Val: 1}}))
	fmt.Println("Should be [1]:", swapPairs(&ListNode{Val: 1}))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	prev := dummy
	current := head

	for current != nil && current.Next != nil {
		first := current
		second := current.Next

		prev.Next = second
		first.Next = second.Next
		second.Next = first

		prev = first
		current = first.Next
	}

	return dummy.Next
}

func printList(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val, " ")
		l = l.Next
	}
	fmt.Println()
}

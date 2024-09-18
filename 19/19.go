package main

import "fmt"

func main() {
	// Given the head of a linked list, remove the nth node from the end of the list and return its head.

	// Test cases
	fmt.Println("Should be [1,2,3,5]:", removeNthFromEnd(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}, 2))
	fmt.Println("Should be [1]:", removeNthFromEnd(&ListNode{Val: 1}, 1))
	fmt.Println("Should be []:", removeNthFromEnd(&ListNode{Val: 1, Next: &ListNode{Val: 2}}, 1))
}

// ListNode is a struct for a linked list node
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	lastNodes := []*ListNode{}
	current := head
	index := 0

	for current != nil {
		//lastNodes[index] = current
		lastNodes = append(lastNodes, current)
		current = current.Next
		index++
	}

	removeIndex := index - n
	if removeIndex == 0 {
		return head.Next
	}

	lastNodes[removeIndex-1].Next = lastNodes[removeIndex].Next

	return head
}
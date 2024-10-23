package main

import (
	"fmt"
)

func main() {
	// You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
	// Merge all the linked-lists into one sorted linked-list and return it.

	// Test cases
	fmt.Println("Should be [1 1 2 3 4 4 5 6]:", mergeKLists([]*ListNode{
		&ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
		&ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
		&ListNode{Val: 2, Next: &ListNode{Val: 6}},
	}))
	fmt.Println("Should be []:", mergeKLists([]*ListNode{}))
	fmt.Println("Should be []:", mergeKLists([]*ListNode{nil}))

}

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	// Return empty ListNode array if input is empty
	if len(lists) == 0 {
		return nil
	}

	mainList := lists[0]

	for i:= 1; i < len(lists); i++ {
		mainList = merge(mainList, lists[i])
	}

	return mergeSort(mainList)
}

// Split a single list
func split(head *ListNode) *ListNode {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		if(fast != nil) {
			slow = slow.Next
		}
	}

	temp := slow.Next
	slow.Next = nil
	return temp
}

func merge(first *ListNode, second *ListNode) *ListNode {
	if(first == nil) {
		return second
	}
	if(second == nil) {
		return first
	}

	if(first.Val < second.Val) {
		first.Next = merge(first.Next, second)
		return first
	} else {
		second.Next = merge(first, second.Next)
		return second
	}
}

func mergeSort(head *ListNode) *ListNode {
	if(head == nil || head.Next == nil) {
		return head
	}

	second := split(head)

	head = mergeSort(head)
	second = mergeSort(second)

	return merge(head, second)
}
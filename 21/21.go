package main

import (
	"fmt"
)

func main() {
	// You are given the heads of two sorted linked lists list1 and list2.
	// Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
	// Return the head of the merged linked list.

	// Test cases
	fmt.Println("Should be [1,1,2,3,4,4]:", mergeTwoLists(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}, &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}))
	fmt.Println("Should be []:", mergeTwoLists(nil, nil))
	fmt.Println("Should be [0]:", mergeTwoLists(nil, &ListNode{Val: 0}))
	fmt.Println("Should be [1]:", mergeTwoLists(&ListNode{Val: 1}, nil))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Recursive solution
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2





    // // printList(list1);
	// node1 := list1
	// node2 := list2

	// if(list1 == nil && list2 == nil) {
	// 	return nil
	// } else if(list1 == nil) {
	// 	return list2
	// } else if(list2 == nil) {
	// 	return list1
	// }

	// output := &ListNode{}
	// if(node1.Val > node2.Val) {
	// 	output = &ListNode{Val: node1.Val, Next: nil}
	// 	node1 = node1.Next
	// } else {
	// 	output = &ListNode{Val: node2.Val, Next: nil}
	// 	node2 = node2.Next
	// }

	// for(node1 != nil && node2 != nil) {
	// 	if(node1.Val > node2.Val) {
	// 		appendList(output, node1)
	// 		node1 = node1.Next
	// 	} else {
	// 		appendList(output, node2)
	// 		node2 = node2.Next
	// 	}
	// }

	// // printList(output)
	// return output
}

func appendList(list *ListNode, newNode *ListNode) {
	for list.Next != nil {
		list = list.Next
	}

	list.Next = &ListNode{Val: newNode.Val, Next: nil}
}

func printList(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val, " ")
		l = l.Next
	}
	fmt.Println()
}
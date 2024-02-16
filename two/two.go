package two

import (
	"fmt"
	"math/big"
	"slices"
	"strconv"
	"strings"
)

func main() {
	num,_ := new(big.Int).SetString("1000000000000000000000000000001", 10)
	num2,_ := new(big.Int).SetString("465", 10)
	output := addTwoNumbers(makeList(num), makeList(num2))
	fmt.Println(output)
}

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    // n1,_ := strconv.Atoi(unrollList(l1))
    // n2,_ := strconv.Atoi(unrollList(l2))

	n1,_ := new(big.Int).SetString(unrollList(l1), 10)
	n2,_ := new(big.Int).SetString(unrollList(l2), 10)

    sum := new(big.Int).Add(n1, n2)
    return makeList(sum)
}

func unrollList(list *ListNode) string {
    currentNode := list
    var digits []string

    for currentNode != nil {
        digits = append(digits, strconv.Itoa(currentNode.Val))
        currentNode = currentNode.Next
    }

    slices.Reverse(digits)
    return strings.Join(digits, "")
}

func makeList(num *big.Int) *ListNode {
	arrString := strings.Split(num.String(), "")
    slices.Reverse(arrString)
	nodes := make([]*ListNode, len(arrString))
	for i := range arrString {
		v,_ := strconv.Atoi(arrString[i])
		nodes[i] = &ListNode{Val: v}
	}
	for n := range nodes[:len(nodes)-1] {
		nodes[n].Next = nodes[n+1]
	}

	return nodes[0]
}
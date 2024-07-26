package main

import (
	"fmt"
)

func main() {
	fmt.Println("Should be true:", isPalindrome(121))
	fmt.Println("Should be false:", isPalindrome(-121))
	fmt.Println("Should be false:", isPalindrome(10))
	fmt.Println("Should be true:", isPalindrome(0))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	// Reverse the number
	rev := 0
	orig := x
	for x > 0 {
		rev = rev*10 + x%10
		x /= 10
	}

	return orig == rev
}
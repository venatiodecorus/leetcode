package main

import (
	"fmt"
)

func main() {
	// Write a function to find the longest common prefix string amongst an array of strings.
	// If there is no common prefix, return an empty string "".
	fmt.Println("Should be 'fl':", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println("Should be '':", longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println("Should be 'a':", longestCommonPrefix([]string{"a"}))
	fmt.Println("Should be 'a':", longestCommonPrefix([]string{"ab", "a"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	pre := ""

	outer:
	for i,c := range strs[0] {
		for i2 := 1; i2 < len(strs); i2++ {
			if i >= len(strs[i2]) || strs[i2][i] != byte(c) {
				break outer
			}
		}
		pre = pre + string(c)
	}

	return pre
}
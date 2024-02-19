package main

import (
	"fmt"
)

func main() {
	fmt.Println("Should be 3: ", lengthOfLongestSubstring("abcabcbb"))
	fmt.Println("Should be 1: ", lengthOfLongestSubstring("bbbbb"))
	fmt.Println("Should be 3: ", lengthOfLongestSubstring("pwwkew"))
	fmt.Println("Should be 1: ", lengthOfLongestSubstring(" "))
	fmt.Println("Should be 2: ", lengthOfLongestSubstring("au"))
	fmt.Println("Should be 2: ", lengthOfLongestSubstring("aab"))
}


func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[rune]int)
    longest := 0
    start := 0

    for i, c := range s {
        if lastSeen, ok := charIndex[c]; ok && lastSeen >= start {
            start = lastSeen + 1
        }

        if i-start+1 > longest {
            longest = i - start + 1
        }

        charIndex[c] = i
    }

    return longest
	// var sub []rune
	// subIdx := 0

	// for i,c := range s {
	// 	if i == 0 {
	// 		sub = []rune(s[i:1])
	// 		continue
	// 	}
	// 	// if slices.Contains(sub, c) {
	// 	// 	if len(sub) <= len(s[subIdx:i]) {
	// 	// 		sub = []rune(s[subIdx:i])
	// 	// 		subIdx = i
	// 	// 		continue
	// 	// 	}
	// 	// }// else {
	// 	// 	sub = append(sub, c)
	// 	// //}


	// 	if strings.ContainsRune(s[subIdx:i], c) {
	// 		if len(sub) <= len(s[subIdx:i]) {
	// 			sub = []rune(s[subIdx:i])
	// 			subIdx = i + 1
	// 		}
	// 	}
	// }

	// if subIdx != len(s) {
	// 	sub = []rune(s[subIdx:])
	// }

    // return len(sub)
}
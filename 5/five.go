package main

import (
	"fmt"
)

func main() {
	fmt.Println("Should be 'bab' or 'aba': ", longestPalindrome("babad"))
	fmt.Println("Should be 'bb': ", longestPalindrome("cbbd"))
	fmt.Println("Should be 'a': ", longestPalindrome("a"))
}

func longestPalindrome(s string) string {
    if len(s) < 2 {
        return s
    }

    start, end := 0, 0
    for i := 0; i < len(s); i++ {
        len1 := expandAroundCenter(s, i, i)
        len2 := expandAroundCenter(s, i, i+1)
        lenMax := max(len1, len2)
        if lenMax > end - start {
            start = i - (lenMax-1)/2
            end = i + lenMax/2
        }
    }
    return s[start:end+1]
}

func expandAroundCenter(s string, left, right int) int {
    L, R := left, right
    for L >= 0 && R < len(s) && s[L] == s[R] {
        L--
        R++
    }
    return R - L - 1
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
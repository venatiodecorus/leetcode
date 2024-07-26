package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Should be 42:", myAtoi("42"))
	fmt.Println("Should be 0:", myAtoi("words and 987"))
	fmt.Println("Should be -42:", myAtoi("   -42"))
	fmt.Println("Should be 0:", myAtoi("+-12"))
}

func myAtoi(s string) int {
	neg := false
	pos := false
	// startNum := false
	offset := 48
	output := 0
    for _,c := range s {

		// // Check if character is numeric
		// if (c < 48 || c > 57) && !startNum {
		// 	continue
		// }
		if(c == 32) {
			continue
		}
		// If we hit non-numeric after starting the number, exit
		if (c < 48 || c > 57) && (c != 43 && c != 45) {
			break;
		}

		// Set negative if we find a '-'
		if c == '-' {
			neg = true
			continue
		}
		if c == '+' {
			pos = true
			continue
		}

		// We've passed whitespace and symbols, read number until non-numeric character
		// startNum = true

		val := int(c) - offset
		// Check for overflow
		if !neg && (output > math.MaxInt32/10 || (output == math.MaxInt32/10 && val > math.MaxInt32%10)) {
			return math.MaxInt32
		}
		if neg && (output > math.MaxInt32/10 || (output == math.MaxInt32/10 && val > (math.MaxInt32%10)+1)) {
			return math.MinInt32
		}
		output = output*10 + val
	}

	if pos && neg {
		return 0
	}
	if neg {
		return -output
	}

	return output
}
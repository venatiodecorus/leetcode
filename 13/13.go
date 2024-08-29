package main

import (
	"fmt"
)

func main() {
	// For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.
	// Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:
    // I can be placed before V (5) and X (10) to make 4 and 9.
    // X can be placed before L (50) and C (100) to make 40 and 90.
    // C can be placed before D (500) and M (1000) to make 400 and 900.
	// Given a roman numeral, convert it to an integer.
	fmt.Println("Should be 3:", romanToInt("III"))
	fmt.Println("Should be 4:", romanToInt("IV"))
	fmt.Println("Should be 9:", romanToInt("IX"))
	fmt.Println("Should be 58:", romanToInt("LVIII"))
	fmt.Println("Should be 1994:", romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	// output := ""
	outputNum := 0

	for i := 0; i < len(s); i++ {
		if s[i] == 'M' {
			outputNum += 1000
		} else if s[i] == 'D' {
			outputNum += 500
		} else if s[i] == 'C' {
			if  i < len(s)-1 && s[i+1] == 'M' {
				outputNum += 900
				i++
			} else if i < len(s)-1 && s[i+1] == 'D' {
				outputNum += 400
				i++
			} else {
				outputNum += 100
			}
		} else if s[i] == 'L' {
			outputNum += 50
		} else if s[i] == 'X' {
			if i < len(s)-1 && s[i+1] == 'C' {
				outputNum += 90
				i++
			} else if i < len(s)-1 && s[i+1] == 'L' {
				outputNum += 40
				i++
			} else {
				outputNum += 10
			}
		} else if s[i] == 'V' {
			outputNum += 5
		} else if s[i] == 'I' {
			if i < len(s)-1 && s[i+1] == 'X' {
				outputNum += 9
				i++
			} else if i < len(s)-1 && s[i+1] == 'V' {
				outputNum += 4
				i++
			} else {
				outputNum += 1
			}
		}
	}

	return outputNum
}

func startsWithFourOrNine(num int) bool {
	for num >= 10 {
		num /= 10
	}

	return num == 4 || num == 9
}

func getDigits(num int) []int {
	if num == 0 {
		return []int{0}
	}

	var digits []int
	for num > 0 {
		digit := num % 10
		digits = append([]int{digit}, digits...)
		num /= 10
	}
	return digits
}
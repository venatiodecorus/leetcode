package main

import (
	"fmt"
)

func main() {
	// Roman numerals are formed by appending the conversions of decimal place values from highest to lowest. Converting a decimal place value into a Roman numeral has the following rules:

    // If the value does not start with 4 or 9, select the symbol of the maximal value that can be subtracted from the input, append that symbol to the result, subtract its value, and convert the remainder to a Roman numeral.
    // If the value starts with 4 or 9 use the subtractive form representing one symbol subtracted from the following symbol, for example, 4 is 1 (I) less than 5 (V): IV and 9 is 1 (I) less than 10 (X): IX. Only the following subtractive forms are used: 4 (IV), 9 (IX), 40 (XL), 90 (XC), 400 (CD) and 900 (CM).
    // Only powers of 10 (I, X, C, M) can be appended consecutively at most 3 times to represent multiples of 10. You cannot append 5 (V), 50 (L), or 500 (D) multiple times. If you need to append a symbol 4 times use the subtractive form.

	// Given an integer, convert it to a Roman numeral.
	fmt.Println("Should be III:", intToRoman(3))
	fmt.Println("Should be IV:", intToRoman(4))
	fmt.Println("Should be IX:", intToRoman(9))
	fmt.Println("Should be LVIII:", intToRoman(58))
	fmt.Println("Should be MCMXCIV:", intToRoman(1994))
	fmt.Println("Should be MMMDCCXLIX:", intToRoman(3749))
}

func intToRoman(num int) string {
	// digits := getDigits(num)
	// output := ""

	// for i,n := range digits {
	// 	place := int(math.Pow(10, float64(len(digits) - i-1)))
	// 	switch place {
	// 	case 1000:
	// 		for j := 0; j < n; j++ {
	// 			output += "M"
	// 		}
	// 	case 100:

	// 	}
	// }

	// return ""

	rem := num
	output := ""
	for rem > 0 {
		if startsWithFourOrNine(rem) {
			if rem >= 900 {
				output += "CM"
				rem -= 900
			} else if rem >= 400 {
				output += "CD"
				rem -= 400
			} else if rem >= 90 {
				output += "XC"
				rem -= 90
			} else if rem >= 40 {
				output += "XL"
				rem -= 40
			} else if rem >= 9 {
				output += "IX"
				rem -= 9
			} else if rem >= 4 {
				output += "IV"
				rem -= 4
			}
		} else {
			if rem >= 1000 {
				output += "M"
				rem -= 1000
			} else if rem >= 500 {
				output += "D"
				rem -= 500
			} else if rem >= 100 {
				output += "C"
				rem -= 100
			} else if rem >= 50 {
				output += "L"
				rem -= 50
			} else if rem >= 10 {
				output += "X"
				rem -= 10
			} else if rem >= 5 {
				output += "V"
				rem -= 5
			} else if rem >= 1 {
				output += "I"
				rem -= 1
			}
		}
	}
	return output
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
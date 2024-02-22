package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
)

func main() {
	fmt.Println("Should be 321", reverse(123))
	fmt.Println("Should be -321", reverse(-123))
	fmt.Println("Should be 9646324351", reverse(1534236469))
}

func reverse(x int) int {
	str := strconv.Itoa(x)
	runes := []rune(str)
	slices.Reverse(runes)
	rStr := string(runes)

	if rStr[len(rStr)-1] == '-' {
		out,err := strconv.Atoi(rStr[:len(rStr)-1])
		if err != nil {
			return 0
		}
		if out < math.MinInt32 || out > math.MaxInt32 {
			return 0
		}
		return -out
	}
	out,err := strconv.Atoi(rStr)
	if err != nil {
		return 0
	}
	if out < math.MinInt32 || out > math.MaxInt32 {
		return 0
	}
    return out
}
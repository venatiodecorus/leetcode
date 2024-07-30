package main

import (
	"fmt"
)

func main() {
	// You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
	// Find two lines that together with the x-axis form a container, such that the container contains the most water.
	// Return the maximum amount of water a container can store.
	// Notice that you may not slant the container.
	fmt.Println("Should be 49:", maxArea([]int{1,8,6,2,5,4,8,3,7}))
	fmt.Println("Should be 1:", maxArea([]int{1,1}))
	fmt.Println("Should be 16:", maxArea([]int{4,3,2,1,4}))
}

func maxArea(height []int) int {
	// Brute force
	// max := 0
    // // Iterate through all lines
	// for i,h := range height {
	// 	// Compare to every other line
	// 	for y := i+1; y < len(height); y++ {
    //         area := (y - i) * min(h, height[y])
    //         // Update max area
    //         max = maxf(max, area)
	// 	}
	// }

	// Two pointer
	left, right := 0, len(height) - 1
	max := 0

	for left < right {
		area := (right - left) * min(height[left], height[right])
		max = maxf(max, area)

		if height[left] < height[right] {
			left ++
		} else {
			right--
		}
	}

	return max
}

func maxf(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
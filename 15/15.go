package main

import (
	"fmt"
	"sort"
)

func main() {
	// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
	// Notice that the solution set must not contain duplicate triplets.

	// Test cases
	fmt.Println("Should be [[-1,-1,2],[-1,0,1]]:", threeSum([]int{-1,0,1,2,-1,-4}))
	fmt.Println("Should be []:", threeSum([]int{}))

}

func threeSum(nums []int) [][]int {
	result := [][]int{}

	// Sort the array to make it easier to skip duplicates and use two pointers
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		// Skip duplicates for the first number
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// Use two pointers to find pairs that sum to -nums[i]
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Skip duplicates for the second number
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// Skip duplicates for the third number
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				// Move the pointers
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}
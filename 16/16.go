package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// Given an integer array nums of length n and an integer target, find three integers in nums such that the sum is closest to target.
	// Return the sum of the three integers.
	// You may assume that each input would have exactly one solution.

	// Test cases
	// fmt.Println("Should be 2:", threeSumClosest([]int{-1,2,1,-4}, 1))
	// fmt.Println("Should be 0:", threeSumClosest([]int{0,0,0}, 1))
	// fmt.Println("Should be 2:", threeSumClosest([]int{1,1,1,0}, -100))
	fmt.Println("Should be 0:", threeSumClosest([]int{-4,2,2,3,3,3}, 0))
}

func threeSumClosest(nums []int, target int) int {
    // Sort the array to make it easier to skip duplicates and use two pointers
    sort.Ints(nums)

    minDiff := math.MaxInt32
    aSum := 0

    for i := 0; i < len(nums)-2; i++ {
        // Skip duplicates for the first number
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        // Use two pointers to find pairs that sum to -nums[i]
        left, right := i+1, len(nums)-1
        for left < right {
            sum := nums[i] + nums[left] + nums[right]
            diff := abs(target - sum)

            if diff < minDiff {
                minDiff = diff
                aSum = sum
            }

            // Move the pointers based on the comparison between sum and target
            if sum < target {
                left++
            } else {
                right--
            }
        }
    }

    return aSum
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
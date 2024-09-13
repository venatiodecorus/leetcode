package main

import (
	"fmt"
	"sort"
)

func main() {
	// Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:
    // 0 <= a, b, c, d < n
    // a, b, c, and d are distinct.
    // nums[a] + nums[b] + nums[c] + nums[d] == target
	// You may return the answer in any order.

	// Test cases
	fmt.Println("Should be [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]:", fourSum([]int{1,0,-1,0,-2,2}, 0))
	fmt.Println("Should be [[2,2,2,2]]:", fourSum([]int{2,2,2,2,2}, 8))
	// fmt.Println("Should be []:", fourSum([]int{}, 0))
	// fmt.Println("Should be []:", fourSum([]int{0}, 0))
	// fmt.Println("Should be []:", fourSum([]int{0,0,0,0}, 0))
	// fmt.Println("Should be [[-3,-2,2,3],[-3,-1,1,3],[-3,0,0,3],[-3,0,1,2],[-2,-1,0,3],[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]:", fourSum([]int{1,0,-1,0,-2,2}, 0))
	// fmt.Println("Should be [[-1,-1,1,1]]:", fourSum([]int{1,-2,-5,-4,-3,3,3,5}, -11))
}

func fourSum(nums []int, target int) [][]int {
	result := [][]int{}
	uniqueQuadruplets := make(map[[4]int]struct{})

	// Sort the array to make it easier to skip duplicates and use two pointers
	sort.Ints(nums)

	for i := 0; i < len(nums)-3; i++ {
		// Skip duplicates for the first number
		// if i > 0 && nums[i] == nums[i-1] {
		// 	continue
		// }

		for j := i + 1; j < len(nums)-2; j++ {
			// Use two pointers to find pairs that sum to -nums[i]
			left, right := j+1, len(nums)-1
			for left < right {
				sum := nums[i] + nums[left] + nums[right] + nums[j]
				if sum == target {
                    quadruplet := [4]int{nums[i], nums[j], nums[left], nums[right]}
                    if _, exists := uniqueQuadruplets[quadruplet]; !exists {
                        uniqueQuadruplets[quadruplet] = struct{}{}
                        result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
                    }
					// // Skip duplicates for the second number
					// for left < right && nums[left] == nums[left+1] {
					// 	left++
					// }
					// // Skip duplicates for the third number
					// for left < right && nums[right] == nums[right-1] {
					// 	right--
					// }

					// Move the pointers
					left++
					right--
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}

	return result
}


// func fourSum(nums []int, target int) [][]int {
// 	result := [][]int{}

// 	// Sort the array to make it easier to skip duplicates and use two pointers
// 	sort.Ints(nums)

// 	for i := 0; i < len(nums)-2; i++ {
// 		// Skip duplicates for the first number
// 		// if i > 0 && nums[i] == nums[i-1] {
// 		// 	continue
// 		// }

// 		// Use two pointers to find pairs that sum to -nums[i]
// 		left, right, right2 := i+1, len(nums)-1, len(nums)-2
// 		for left < right {
// 			sum := nums[i] + nums[left] + nums[right] + nums[right2]
// 			if sum == target {
// 				result = append(result, []int{nums[i], nums[left], nums[right], nums[right2]})

// 				// Skip duplicates for the second number
// 				for left < right && nums[left] == nums[left+1] {
// 					left++
// 				}
// 				// Skip duplicates for the third number
// 				for left < right && nums[right] == nums[right-1] {
// 					right -= 2
// 				}

// 				// Move the pointers
// 				left++
// 				right -= 2
// 			} else if sum < 0 {
// 				left++
// 			} else {
// 				right -=2
// 			}
// 		}
// 	}

// 	return result
// }
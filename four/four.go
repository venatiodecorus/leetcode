package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findMedianSortedArrays([]int{1,3}, []int{2}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    nums := append(nums1, nums2...)
	sort.Ints(nums)
	
	if len(nums) % 2 == 0 {
		return float64(nums[len(nums)/2] + nums[len(nums)/2-1]) / 2
	} else {
		return float64(nums[len(nums)/2])
	}
}
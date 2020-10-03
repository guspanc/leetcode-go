package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println([]int{1, 5, 11, 5})
}

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum&1 == 1 {
		return false
	}
	sort.Slice(nums, func(a, b int) bool {
		return nums[a] > nums[b]
	})
	return helper(nums, 0, 0, sum/2)
}

func helper(nums []int, index, num, target int) bool {
	if num == target {
		return true
	}
	if index >= len(nums) {
		return false
	}
	if num+nums[index] > target {
		return false
	}
	// next line, num+nums[index] must place at front. else Time Limit Exceeded
	return helper(nums, index+1, num+nums[index], target) || helper(nums, index+1, num, target)
}

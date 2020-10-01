package main

import "fmt"

func main() {
	input := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDuplicates(input))
}

/*
- 1 ≤ a[i] ≤ n

- where j = abs(a[i]) - 1
- set a[j] to a[j]*-1 to know that we have seen this abs(a[i])
- then if a[j] < 0, we know we have seen abs(a[i])
*/

func findDuplicates(nums []int) []int {
	res := []int{}
	for _, v := range nums {
		if nums[abs(v)-1] < 0 {
			res = append(res, abs(v))
		} else {
			nums[abs(v)-1] = nums[abs(v)-1] * -1
		}
	}
	return res
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

package main

import "fmt"

func main() {
	fmt.Println(sumOddLengthSubarrays([]int{1, 4, 2, 5, 3}))
}

func sumOddLengthSubarrays(arr []int) int {
	res := 0

	for i := 0; i < len(arr); i++ {
		left := i + 1
		right := len(arr) - i

		res += (left*right + 1) / 2 * arr[i]
	}

	return res
}

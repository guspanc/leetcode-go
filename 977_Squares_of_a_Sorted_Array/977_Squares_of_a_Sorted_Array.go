package main

import "fmt"

func main() {
	input := []int{-4, -1, 0, 3, 10}
	res := sortedSquares(input)
	fmt.Println(input)
	fmt.Println(res)
}

func sortedSquares(A []int) []int {
	res := make([]int, len(A))
	i, j, p := 0, len(A)-1, len(A)-1
	for i <= j {
		if abs(A[i]) > abs(A[j]) {
			res[p] = A[i] * A[i]
			i++
		} else {
			res[p] = A[j] * A[j]
			j--
		}
		p--
	}
	return res
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

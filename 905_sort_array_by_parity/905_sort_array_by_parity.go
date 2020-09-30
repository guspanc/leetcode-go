package main

import "fmt"

func main() {
	before := []int{3, 1, 2, 4}
	after := sortArrayByParity(before)
	fmt.Println(before)
	fmt.Println(after)
}

func sortArrayByParity(A []int) []int {
	res := make([]int, len(A))
	i, j := 0, len(A)-1
	for _, v := range A {
		if v%2 == 0 {
			res[i] = v
			i++
		} else {
			res[j] = v
			j--
		}
	}
	return res
}

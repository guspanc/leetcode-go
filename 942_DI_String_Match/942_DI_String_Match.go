package main

import "fmt"

func main() {
	input := "III"
	fmt.Println(diStringMatch(input))
}

func diStringMatch(S string) []int {
	i, ic, dc := 0, 0, len(S)
	res := make([]int, len(S)+1)
	for _, v := range S {
		if string(v) == "I" {
			res[i] = ic
			ic++
		} else {
			res[i] = dc
			dc--
		}
		i++
	}
	res[i] = ic
	return res
}

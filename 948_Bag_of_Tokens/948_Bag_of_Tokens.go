package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(bagOfTokensScore([]int{100, 200, 300, 400}, 200))
}

func bagOfTokensScore(tokens []int, P int) int {
	sort.Ints(tokens)
	var maxpoints, points int
	l, r := 0, len(tokens)-1
	for l <= r {
		if P >= tokens[l] {
			P -= tokens[l]
			points++
			maxpoints = max(maxpoints, points)
			l++
		} else if points > 0 {
			points--
			P += tokens[r]
			r--
		} else {
			return maxpoints
		}
	}
	return maxpoints
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

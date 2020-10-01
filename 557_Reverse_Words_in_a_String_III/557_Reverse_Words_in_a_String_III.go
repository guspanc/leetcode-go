package main

import "fmt"

func main() {
	input := "Let's take LeetCode contest"
	fmt.Println(reverseWords(input))
}

func reverseWords(s string) string {
	// get all splits
	splits := []int{}
	for i, v := range s {
		if v == ' ' {
			splits = append(splits, i)
		}
	}
	splits = append(splits, len(s))

	// pointers
	is := 0
	split := splits[is]
	j := 1
	res := make([]rune, len(s))
	for i, v := range s {
		// if ' ' reset j pointer and shift to next split
		if v == ' ' {
			is++
			split = splits[is]
			j = 1
			res[i] = ' '
		} else {
			res[split-j] = v
			j++
		}
	}

	return string(res)
}

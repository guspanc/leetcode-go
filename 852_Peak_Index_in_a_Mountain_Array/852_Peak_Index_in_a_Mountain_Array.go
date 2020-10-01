package main

import "fmt"

func main() {
	input := []int{40, 48, 61, 75, 100, 99, 98, 39, 30, 10}
	fmt.Println(peakIndexInMountainArray(input))
}

func peakIndexInMountainArray(arr []int) int {
	var m int
	l, r := 0, len(arr)-2
	for l < r {
		m = (r + l) / 2
		if arr[m] > arr[m+1] {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

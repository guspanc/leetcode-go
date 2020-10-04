package main

/*
Input: S = "ababcbacadefegdehijhklij"
Output: [9,7,8]
Explanation:
The partition is "ababcbaca", "defegde", "hijhklij".
This is a partition so that each letter appears in at most one part.
A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it splits S into less parts.
*/

func main() {

}

func partitionLabels(S string) []int {
	result := []int{}

	// get locations
	location := map[byte]int{}
	for i := range S {
		v := S[i]
		location[v] = i
	}

	// get partitions
	start := 0
	end := location[S[0]]
	for i := 0; i < len(S); i++ {
		end = max(end, location[S[i]])
		if i == end {
			result = append(result, end-start+1)
			start = i + 1
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

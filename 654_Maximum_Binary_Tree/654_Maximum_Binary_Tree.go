package main

import "fmt"

func main() {
	fmt.Println([]int{3, 2, 1, 6, 0, 5})
}

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	index := 0
	max := nums[index]
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
			index = i
		}
	}
	node := &TreeNode{
		Val: max,
	}
	node.Left = constructMaximumBinaryTree(nums[:index])
	node.Right = constructMaximumBinaryTree(nums[index+1:])
	return node
}
